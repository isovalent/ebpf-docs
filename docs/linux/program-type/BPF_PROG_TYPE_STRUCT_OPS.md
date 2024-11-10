---
title: "Program Type 'BPF_PROG_TYPE_STRUCT_OPS'"
description: "This page documents the 'BPF_PROG_TYPE_STRUCT_OPS' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_STRUCT_OPS`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_STRUCT_OPS) -->
[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/27ae7997a66174cb8afd6a75b3989f5e0c1b9e5a)
<!-- [/FEATURE_TAG] -->

This program types allows for the implementation of certain kernel features in BPF. 

## Usage

The kernel uses the "struct ops" pattern in C to implement interfaces. The kernel defines a struct with function pointers as field types. An implementation can create an instance of this struct and provide pointers to its own functions that implement the signatures.

This program type allows for the creation of these struct implementations with BPF so locations in the kernel that allow it can BPF to implement the functionality. For example the TCP congestion control algorithm.

## Context

The context provided to each function is an array of 64 bit integers, representing the argument list of the function in the struct it is implementing. The `BPF_PROG` macro from libbpf can be used to write your program in a recognizable way, the macro will take care of unpacking the array into named variables of the provided type.

## Attachment

### User perspective

From a eBPF user perspective, all one has to do is to create an instance of the struct to be used as implementation, any function pointer fields you wish to use should point to eBPF programs. This struct must be placed in the `.struct_ops` section like so:

```c
/* Copyright (c) 2019 Facebook */

SEC(".struct_ops")
struct tcp_congestion_ops dctcp_nouse = {
	.init		= (void *)dctcp_init,
	.set_state	= (void *)dctcp_state,
	.flags		= TCP_CONG_NEEDS_ECN,
	.name		= "bpf_dctcp_nouse",
};
```

The full file is turned into a light skeleton. After loading the light skeleton the struct will appear as a map in the `skel->maps` structure. Then the map is passed to the `bpf_map__attach_struct_ops` function to attach it.

### Loader/kernel perspective

The loader locates the struct(s) in the `.struct_ops` section. It uses BTF to get the name and structure of the struct. The values of non-function fields are present in the `.struct_ops` section itself and for the function pointer fields that were set there will be ELF relocation entries pointing to the respective eBPF programs.

