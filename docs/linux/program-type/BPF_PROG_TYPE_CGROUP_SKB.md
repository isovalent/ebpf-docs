---
title: "Program Type 'BPF_PROG_TYPE_CGROUP_SKB'"
description: "This page documents the 'BPF_PROG_TYPE_CGROUP_SKB' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_CGROUP_SKB`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_CGROUP_SKB) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/0e33661de493db325435d565a4a722120ae4cbf3)
<!-- [/FEATURE_TAG] -->

cGroup socket buffer programs are attached to a cGroup and are called for incoming or outgoing packets to or from processes within that cGroup. The programs can filter packets but not modify them.

## Usage

cGroup socket buffer programs can be used to filter packets and/or monitor/measure traffic on a per cGroup level.

cGroup socket buffer programs are typically placed in `cgroup_skb/egress` or `cgroup_skb/ingress` ELF sections depending on if they the program should be attached on the ingress or egress side, the respective `attach_type` value is then set to `BPF_CGROUP_INET_EGRESS` or `BPF_CGROUP_INET_INGRESS`.

Recognized return values are `0` to drop the packet or `1` to pass the packet, there are no enums for these return values. 

After [:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/5cf1e91456301f8c4f6bbc63ff76cff12f92f31b) `BPF_CGROUP_INET_EGRESS` programs can also set the 2's bit to indicate congestion occurred and signal to the higher level protocols that they should decrease rate. So `2` is drop + mark congestion and `3` is pass + mark congestion.

## Context

cGroup socket buffer programs are called by the kernel with a [`__sk_buff`](../program-context/__sk_buff.md) context.

This program type isn't allowed to read from and write to all fields of the context since doing so might break assumptions in the kernel or because data isn't available at the point where the program is hooked into the kernel.

