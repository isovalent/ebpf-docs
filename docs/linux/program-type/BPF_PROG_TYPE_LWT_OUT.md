---
title: "Program Type 'BPF_PROG_TYPE_LWT_OUT'"
description: "This page documents the 'BPF_PROG_TYPE_LWT_OUT' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_LWT_OUT`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_LWT_OUT) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/3a0af8fd61f90920f6fa04e4f1e9a6a73c1b4fd2)
<!-- [/FEATURE_TAG] -->

LWT (Light Weight Tunnel) Output programs attach to the egress path of a light weight tunnel.

## Usage

This program type can be attached to route, the program will be called for outgoing packets to said route. This is already at the point after the effects of the routing decision have been applied, thus this program type is purely for observation. For example

`ip route add 192.168.253.2/32 encap bpf out obj {elf file}.o section {prog section} dev veth0`

The initial use cases listed for this program type are:

- Collect statistics and generate sampling data for a subset of traffic based on the destination utilized by the packet thus allowing to extend the existing realms.
- Apply additional per route/destination filters to prohibit certain outgoing or incoming packets based on BPF filters. In particular, this allows to maintain per destination custom state across multiple packets in BPF maps and apply filters based on statistics and behaviour observed over time.

This output variant of the LWT programs can read the packet and make a PASS/DROP decision but it isn't allowed to modify the packet.

## Context

Socket SKB programs are called by the kernel with a [`__sk_buff`](../program-context/__sk_buff.md) context.

This program type isn't allowed to read from and write to all fields of the context since doing so might break assumptions in the kernel or because data isn't available at the point where the program is hooked into the kernel.

<!-- Information based on `lwt_is_valid_access` and `bpf_skb_is_valid_access` functions in the kernel sources -->