The loader will figure out which eBPF programs were referenced and load them. The loader then will inspect the Vmlinux BTF of the host to find the BTF type ID of the struct_ops struct. Once we have that, the loader creates a [`BPF_MAP_TYPE_STRUCT_OPS`](../map-type/BPF_MAP_TYPE_STRUCT_OPS.md) map. While loading the BTF type ID of the struct_ops struct is provided via the [`btf_vmlinux_value_type_id`](../syscall/BPF_MAP_CREATE.md#btf_vmlinux_value_type_id) argument. The [`btf_value_type_id`](../syscall/BPF_MAP_CREATE.md#btf_value_type_id) is set to the BTF ID of the programs struct. The key of the map must be an `int`/4 bytes and `max_entries` of 1. 

The kernel has now allocated memory for one instance of the struct_ops struct and associated it with the call location. So the next step for the loader is to write to the only element in the map, element `0`. The value of the map is the instance of the struct as C would lay it out in memory. Except the values of the function pointer fields are populated with the file descriptors of the eBPF programs, the kernel will transform these into the actual memory locations of the JITed eBPF programs. As soon as the update happens, the struct_ops programs are attached. The map holds the refcounts to the programs, so they don't have to be pinned. As soon as element `0` of the map is deleted or the map is cleaned up due to being unreferenced, the struct_ops programs are detached. So typically the map gets pinned.

## Example

??? example "DC TCP"
    ```c
    // SPDX-License-Identifier: GPL-2.0
    /* Copyright (c) 2019 Facebook */

    /* WARNING: This implemenation is not necessarily the same
    * as the tcp_dctcp.c.  The purpose is mainly for testing
    * the kernel BPF logic.
    */

    #include <stddef.h>
    #include <linux/bpf.h>
    #include <linux/types.h>
    #include <linux/stddef.h>
    #include <linux/tcp.h>
    #include <bpf/bpf_helpers.h>
    #include <bpf/bpf_tracing.h>
    #include "bpf_tcp_helpers.h"

    char _license[] SEC("license") = "GPL";

    int stg_result = 0;

    struct {
        __uint(type, BPF_MAP_TYPE_SK_STORAGE);
        __uint(map_flags, BPF_F_NO_PREALLOC);
        __type(key, int);
        __type(value, int);
    } sk_stg_map SEC(".maps");

    #define DCTCP_MAX_ALPHA	1024U

    struct dctcp {
        __u32 old_delivered;
        __u32 old_delivered_ce;
        __u32 prior_rcv_nxt;
        __u32 dctcp_alpha;
        __u32 next_seq;
        __u32 ce_state;
        __u32 loss_cwnd;
    };

    static unsigned int dctcp_shift_g = 4; /* g = 1/2^4 */
    static unsigned int dctcp_alpha_on_init = DCTCP_MAX_ALPHA;

    static __always_inline void dctcp_reset(const struct tcp_sock *tp,
                        struct dctcp *ca)
    {
        ca->next_seq = tp->snd_nxt;

        ca->old_delivered = tp->delivered;
        ca->old_delivered_ce = tp->delivered_ce;
    }

    SEC("struct_ops/dctcp_init")
    void BPF_PROG(dctcp_init, struct sock *sk)
    {
        const struct tcp_sock *tp = tcp_sk(sk);
        struct dctcp *ca = inet_csk_ca(sk);
        int *stg;

        ca->prior_rcv_nxt = tp->rcv_nxt;
        ca->dctcp_alpha = min(dctcp_alpha_on_init, DCTCP_MAX_ALPHA);
        ca->loss_cwnd = 0;
        ca->ce_state = 0;

        stg = bpf_sk_storage_get(&sk_stg_map, (void *)tp, NULL, 0);
        if (stg) {
            stg_result = *stg;
            bpf_sk_storage_delete(&sk_stg_map, (void *)tp);
        }
        dctcp_reset(tp, ca);
    }

    SEC("struct_ops/dctcp_ssthresh")
    __u32 BPF_PROG(dctcp_ssthresh, struct sock *sk)
    {
        struct dctcp *ca = inet_csk_ca(sk);
        struct tcp_sock *tp = tcp_sk(sk);

        ca->loss_cwnd = tp->snd_cwnd;
        return max(tp->snd_cwnd - ((tp->snd_cwnd * ca->dctcp_alpha) >> 11U), 2U);
    }

    SEC("struct_ops/dctcp_update_alpha")
    void BPF_PROG(dctcp_update_alpha, struct sock *sk, __u32 flags)
    {
        const struct tcp_sock *tp = tcp_sk(sk);
        struct dctcp *ca = inet_csk_ca(sk);

        /* Expired RTT */
        if (!before(tp->snd_una, ca->next_seq)) {
            __u32 delivered_ce = tp->delivered_ce - ca->old_delivered_ce;
            __u32 alpha = ca->dctcp_alpha;

            /* alpha = (1 - g) * alpha + g * F */

            alpha -= min_not_zero(alpha, alpha >> dctcp_shift_g);
            if (delivered_ce) {
                __u32 delivered = tp->delivered - ca->old_delivered;

                /* If dctcp_shift_g == 1, a 32bit value would overflow
                * after 8 M packets.
                */
                delivered_ce <<= (10 - dctcp_shift_g);
                delivered_ce /= max(1U, delivered);

                alpha = min(alpha + delivered_ce, DCTCP_MAX_ALPHA);
            }
            ca->dctcp_alpha = alpha;
            dctcp_reset(tp, ca);
        }
    }

    static __always_inline void dctcp_react_to_loss(struct sock *sk)
    {
        struct dctcp *ca = inet_csk_ca(sk);
        struct tcp_sock *tp = tcp_sk(sk);

        ca->loss_cwnd = tp->snd_cwnd;
        tp->snd_ssthresh = max(tp->snd_cwnd >> 1U, 2U);
    }

    SEC("struct_ops/dctcp_state")
    void BPF_PROG(dctcp_state, struct sock *sk, __u8 new_state)
    {
        if (new_state == TCP_CA_Recovery &&
            new_state != BPF_CORE_READ_BITFIELD(inet_csk(sk), icsk_ca_state))
            dctcp_react_to_loss(sk);
        /* We handle RTO in dctcp_cwnd_event to ensure that we perform only
        * one loss-adjustment per RTT.
        */
    }

    static __always_inline void dctcp_ece_ack_cwr(struct sock *sk, __u32 ce_state)
    {
        struct tcp_sock *tp = tcp_sk(sk);

        if (ce_state == 1)
            tp->ecn_flags |= TCP_ECN_DEMAND_CWR;
        else
            tp->ecn_flags &= ~TCP_ECN_DEMAND_CWR;
    }

    /* Minimal DCTP CE state machine:
    *
    * S:	0 <- last pkt was non-CE
    *	1 <- last pkt was CE
    */
    static __always_inline
    void dctcp_ece_ack_update(struct sock *sk, enum tcp_ca_event evt,
                __u32 *prior_rcv_nxt, __u32 *ce_state)
    {
        __u32 new_ce_state = (evt == CA_EVENT_ECN_IS_CE) ? 1 : 0;

        if (*ce_state != new_ce_state) {
            /* CE state has changed, force an immediate ACK to
            * reflect the new CE state. If an ACK was delayed,
            * send that first to reflect the prior CE state.
            */
            if (inet_csk(sk)->icsk_ack.pending & ICSK_ACK_TIMER) {
                dctcp_ece_ack_cwr(sk, *ce_state);
                bpf_tcp_send_ack(sk, *prior_rcv_nxt);
            }
            inet_csk(sk)->icsk_ack.pending |= ICSK_ACK_NOW;
        }
        *prior_rcv_nxt = tcp_sk(sk)->rcv_nxt;
        *ce_state = new_ce_state;
        dctcp_ece_ack_cwr(sk, new_ce_state);
    }

    SEC("struct_ops/dctcp_cwnd_event")
    void BPF_PROG(dctcp_cwnd_event, struct sock *sk, enum tcp_ca_event ev)
    {
        struct dctcp *ca = inet_csk_ca(sk);

        switch (ev) {
        case CA_EVENT_ECN_IS_CE:
        case CA_EVENT_ECN_NO_CE:
            dctcp_ece_ack_update(sk, ev, &ca->prior_rcv_nxt, &ca->ce_state);
            break;
        case CA_EVENT_LOSS:
            dctcp_react_to_loss(sk);
            break;
        default:
            /* Don't care for the rest. */
            break;
        }
    }

    SEC("struct_ops/dctcp_cwnd_undo")
    __u32 BPF_PROG(dctcp_cwnd_undo, struct sock *sk)
    {
        const struct dctcp *ca = inet_csk_ca(sk);

        return max(tcp_sk(sk)->snd_cwnd, ca->loss_cwnd);
    }

    extern void tcp_reno_cong_avoid(struct sock *sk, __u32 ack, __u32 acked) __ksym;

    SEC("struct_ops/dctcp_reno_cong_avoid")
    void BPF_PROG(dctcp_cong_avoid, struct sock *sk, __u32 ack, __u32 acked)
    {
        tcp_reno_cong_avoid(sk, ack, acked);
    }

    SEC(".struct_ops")
    struct tcp_congestion_ops dctcp_nouse = {
        .init		= (void *)dctcp_init,
        .set_state	= (void *)dctcp_state,
        .flags		= TCP_CONG_NEEDS_ECN,
        .name		= "bpf_dctcp_nouse",
    };

    SEC(".struct_ops")
    struct tcp_congestion_ops dctcp = {
        .init		= (void *)dctcp_init,
        .in_ack_event   = (void *)dctcp_update_alpha,
        .cwnd_event	= (void *)dctcp_cwnd_event,
        .ssthresh	= (void *)dctcp_ssthresh,
        .cong_avoid	= (void *)dctcp_cong_avoid,
        .undo_cwnd	= (void *)dctcp_cwnd_undo,
        .set_state	= (void *)dctcp_state,
        .flags		= TCP_CONG_NEEDS_ECN,
        .name		= "bpf_dctcp",
    };
    ```

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for struct ops programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
    * [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md)
    * [`bpf_ktime_get_tai_ns`](../helper-function/bpf_ktime_get_tai_ns.md)
    * [`bpf_loop`](../helper-function/bpf_loop.md)
    * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
    * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
    * [`bpf_map_lookup_percpu_elem`](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
    * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
    * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
    * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
    * [`bpf_per_cpu_ptr`](../helper-function/bpf_per_cpu_ptr.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
    * [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md)
    * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
    * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
    * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
    * [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md)
    * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
    * [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md)
    * [`bpf_sk_storage_delete`](../helper-function/bpf_sk_storage_delete.md)
    * [`bpf_sk_storage_get`](../helper-function/bpf_sk_storage_get.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_tcp_send_ack`](../helper-function/bpf_tcp_send_ack.md)
    * [`bpf_this_cpu_ptr`](../helper-function/bpf_this_cpu_ptr.md)
    * [`bpf_timer_cancel`](../helper-function/bpf_timer_cancel.md)
    * [`bpf_timer_init`](../helper-function/bpf_timer_init.md)
    * [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md)
    * [`bpf_timer_start`](../helper-function/bpf_timer_start.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
    * [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md)
    * [`bpf_user_ringbuf_drain`](../helper-function/bpf_user_ringbuf_drain.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
??? abstract "Supported kfuncs"
    - [`bbr_cwnd_event`](../kfuncs/bbr_cwnd_event.md)
    - [`bbr_init`](../kfuncs/bbr_init.md)
    - [`bbr_main`](../kfuncs/bbr_main.md)
    - [`bbr_min_tso_segs`](../kfuncs/bbr_min_tso_segs.md)
    - [`bbr_set_state`](../kfuncs/bbr_set_state.md)
    - [`bbr_sndbuf_expand`](../kfuncs/bbr_sndbuf_expand.md)
    - [`bbr_ssthresh`](../kfuncs/bbr_ssthresh.md)
    - [`bbr_undo_cwnd`](../kfuncs/bbr_undo_cwnd.md)
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md)
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md)
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md)
    - [`bpf_cgroup_acquire`](../kfuncs/bpf_cgroup_acquire.md)
    - [`bpf_cgroup_ancestor`](../kfuncs/bpf_cgroup_ancestor.md)
    - [`bpf_cgroup_from_id`](../kfuncs/bpf_cgroup_from_id.md)
    - [`bpf_cgroup_release`](../kfuncs/bpf_cgroup_release.md)
    - [`bpf_cpumask_acquire`](../kfuncs/bpf_cpumask_acquire.md)
    - [`bpf_cpumask_and`](../kfuncs/bpf_cpumask_and.md)
    - [`bpf_cpumask_any_and_distribute`](../kfuncs/bpf_cpumask_any_and_distribute.md)
    - [`bpf_cpumask_any_distribute`](../kfuncs/bpf_cpumask_any_distribute.md)
    - [`bpf_cpumask_clear`](../kfuncs/bpf_cpumask_clear.md)
    - [`bpf_cpumask_clear_cpu`](../kfuncs/bpf_cpumask_clear_cpu.md)
    - [`bpf_cpumask_copy`](../kfuncs/bpf_cpumask_copy.md)
    - [`bpf_cpumask_create`](../kfuncs/bpf_cpumask_create.md)
    - [`bpf_cpumask_empty`](../kfuncs/bpf_cpumask_empty.md)
    - [`bpf_cpumask_equal`](../kfuncs/bpf_cpumask_equal.md)
    - [`bpf_cpumask_first`](../kfuncs/bpf_cpumask_first.md)
    - [`bpf_cpumask_first_and`](../kfuncs/bpf_cpumask_first_and.md)
    - [`bpf_cpumask_first_zero`](../kfuncs/bpf_cpumask_first_zero.md)
    - [`bpf_cpumask_full`](../kfuncs/bpf_cpumask_full.md)
    - [`bpf_cpumask_intersects`](../kfuncs/bpf_cpumask_intersects.md)
    - [`bpf_cpumask_or`](../kfuncs/bpf_cpumask_or.md)
    - [`bpf_cpumask_release`](../kfuncs/bpf_cpumask_release.md)
    - [`bpf_cpumask_set_cpu`](../kfuncs/bpf_cpumask_set_cpu.md)
    - [`bpf_cpumask_setall`](../kfuncs/bpf_cpumask_setall.md)
    - [`bpf_cpumask_subset`](../kfuncs/bpf_cpumask_subset.md)
    - [`bpf_cpumask_test_and_clear_cpu`](../kfuncs/bpf_cpumask_test_and_clear_cpu.md)
    - [`bpf_cpumask_test_and_set_cpu`](../kfuncs/bpf_cpumask_test_and_set_cpu.md)
    - [`bpf_cpumask_test_cpu`](../kfuncs/bpf_cpumask_test_cpu.md)
    - [`bpf_cpumask_weight`](../kfuncs/bpf_cpumask_weight.md)
    - [`bpf_cpumask_xor`](../kfuncs/bpf_cpumask_xor.md)
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md)
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md)
    - [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md)
    - [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md)
    - [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md)
    - [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md)
    - [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md)
    - [`bpf_iter_bits_destroy`](../kfuncs/bpf_iter_bits_destroy.md)
    - [`bpf_iter_bits_new`](../kfuncs/bpf_iter_bits_new.md)
    - [`bpf_iter_bits_next`](../kfuncs/bpf_iter_bits_next.md)
    - [`bpf_iter_css_destroy`](../kfuncs/bpf_iter_css_destroy.md)
    - [`bpf_iter_css_new`](../kfuncs/bpf_iter_css_new.md)
    - [`bpf_iter_css_next`](../kfuncs/bpf_iter_css_next.md)
    - [`bpf_iter_css_task_destroy`](../kfuncs/bpf_iter_css_task_destroy.md)
    - [`bpf_iter_css_task_new`](../kfuncs/bpf_iter_css_task_new.md)
    - [`bpf_iter_css_task_next`](../kfuncs/bpf_iter_css_task_next.md)
    - [`bpf_iter_num_destroy`](../kfuncs/bpf_iter_num_destroy.md)
    - [`bpf_iter_num_new`](../kfuncs/bpf_iter_num_new.md)
    - [`bpf_iter_num_next`](../kfuncs/bpf_iter_num_next.md)
    - [`bpf_iter_task_destroy`](../kfuncs/bpf_iter_task_destroy.md)
    - [`bpf_iter_task_new`](../kfuncs/bpf_iter_task_new.md)
    - [`bpf_iter_task_next`](../kfuncs/bpf_iter_task_next.md)
    - [`bpf_iter_task_vma_destroy`](../kfuncs/bpf_iter_task_vma_destroy.md)
    - [`bpf_iter_task_vma_new`](../kfuncs/bpf_iter_task_vma_new.md)
    - [`bpf_iter_task_vma_next`](../kfuncs/bpf_iter_task_vma_next.md)
    - [`bpf_list_pop_back`](../kfuncs/bpf_list_pop_back.md)
    - [`bpf_list_pop_front`](../kfuncs/bpf_list_pop_front.md)
    - [`bpf_list_push_back_impl`](../kfuncs/bpf_list_push_back_impl.md)
    - [`bpf_list_push_front_impl`](../kfuncs/bpf_list_push_front_impl.md)
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md)
    - [`bpf_obj_drop_impl`](../kfuncs/bpf_obj_drop_impl.md)
    - [`bpf_obj_new_impl`](../kfuncs/bpf_obj_new_impl.md)
    - [`bpf_percpu_obj_drop_impl`](../kfuncs/bpf_percpu_obj_drop_impl.md)
    - [`bpf_percpu_obj_new_impl`](../kfuncs/bpf_percpu_obj_new_impl.md)
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md)
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md)
    - [`bpf_rbtree_add_impl`](../kfuncs/bpf_rbtree_add_impl.md)
    - [`bpf_rbtree_first`](../kfuncs/bpf_rbtree_first.md)
    - [`bpf_rbtree_remove`](../kfuncs/bpf_rbtree_remove.md)
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md)
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md)
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md)
    - [`bpf_refcount_acquire_impl`](../kfuncs/bpf_refcount_acquire_impl.md)
    - [`bpf_task_acquire`](../kfuncs/bpf_task_acquire.md)
    - [`bpf_task_from_pid`](../kfuncs/bpf_task_from_pid.md)
    - [`bpf_task_get_cgroup1`](../kfuncs/bpf_task_get_cgroup1.md)
    - [`bpf_task_release`](../kfuncs/bpf_task_release.md)
    - [`bpf_task_under_cgroup`](../kfuncs/bpf_task_under_cgroup.md)
    - [`bpf_throw`](../kfuncs/bpf_throw.md)
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md)
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md)
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md)
    - [`crash_kexec`](../kfuncs/crash_kexec.md)
    - [`cubictcp_acked`](../kfuncs/cubictcp_acked.md)
    - [`cubictcp_cong_avoid`](../kfuncs/cubictcp_cong_avoid.md)
    - [`cubictcp_cwnd_event`](../kfuncs/cubictcp_cwnd_event.md)
    - [`cubictcp_init`](../kfuncs/cubictcp_init.md)
    - [`cubictcp_recalc_ssthresh`](../kfuncs/cubictcp_recalc_ssthresh.md)
    - [`cubictcp_state`](../kfuncs/cubictcp_state.md)
    - [`dctcp_cwnd_event`](../kfuncs/dctcp_cwnd_event.md)
    - [`dctcp_cwnd_undo`](../kfuncs/dctcp_cwnd_undo.md)
    - [`dctcp_init`](../kfuncs/dctcp_init.md)
    - [`dctcp_ssthresh`](../kfuncs/dctcp_ssthresh.md)
    - [`dctcp_state`](../kfuncs/dctcp_state.md)
    - [`dctcp_update_alpha`](../kfuncs/dctcp_update_alpha.md)
    - [`tcp_cong_avoid_ai`](../kfuncs/tcp_cong_avoid_ai.md)
    - [`tcp_reno_cong_avoid`](../kfuncs/tcp_reno_cong_avoid.md)
    - [`tcp_reno_ssthresh`](../kfuncs/tcp_reno_ssthresh.md)
    - [`tcp_reno_undo_cwnd`](../kfuncs/tcp_reno_undo_cwnd.md)
    - [`tcp_slow_start`](../kfuncs/tcp_slow_start.md)
<!-- [/PROG_KFUNC_REF] -->