<!-- Information based on `cg_skb_is_valid_access` and `bpf_skb_is_valid_access` functions in the kernel sources -->

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
    | [`family`](../program-context/__sk_buff.md#family)                   | :material-check: | :material-close: |
    | [`remote_ip4`](../program-context/__sk_buff.md#remote_ip4)           | :material-check: | :material-close: |
    | [`remote_ip6`](../program-context/__sk_buff.md#remote_ip6)           | :material-check: | :material-close: |
    | [`local_ip4`](../program-context/__sk_buff.md#local_ip4)             | :material-check: | :material-close: |
    | [`local_ip6`](../program-context/__sk_buff.md#local_ip6)             | :material-check: | :material-close: |
    | [`remote_port`](../program-context/__sk_buff.md#remote_port)         | :material-check: | :material-close: |
    | [`local_port`](../program-context/__sk_buff.md#local_port)           | :material-check: | :material-close: |
    | [`data_meta`](../program-context/__sk_buff.md#data_meta)             | :material-close: | :material-close: |
    | [`flow_keys`](../program-context/__sk_buff.md#flow_keys)             | :material-close: | :material-close: |
    | [`tstamp`](../program-context/__sk_buff.md#tstamp)                   | :material-check: | :material-check: |
    | [`wire_len`](../program-context/__sk_buff.md#wire_len)               | :material-close: | :material-close: |
    | [`tstamp`](../program-context/__sk_buff.md#tstamp)                   | :material-check: | :material-close: |
    | [`gso_segs`](../program-context/__sk_buff.md#gso_segs)               | :material-check: | :material-close: |
    | [`sk`](../program-context/__sk_buff.md#sk)                           | :material-check: | :material-close: |
    | [`gso_size`](../program-context/__sk_buff.md#gso_size)               | :material-check: | :material-close: |
    | [`tstamp_type`](../program-context/__sk_buff.md#tstamp_type)         | :material-close: | :material-close: |
    | [`hwtstamp`](../program-context/__sk_buff.md#hwtstamp)               | :material-check: | :material-close: |

## Attachment

cGroup socket buffer programs are attached to cGroups via the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall or via [BPF link](../syscall/BPF_LINK_CREATE.md).

## Example

Example BPF program:

```c
/* Copyright (c) 2019 Facebook
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of version 2 of the GNU General Public
 * License as published by the Free Software Foundation.
 *
 * Sample Host Bandwidth Manager (HBM) BPF program.
 *
 * A cgroup skb BPF egress program to limit cgroup output bandwidth.
 * It uses a modified virtual token bucket queue to limit average
 * egress bandwidth. The implementation uses credits instead of tokens.
 * Negative credits imply that queueing would have happened (this is
 * a virtual queue, so no queueing is done by it. However, queueing may
 * occur at the actual qdisc (which is not used for rate limiting).
 *
 * This implementation uses 3 thresholds, one to start marking packets and
 * the other two to drop packets:
 *                                  CREDIT
 *        - <--------------------------|------------------------> +
 *              |    |          |      0
 *              |  Large pkt    |
 *              |  drop thresh  |
 *   Small pkt drop             Mark threshold
 *       thresh
 *
 * The effect of marking depends on the type of packet:
 * a) If the packet is ECN enabled and it is a TCP packet, then the packet
 *    is ECN marked.
 * b) If the packet is a TCP packet, then we probabilistically call tcp_cwr
 *    to reduce the congestion window. The current implementation uses a linear
 *    distribution (0% probability at marking threshold, 100% probability
 *    at drop threshold).
 * c) If the packet is not a TCP packet, then it is dropped.
 *
 * If the credit is below the drop threshold, the packet is dropped. If it
 * is a TCP packet, then it also calls tcp_cwr since packets dropped by
 * a cgroup skb BPF program do not automatically trigger a call to
 * tcp_cwr in the current kernel code.
 *
 * This BPF program actually uses 2 drop thresholds, one threshold
 * for larger packets (>= 120 bytes) and another for smaller packets. This
 * protects smaller packets such as SYNs, ACKs, etc.
 *
 * The default bandwidth limit is set at 1Gbps but this can be changed by
 * a user program through a shared BPF map. In addition, by default this BPF
 * program does not limit connections using loopback. This behavior can be
 * overwritten by the user program. There is also an option to calculate
 * some statistics, such as percent of packets marked or dropped, which
 * a user program, such as hbm, can access.
 */

#include "hbm_kern.h"

SEC("cgroup_skb/egress")
int _hbm_out_cg(struct __sk_buff *skb)
{
	long long delta = 0, delta_send;
	unsigned long long curtime, sendtime;
	struct hbm_queue_stats *qsp = NULL;
	unsigned int queue_index = 0;
	bool congestion_flag = false;
	bool ecn_ce_flag = false;
	struct hbm_pkt_info pkti = {};
	struct hbm_vqueue *qdp;
	bool drop_flag = false;
	bool cwr_flag = false;
	int len = skb->len;
	int rv = ALLOW_PKT;

	qsp = bpf_map_lookup_elem(&queue_stats, &queue_index);

	// Check if we should ignore loopback traffic
	if (qsp != NULL && !qsp->loopback && (skb->ifindex == 1))
		return ALLOW_PKT;

	hbm_get_pkt_info(skb, &pkti);

	// We may want to account for the length of headers in len
	// calculation, like ETH header + overhead, specially if it
	// is a gso packet. But I am not doing it right now.

	qdp = bpf_get_local_storage(&queue_state, 0);
	if (!qdp)
		return ALLOW_PKT;
	if (qdp->lasttime == 0)
		hbm_init_edt_vqueue(qdp, 1024);

	curtime = bpf_ktime_get_ns();

	// Begin critical section
	bpf_spin_lock(&qdp->lock);
	delta = qdp->lasttime - curtime;
	// bound bursts to 100us
	if (delta < -BURST_SIZE_NS) {
		// negative delta is a credit that allows bursts
		qdp->lasttime = curtime - BURST_SIZE_NS;
		delta = -BURST_SIZE_NS;
	}
	sendtime = qdp->lasttime;
	delta_send = BYTES_TO_NS(len, qdp->rate);
	__sync_add_and_fetch(&(qdp->lasttime), delta_send);
	bpf_spin_unlock(&qdp->lock);
	// End critical section

	// Set EDT of packet
	skb->tstamp = sendtime;

	// Check if we should update rate
	if (qsp != NULL && (qsp->rate * 128) != qdp->rate)
		qdp->rate = qsp->rate * 128;

	// Set flags (drop, congestion, cwr)
	// last packet will be sent in the future, bound latency
	if (delta > DROP_THRESH_NS || (delta > LARGE_PKT_DROP_THRESH_NS &&
				       len > LARGE_PKT_THRESH)) {
		drop_flag = true;
		if (pkti.is_tcp && pkti.ecn == 0)
			cwr_flag = true;
	} else if (delta > MARK_THRESH_NS) {
		if (pkti.is_tcp)
			congestion_flag = true;
		else
			drop_flag = true;
	}

	if (congestion_flag) {
		if (bpf_skb_ecn_set_ce(skb)) {
			ecn_ce_flag = true;
		} else {
			if (pkti.is_tcp) {
				unsigned int rand = bpf_get_prandom_u32();

				if (delta >= MARK_THRESH_NS +
				    (rand % MARK_REGION_SIZE_NS)) {
					// Do congestion control
					cwr_flag = true;
				}
			} else if (len > LARGE_PKT_THRESH) {
				// Problem if too many small packets?
				drop_flag = true;
				congestion_flag = false;
			}
		}
	}

	if (pkti.is_tcp && drop_flag && pkti.packets_out <= 1) {
		drop_flag = false;
		cwr_flag = true;
		congestion_flag = false;
	}

	if (qsp != NULL && qsp->no_cn)
			cwr_flag = false;

	hbm_update_stats(qsp, len, curtime, congestion_flag, drop_flag,
			 cwr_flag, ecn_ce_flag, &pkti, (int) delta);

	if (drop_flag) {
		__sync_add_and_fetch(&(qdp->lasttime), -delta_send);
		rv = DROP_PKT;
	}

	if (cwr_flag)
		rv |= CWR;
	return rv;
}
char _license[] SEC("license") = "GPL";
```

## Helper functions

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
    * [`bpf_get_listener_sock`](../helper-function/bpf_get_listener_sock.md)
    * [`bpf_get_local_storage`](../helper-function/bpf_get_local_storage.md)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_get_socket_cookie`](../helper-function/bpf_get_socket_cookie.md)
    * [`bpf_get_socket_uid`](../helper-function/bpf_get_socket_uid.md)
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
    * [`bpf_sk_ancestor_cgroup_id`](../helper-function/bpf_sk_ancestor_cgroup_id.md)
    * [`bpf_sk_cgroup_id`](../helper-function/bpf_sk_cgroup_id.md)
    * [`bpf_sk_fullsock`](../helper-function/bpf_sk_fullsock.md)
    * [`bpf_sk_lookup_tcp`](../helper-function/bpf_sk_lookup_tcp.md)
    * [`bpf_sk_lookup_udp`](../helper-function/bpf_sk_lookup_udp.md)
    * [`bpf_sk_release`](../helper-function/bpf_sk_release.md)
    * [`bpf_sk_storage_delete`](../helper-function/bpf_sk_storage_delete.md)
    * [`bpf_sk_storage_get`](../helper-function/bpf_sk_storage_get.md)
    * [`bpf_skb_ancestor_cgroup_id`](../helper-function/bpf_skb_ancestor_cgroup_id.md)
    * [`bpf_skb_cgroup_id`](../helper-function/bpf_skb_cgroup_id.md)
    * [`bpf_skb_ecn_set_ce`](../helper-function/bpf_skb_ecn_set_ce.md)
    * [`bpf_skb_load_bytes`](../helper-function/bpf_skb_load_bytes.md)
    * [`bpf_skb_load_bytes_relative`](../helper-function/bpf_skb_load_bytes_relative.md)
    * [`bpf_skc_lookup_tcp`](../helper-function/bpf_skc_lookup_tcp.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_tcp_sock`](../helper-function/bpf_tcp_sock.md)
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
    - [`bpf_sock_addr_set_sun_path`](../kfuncs/bpf_sock_addr_set_sun_path.md)
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md)
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md)
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md)
<!-- [/PROG_KFUNC_REF] -->