??? abstract "Context fields"
    | Field                                                                | Read             | Write            |
    | -------------------------------------------------------------------- | ---------------- | ---------------- |
    | [`len`](../program-context/__sk_buff.md#len)                         | :material-check: | :material-close: |
    | [`pkt_type`](../program-context/__sk_buff.md#pkt_type)               | :material-check: | :material-close: |
    | [`mark`](../program-context/__sk_buff.md#mark)                       | :material-check: | :material-check: |
    | [`queue_mapping`](../program-context/__sk_buff.md#queue_mapping)     | :material-check: | :material-close: |
    | [`protocol`](../program-context/__sk_buff.md#protocol)               | :material-check: | :material-close: |
    | [`vlan_present`](../program-context/__sk_buff.md#vlan_present)       | :material-check: | :material-close: |
    | [`vlan_tci`](../program-context/__sk_buff.md#vlan_tci)               | :material-check: | :material-close: |
    | [`vlan_proto`](../program-context/__sk_buff.md#vlan_proto)           | :material-check: | :material-close: |
    | [`priority`](../program-context/__sk_buff.md#priority)               | :material-check: | :material-check: |
    | [`ingress_ifindex`](../program-context/__sk_buff.md#ingress_ifindex) | :material-check: | :material-close: |
    | [`ifindex`](../program-context/__sk_buff.md#ifindex)                 | :material-check: | :material-close: |
    | [`tc_index`](../program-context/__sk_buff.md#tc_index)               | :material-check: | :material-close: |
    | [`cb`](../program-context/__sk_buff.md#cb)                           | :material-check: | :material-check: |
    | [`hash`](../program-context/__sk_buff.md#hash)                       | :material-check: | :material-close: |
    | [`tc_classid`](../program-context/__sk_buff.md#tc_classid)           | :material-close: | :material-close: |
    | [`data`](../program-context/__sk_buff.md#data)                       | :material-check: | :material-close: |
    | [`data_end`](../program-context/__sk_buff.md#data_end)               | :material-check: | :material-close: |
    | [`napi_id`](../program-context/__sk_buff.md#napi_id)                 | :material-check: | :material-close: |
    | [`family`](../program-context/__sk_buff.md#family)                   | :material-close: | :material-close: |
    | [`remote_ip4`](../program-context/__sk_buff.md#remote_ip4)           | :material-close: | :material-close: |
    | [`local_ip4`](../program-context/__sk_buff.md#local_ip4)             | :material-close: | :material-close: |
    | [`remote_ip4`](../program-context/__sk_buff.md#remote_ip4)           | :material-close: | :material-close: |
    | [`remote_ip6`](../program-context/__sk_buff.md#remote_ip6)           | :material-close: | :material-close: |
    | [`local_ip6`](../program-context/__sk_buff.md#local_ip6)             | :material-close: | :material-close: |
    | [`remote_port`](../program-context/__sk_buff.md#remote_port)         | :material-close: | :material-close: |
    | [`local_port`](../program-context/__sk_buff.md#local_port)           | :material-close: | :material-close: |
    | [`data_meta`](../program-context/__sk_buff.md#data_meta)             | :material-close: | :material-close: |
    | [`flow_keys`](../program-context/__sk_buff.md#flow_keys)             | :material-close: | :material-close: |
    | [`tstamp`](../program-context/__sk_buff.md#tstamp)                   | :material-close: | :material-close: |
    | [`wire_len`](../program-context/__sk_buff.md#wire_len)               | :material-close: | :material-close: |
    | [`gso_segs`](../program-context/__sk_buff.md#gso_segs)               | :material-check: | :material-close: |
    | [`sk`](../program-context/__sk_buff.md#sk)                           | :material-check: | :material-close: |
    | [`gso_size`](../program-context/__sk_buff.md#gso_size)               | :material-check: | :material-close: |
    | [`tstamp_type`](../program-context/__sk_buff.md#tstamp_type)         | :material-close: | :material-close: |
    | [`hwtstamp`](../program-context/__sk_buff.md#hwtstamp)               | :material-close: | :material-close: |

## Attachment

This program type can only be attached via netlink or commands such as `ip` from (iproute2) which use netlink under the hood:

`ip route add 192.168.253.2/32 encap bpf out obj {elf file}.o section {prog section} dev veth0`

## Example

!!! example
    ```c
    /* Copyright (c) 2016 Thomas Graf <tgraf@tgraf.ch>
    *
    * This program is free software; you can redistribute it and/or
    * modify it under the terms of version 2 of the GNU General Public
    * License as published by the Free Software Foundation.
    *
    * This program is distributed in the hope that it will be useful, but
    * WITHOUT ANY WARRANTY; without even the implied warranty of
    * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
    * General Public License for more details.
    */

    #include "vmlinux.h"
    #include <bpf/bpf_helpers.h>

    struct {
        __uint(type, BPF_MAP_TYPE_PERCPU_HASH);
        __type(key, u64);
        __type(value, u64);
        __uint(pinning, LIBBPF_PIN_BY_NAME);
        __uint(max_entries, 1024);
    } lwt_len_hist_map SEC(".maps");

    static unsigned int log2(unsigned int v)
    {
        unsigned int r;
        unsigned int shift;

        r = (v > 0xFFFF) << 4; v >>= r;
        shift = (v > 0xFF) << 3; v >>= shift; r |= shift;
        shift = (v > 0xF) << 2; v >>= shift; r |= shift;
        shift = (v > 0x3) << 1; v >>= shift; r |= shift;
        r |= (v >> 1);
        return r;
    }

    static unsigned int log2l(unsigned long v)
    {
        unsigned int hi = v >> 32;
        if (hi)
            return log2(hi) + 32;
        else
            return log2(v);
    }

    SEC("len_hist")
    int do_len_hist(struct __sk_buff *skb)
    {
        __u64 *value, key, init_val = 1;

        key = log2l(skb->len);

        value = bpf_map_lookup_elem(&lwt_len_hist_map, &key);
        if (value)
            __sync_fetch_and_add(value, 1);
        else
            bpf_map_update_elem(&lwt_len_hist_map, &key, &init_val, BPF_ANY);

        return BPF_OK;
    }

    char _license[] SEC("license") = "GPL";`
    ```

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for LWT programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_csum_diff`](../helper-function/bpf_csum_diff.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_get_cgroup_classid`](../helper-function/bpf_get_cgroup_classid.md)
    * [`bpf_get_current_pid_tgid`](../helper-function/bpf_get_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_get_hash_recalc`](../helper-function/bpf_get_hash_recalc.md)
    * [`bpf_get_ns_current_pid_tgid`](../helper-function/bpf_get_ns_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_route_realm`](../helper-function/bpf_get_route_realm.md)
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
    * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
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
    * [`bpf_skb_load_bytes`](../helper-function/bpf_skb_load_bytes.md)
    * [`bpf_skb_pull_data`](../helper-function/bpf_skb_pull_data.md)
    * [`bpf_skb_under_cgroup`](../helper-function/bpf_skb_under_cgroup.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
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
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md)
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md)
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md)
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md)
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md)
    - [`bpf_dynptr_from_skb`](../kfuncs/bpf_dynptr_from_skb.md)
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
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md)
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md)
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md)
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md)
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md)
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md)
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md)
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md)
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md)
<!-- [/PROG_KFUNC_REF] -->

