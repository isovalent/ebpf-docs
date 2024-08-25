---
title: "Program Type 'BPF_PROG_TYPE_NETFILTER'"
description: "This page documents the 'BPF_PROG_TYPE_NETFILTER' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_NETFILTER`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_NETFILTER) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/84601d6ee68ae820dec97450934797046d62db4b)
<!-- [/FEATURE_TAG] -->

## Usage

This program type is used to implement a netfilter (aka `iptables` / `nftables`) hook in eBPF. 

The hook can make a decision to drop or accept the packet by returning `NF_DROP` (0) or `NF_ACCEPT` (1) respectively.

## Context

The context that is passed in contains pointers to the hook state and to a full `sk_buff` as opposed to the `__sk_buff` we typically see as the context in other program types. The whole ctx is read-only.

```c
struct bpf_nf_ctx {
	const struct nf_hook_state *state;
	struct sk_buff *skb;
};
```

The `ctx->skb` pointer can be used in combination with the [`bpf_dynptr_from_skb`](../kfuncs/bpf_dynptr_from_skb.md) kfunc to access the packet data. The returned dynptr will be read-only.

The hook state contains a lot of information about the current hook and state of the packet.

```c
struct nf_hook_state {
	u8 hook;
	u8 pf;
	struct net_device *in;
	struct net_device *out;
	struct sock *sk;
	struct net *net;
	int (*okfn)(struct net *, struct sock *, struct sk_buff *);
};
```

## Attachment

These programs are attached via the link API. The netlink portion of the link create attributes look like:

```c
struct {
    __u32		pf;
    __u32		hooknum;
    __s32		priority;
    __u32		flags;
} netfilter;
```

`pf` is the protocol family, supported values are `NFPROTO_IPV4` (2) and `NFPROTO_IPV6` (10).

`hooknum` is the hook number, supported values are `NF_INET_PRE_ROUTING` (0), `NF_INET_LOCAL_IN` (1), `NF_INET_FORWARD` (2), `NF_INET_LOCAL_OUT` (3), and `NF_INET_POST_ROUTING` (4).

`priority` is the priority of the hook, lower values are called first. `NF_IP_PRI_FIRST` (-2147483648) and `NF_IP_PRI_LAST` (2147483647) are not allowed. If the `BPF_F_NETFILTER_IP_DEFRAG` flag is set, the priority must be higher than `NF_IP_PRI_CONNTRACK_DEFRAG` (-400).

`flags` is a bitmask of flags. Supported flags are:

* `NF_IP_PRI_CONNTRACK_DEFRAG` - Enable defragmentation of IP fragments, this hook will only see defragmented packets.

## Example

```c
// SPDX-License-Identifier: GPL-2.0-only
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>
#include "bpf_tracing_net.h"

#define NF_DROP			0
#define NF_ACCEPT		1
#define ETH_P_IP		0x0800
#define ETH_P_IPV6		0x86DD
#define IP_MF			0x2000
#define IP_OFFSET		0x1FFF
#define NEXTHDR_FRAGMENT	44

extern int bpf_dynptr_from_skb(struct __sk_buff *skb, __u64 flags,
			      struct bpf_dynptr *ptr__uninit) __ksym;
extern void *bpf_dynptr_slice(const struct bpf_dynptr *ptr, uint32_t offset,
			      void *buffer, uint32_t buffer__sz) __ksym;

volatile int shootdowns = 0;

static bool is_frag_v4(struct iphdr *iph)
{
	int offset;
	int flags;

	offset = bpf_ntohs(iph->frag_off);
	flags = offset & ~IP_OFFSET;
	offset &= IP_OFFSET;
	offset <<= 3;

	return (flags & IP_MF) || offset;
}

static bool is_frag_v6(struct ipv6hdr *ip6h)
{
	/* Simplifying assumption that there are no extension headers
	 * between fixed header and fragmentation header. This assumption
	 * is only valid in this test case. It saves us the hassle of
	 * searching all potential extension headers.
	 */
	return ip6h->nexthdr == NEXTHDR_FRAGMENT;
}

static int handle_v4(struct __sk_buff *skb)
{
	struct bpf_dynptr ptr;
	u8 iph_buf[20] = {};
	struct iphdr *iph;

	if (bpf_dynptr_from_skb(skb, 0, &ptr))
		return NF_DROP;

	iph = bpf_dynptr_slice(&ptr, 0, iph_buf, sizeof(iph_buf));
	if (!iph)
		return NF_DROP;

	/* Shootdown any frags */
	if (is_frag_v4(iph)) {
		shootdowns++;
		return NF_DROP;
	}

	return NF_ACCEPT;
}

static int handle_v6(struct __sk_buff *skb)
{
	struct bpf_dynptr ptr;
	struct ipv6hdr *ip6h;
	u8 ip6h_buf[40] = {};

	if (bpf_dynptr_from_skb(skb, 0, &ptr))
		return NF_DROP;

	ip6h = bpf_dynptr_slice(&ptr, 0, ip6h_buf, sizeof(ip6h_buf));
	if (!ip6h)
		return NF_DROP;

	/* Shootdown any frags */
	if (is_frag_v6(ip6h)) {
		shootdowns++;
		return NF_DROP;
	}

	return NF_ACCEPT;
}

SEC("netfilter")
int defrag(struct bpf_nf_ctx *ctx)
{
	struct __sk_buff *skb = (struct __sk_buff *)ctx->skb;

	switch (bpf_ntohs(ctx->skb->protocol)) {
	case ETH_P_IP:
		return handle_v4(skb);
	case ETH_P_IPV6:
		return handle_v6(skb);
	default:
		return NF_ACCEPT;
	}
}

char _license[] SEC("license") = "GPL";
```

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for `BPF_PROG_TYPE_NETFILTER` programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
    * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
    * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
    * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
    * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
    * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
    * [`bpf_map_lookup_percpu_elem`](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
    * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
    * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
    * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
    * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
    * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_loop`](../helper-function/bpf_loop.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_per_cpu_ptr`](../helper-function/bpf_per_cpu_ptr.md)
    * [`bpf_this_cpu_ptr`](../helper-function/bpf_this_cpu_ptr.md)
    * [`bpf_timer_init`](../helper-function/bpf_timer_init.md)
    * [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md)
    * [`bpf_timer_start`](../helper-function/bpf_timer_start.md)
    * [`bpf_timer_cancel`](../helper-function/bpf_timer_cancel.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)
    * [`bpf_ktime_get_tai_ns`](../helper-function/bpf_ktime_get_tai_ns.md)
    * [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md)
    * [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md)
    * [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md)
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
