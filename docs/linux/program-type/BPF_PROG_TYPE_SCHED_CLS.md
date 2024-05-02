---
title: "Program Type 'BPF_PROG_TYPE_SCHED_CLS'"
description: "This page documents the 'BPF_PROG_TYPE_SCHED_CLS' eBPF program type, including its defintion, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_SCHED_CLS`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_SCHED_CLS) -->
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/96be4325f443dbbfeb37d2a157675ac0736531a1)
<!-- [/FEATURE_TAG] -->

This program type allows for the implementation of a Traffic Control (TC) classifier (aka filter) in eBPF. TC can be used for a number of use cases, all of them having to do with the manipulation of traffic. TC is for example used to implement QoS (Quality of Service) allowing latency sensitive traffic like VoIP (Voice over IP) to be processed ahead of lets say web traffic. It can also drop packets to simulate packet-loss, add latency to simulate distant clients or apply bandwidth limitations for applications or users, to name a few.

TC allows an admin to filter traffic using a hierarchical model of qdisc's (Queuing DISCipline). A root qdisc is attached to a network interface with certain actions. This qdisc can also have child qdiscs which will be used over the root if their filter matches the traffic. This program type allows us to implement such a filter in eBPF.

## Usage

TC CLS programs are typically put into an [ELF](../../concepts/elf.md) section prefixed with `tc/` or `classifier/`. The TC CLS program is called by the kernel with a [__sk_buff](../program-context/__sk_buff.md) context. The return value indicates what action the kernel should take with the packet, the following values are permitted:

### Regular classifier

By default, when a BPF classifier is attached to a qdisc it will act as any other classifier. It can't take actions such as dropping or redirecting packets, instead its return value is used to pick a class based on the contents of the packet. A return value of `-1` indicates the default class should be picked, a return value of `0` means the filter did not match, and that the next filter should try, and any positive number indicates the id of the class.

While possible, this is a rarely used use-case, eBPF programs are typically used for direct actions.

### Direct action

When attached in direct action mode, the eBPF program will act as both a classifier and an action. This mode simplifies setups for the most common use cases where we just want to always execute an action. In direct action mode the return value can be one of:

* `TC_ACT_UNSPEC` (-1) - Signals that the default configured action should be taken.
* `TC_ACT_OK` (0) - Signals that the packet should proceed.
* `TC_ACT_RECLASSIFY` (1) - Signals that the packet has to re-start classification from the root qdisc. This is typically used after modifying the packet so its classification might have different results.
* `TC_ACT_SHOT` (2) - Signals that the packet should be dropped, no other TC processing should happen.
* `TC_ACT_PIPE`	(3) - While defined, this action should not be used and holds no particular meaning for eBPF classifiers.
* `TC_ACT_STOLEN` (4) - While defined, this action should not be used and holds no particular meaning for eBPF classifiers.
* `TC_ACT_QUEUED` (5) - While defined, this action should not be used and holds no particular meaning for eBPF classifiers.
* `TC_ACT_REPEAT` (6) - While defined, this action should not be used and holds no particular meaning for eBPF classifiers.
* `TC_ACT_REDIRECT`	(7) - Signals that the packet should be redirected, the details of how and where to are set as side effects by [helpers functions](../helper-function/index.md).

Classifiers in direct action mode can still set a class id by setting the `tc_classid` field 

## Context

This program type is not allowed to read from and write to all fields of the context since doing so might break assumptions in the kernel or because data is not available at the point where the program is hooked into the kernel.

<!-- Information based on `tc_cls_act_is_valid_access` and `bpf_skb_is_valid_access` functions in the kernel sources -->

??? abstract "Context fields"
    | Field                                                                | Read             | Write            |
    | -------------------------------------------------------------------- | ---------------- | ---------------- |
    | [`len`](../program-context/__sk_buff.md#len)                         | :material-check: | :material-close: |
    | [`pkt_type`](../program-context/__sk_buff.md#pkt_type)               | :material-check: | :material-close: |
    | [`mark`](../program-context/__sk_buff.md#mark)                       | :material-check: | :material-check: |
    | [`queue_mapping`](../program-context/__sk_buff.md#queue_mapping)     | :material-check: | :material-check: |
    | [`protocol`](../program-context/__sk_buff.md#protocol)               | :material-check: | :material-close: |
    | [`vlan_present`](../program-context/__sk_buff.md#vlan_present)       | :material-check: | :material-close: |
    | [`vlan_tci`](../program-context/__sk_buff.md#vlan_tci)               | :material-check: | :material-close: |
    | [`vlan_proto`](../program-context/__sk_buff.md#vlan_proto)           | :material-check: | :material-close: |
    | [`priority`](../program-context/__sk_buff.md#priority)               | :material-check: | :material-check: |
    | [`ingress_ifindex`](../program-context/__sk_buff.md#ingress_ifindex) | :material-check: | :material-close: |
    | [`ifindex`](../program-context/__sk_buff.md#ifindex)                 | :material-check: | :material-close: |
    | [`tc_index`](../program-context/__sk_buff.md#tc_index)               | :material-check: | :material-check: |
    | [`cb`](../program-context/__sk_buff.md#cb)                           | :material-check: | :material-check: |
    | [`hash`](../program-context/__sk_buff.md#hash)                       | :material-check: | :material-close: |
    | [`tc_classid`](../program-context/__sk_buff.md#tc_classid)           | :material-check: | :material-check: |
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
    | [`data_meta`](../program-context/__sk_buff.md#data_meta)             | :material-check: | :material-close: |
    | [`flow_keys`](../program-context/__sk_buff.md#flow_keys)             | :material-close: | :material-close: |
    | [`tstamp`](../program-context/__sk_buff.md#tstamp)                   | :material-check: | :material-check: |
    | [`wire_len`](../program-context/__sk_buff.md#wire_len)               | :material-check: | :material-close: |
    | [`tstamp`](../program-context/__sk_buff.md#tstamp)                   | :material-check: | :material-close: |
    | [`gso_segs`](../program-context/__sk_buff.md#gso_segs)               | :material-check: | :material-close: |
    | [`sk`](../program-context/__sk_buff.md#sk)                           | :material-check: | :material-close: |
    | [`gso_size`](../program-context/__sk_buff.md#gso_size)               | :material-check: | :material-close: |
    | [`tstamp_type`](../program-context/__sk_buff.md#tstamp_type)         | :material-check: | :material-close: |
    | [`hwtstamp`](../program-context/__sk_buff.md#hwtstamp)               | :material-check: | :material-close: |

## Attachment

As of kernel version v6.2 the only way to attach eBPF programs to TC is via a [netlink](https://man7.org/linux/man-pages/man7/netlink.7.html) socket the details of which are complex. The usage of a netlink library is recommended if you wish to manage attachment via an API. However, the most common way to go about this is via the iproute2 `tc` CLI tool which is the standard implementation for network utilities using the netlink protocol.

The most basic example of attaching a TC classifier is:

```bash
# Add a qdisc of type `clsact` to device `eth1`
$ tc qdisc add dev eth1 clsact
# Load the `program.o` ELF file, and attach the `my_func` section to the qdisc of eth1 on the ingress side.
$ tc filter add dev eth1 ingress bpf obj program.o sec my_func
```

For more details on the tc command, see the general [man page](https://man7.org/linux/man-pages/man8/tc.8.html).

For more details on the bpf filter options, see the tc-bpf [man page](https://man7.org/linux/man-pages/man8/tc-bpf.8.html).

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for TC classifier programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [bpf_skb_store_bytes](../helper-function/bpf_skb_store_bytes.md)
    * [bpf_skb_load_bytes](../helper-function/bpf_skb_load_bytes.md)
    * [bpf_skb_load_bytes_relative](../helper-function/bpf_skb_load_bytes_relative.md)
    * [bpf_skb_pull_data](../helper-function/bpf_skb_pull_data.md)
    * [bpf_csum_diff](../helper-function/bpf_csum_diff.md)
    * [bpf_csum_update](../helper-function/bpf_csum_update.md)
    * [bpf_csum_level](../helper-function/bpf_csum_level.md)
    * [bpf_l3_csum_replace](../helper-function/bpf_l3_csum_replace.md)
    * [bpf_l4_csum_replace](../helper-function/bpf_l4_csum_replace.md)
    * [bpf_clone_redirect](../helper-function/bpf_clone_redirect.md)
    * [bpf_get_cgroup_classid](../helper-function/bpf_get_cgroup_classid.md)
    * [bpf_skb_vlan_push](../helper-function/bpf_skb_vlan_push.md)
    * [bpf_skb_vlan_pop](../helper-function/bpf_skb_vlan_pop.md)
    * [bpf_skb_change_proto](../helper-function/bpf_skb_change_proto.md)
    * [bpf_skb_change_type](../helper-function/bpf_skb_change_type.md)
    * [bpf_skb_adjust_room](../helper-function/bpf_skb_adjust_room.md)
    * [bpf_skb_change_tail](../helper-function/bpf_skb_change_tail.md)
    * [bpf_skb_change_head](../helper-function/bpf_skb_change_head.md) [:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/6f3f65d80dac8f2bafce2213005821fccdce194c)
    * [bpf_skb_get_tunnel_key](../helper-function/bpf_skb_get_tunnel_key.md)
    * [bpf_skb_set_tunnel_key](../helper-function/bpf_skb_set_tunnel_key.md)
    * [bpf_skb_get_tunnel_opt](../helper-function/bpf_skb_get_tunnel_opt.md)
    * [bpf_skb_set_tunnel_opt](../helper-function/bpf_skb_set_tunnel_opt.md)
    * [bpf_redirect](../helper-function/bpf_redirect.md)
    * [bpf_redirect_neigh](../helper-function/bpf_redirect_neigh.md)
    * [bpf_redirect_peer](../helper-function/bpf_redirect_peer.md)
    * [bpf_get_route_realm](../helper-function/bpf_get_route_realm.md)
    * [bpf_get_hash_recalc](../helper-function/bpf_get_hash_recalc.md)
    * [bpf_set_hash_invalid](../helper-function/bpf_set_hash_invalid.md)
    * [bpf_set_hash](../helper-function/bpf_set_hash.md)
    * [bpf_perf_event_output](../helper-function/bpf_perf_event_output.md)
    * [bpf_get_smp_processor_id](../helper-function/bpf_get_smp_processor_id.md)
    * [bpf_skb_under_cgroup](../helper-function/bpf_skb_under_cgroup.md)
    * [bpf_get_socket_cookie](../helper-function/bpf_get_socket_cookie.md)
    * [bpf_get_socket_uid](../helper-function/bpf_get_socket_uid.md)
    * [bpf_fib_lookup](../helper-function/bpf_fib_lookup.md)
    * [bpf_check_mtu](../helper-function/bpf_check_mtu.md)
    * [bpf_sk_fullsock](../helper-function/bpf_sk_fullsock.md)
    * [bpf_sk_storage_get](../helper-function/bpf_sk_storage_get.md)
    * [bpf_sk_storage_delete](../helper-function/bpf_sk_storage_delete.md)
    * [bpf_skb_get_xfrm_state](../helper-function/bpf_skb_get_xfrm_state.md)
    * [bpf_skb_cgroup_classid](../helper-function/bpf_skb_cgroup_classid.md)
    * [bpf_skb_cgroup_id](../helper-function/bpf_skb_cgroup_id.md)
    * [bpf_skb_ancestor_cgroup_id](../helper-function/bpf_skb_ancestor_cgroup_id.md)
    * [bpf_sk_lookup_tcp](../helper-function/bpf_sk_lookup_tcp.md)
    * [bpf_sk_lookup_udp](../helper-function/bpf_sk_lookup_udp.md)
    * [bpf_sk_release](../helper-function/bpf_sk_release.md)
    * [bpf_tcp_sock](../helper-function/bpf_tcp_sock.md)
    * [bpf_get_listener_sock](../helper-function/bpf_get_listener_sock.md)
    * [bpf_skc_lookup_tcp](../helper-function/bpf_skc_lookup_tcp.md)
    * [bpf_tcp_check_syncookie](../helper-function/bpf_tcp_check_syncookie.md)
    * [bpf_skb_ecn_set_ce](../helper-function/bpf_skb_ecn_set_ce.md)
    * [bpf_tcp_gen_syncookie](../helper-function/bpf_tcp_gen_syncookie.md)
    * [bpf_sk_assign](../helper-function/bpf_sk_assign.md)
    * [bpf_skb_set_tstamp](../helper-function/bpf_skb_set_tstamp.md)
    * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
    * [bpf_map_update_elem](../helper-function/bpf_map_update_elem.md)
    * [bpf_map_delete_elem](../helper-function/bpf_map_delete_elem.md)
    * [bpf_map_push_elem](../helper-function/bpf_map_push_elem.md)
    * [bpf_map_pop_elem](../helper-function/bpf_map_pop_elem.md)
    * [bpf_map_peek_elem](../helper-function/bpf_map_peek_elem.md)
    * [bpf_map_lookup_percpu_elem](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [bpf_get_prandom_u32](../helper-function/bpf_get_prandom_u32.md)
    * [bpf_get_smp_processor_id](../helper-function/bpf_get_smp_processor_id.md)
    * [bpf_get_numa_node_id](../helper-function/bpf_get_numa_node_id.md)
    * [bpf_tail_call](../helper-function/bpf_tail_call.md)
    * [bpf_ktime_get_ns](../helper-function/bpf_ktime_get_ns.md)
    * [bpf_ktime_get_boot_ns](../helper-function/bpf_ktime_get_boot_ns.md)
    * [bpf_ringbuf_output](../helper-function/bpf_ringbuf_output.md)
    * [bpf_ringbuf_reserve](../helper-function/bpf_ringbuf_reserve.md)
    * [bpf_ringbuf_submit](../helper-function/bpf_ringbuf_submit.md)
    * [bpf_ringbuf_discard](../helper-function/bpf_ringbuf_discard.md)
    * [bpf_ringbuf_query](../helper-function/bpf_ringbuf_query.md)
    * [bpf_for_each_map_elem](../helper-function/bpf_for_each_map_elem.md)
    * [bpf_loop](../helper-function/bpf_loop.md)
    * [bpf_strncmp](../helper-function/bpf_strncmp.md)
    * [bpf_spin_lock](../helper-function/bpf_spin_lock.md)
    * [bpf_spin_unlock](../helper-function/bpf_spin_unlock.md)
    * [bpf_jiffies64](../helper-function/bpf_jiffies64.md)
    * [bpf_per_cpu_ptr](../helper-function/bpf_per_cpu_ptr.md)
    * [bpf_this_cpu_ptr](../helper-function/bpf_this_cpu_ptr.md)
    * [bpf_timer_init](../helper-function/bpf_timer_init.md)
    * [bpf_timer_set_callback](../helper-function/bpf_timer_set_callback.md)
    * [bpf_timer_start](../helper-function/bpf_timer_start.md)
    * [bpf_timer_cancel](../helper-function/bpf_timer_cancel.md)
    * [bpf_trace_printk](../helper-function/bpf_trace_printk.md)
    * [bpf_get_current_task](../helper-function/bpf_get_current_task.md)
    * [bpf_get_current_task_btf](../helper-function/bpf_get_current_task_btf.md)
    * [bpf_probe_read_user](../helper-function/bpf_probe_read_user.md)
    * [bpf_probe_read_kernel](../helper-function/bpf_probe_read_kernel.md)
    * [bpf_probe_read_user_str](../helper-function/bpf_probe_read_user_str.md)
    * [bpf_probe_read_kernel_str](../helper-function/bpf_probe_read_kernel_str.md)
    * [bpf_snprintf_btf](../helper-function/bpf_snprintf_btf.md)
    * [bpf_snprintf](../helper-function/bpf_snprintf.md)
    * [bpf_task_pt_regs](../helper-function/bpf_task_pt_regs.md)
    * [bpf_trace_vprintk](../helper-function/bpf_trace_vprintk.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
??? abstract "Supported kfuncs"
    - [bpf_cast_to_kern_ctx](../kfuncs/bpf_cast_to_kern_ctx.md)
    - [bpf_cgroup_acquire](../kfuncs/bpf_cgroup_acquire.md)
    - [bpf_cgroup_ancestor](../kfuncs/bpf_cgroup_ancestor.md)
    - [bpf_cgroup_from_id](../kfuncs/bpf_cgroup_from_id.md)
    - [bpf_cgroup_release](../kfuncs/bpf_cgroup_release.md)
    - [bpf_ct_change_status](../kfuncs/bpf_ct_change_status.md)
    - [bpf_ct_change_timeout](../kfuncs/bpf_ct_change_timeout.md)
    - [bpf_ct_insert_entry](../kfuncs/bpf_ct_insert_entry.md)
    - [bpf_ct_release](../kfuncs/bpf_ct_release.md)
    - [bpf_ct_set_nat_info](../kfuncs/bpf_ct_set_nat_info.md)
    - [bpf_ct_set_status](../kfuncs/bpf_ct_set_status.md)
    - [bpf_ct_set_timeout](../kfuncs/bpf_ct_set_timeout.md)
    - [bpf_dynptr_adjust](../kfuncs/bpf_dynptr_adjust.md)
    - [bpf_dynptr_clone](../kfuncs/bpf_dynptr_clone.md)
    - [bpf_dynptr_from_skb](../kfuncs/bpf_dynptr_from_skb.md)
    - [bpf_dynptr_is_null](../kfuncs/bpf_dynptr_is_null.md)
    - [bpf_dynptr_is_rdonly](../kfuncs/bpf_dynptr_is_rdonly.md)
    - [bpf_dynptr_size](../kfuncs/bpf_dynptr_size.md)
    - [bpf_dynptr_slice](../kfuncs/bpf_dynptr_slice.md)
    - [bpf_dynptr_slice_rdwr](../kfuncs/bpf_dynptr_slice_rdwr.md)
    - [bpf_iter_css_destroy](../kfuncs/bpf_iter_css_destroy.md)
    - [bpf_iter_css_new](../kfuncs/bpf_iter_css_new.md)
    - [bpf_iter_css_next](../kfuncs/bpf_iter_css_next.md)
    - [bpf_iter_css_task_destroy](../kfuncs/bpf_iter_css_task_destroy.md)
    - [bpf_iter_css_task_new](../kfuncs/bpf_iter_css_task_new.md)
    - [bpf_iter_css_task_next](../kfuncs/bpf_iter_css_task_next.md)
    - [bpf_iter_num_destroy](../kfuncs/bpf_iter_num_destroy.md)
    - [bpf_iter_num_new](../kfuncs/bpf_iter_num_new.md)
    - [bpf_iter_num_next](../kfuncs/bpf_iter_num_next.md)
    - [bpf_iter_task_destroy](../kfuncs/bpf_iter_task_destroy.md)
    - [bpf_iter_task_new](../kfuncs/bpf_iter_task_new.md)
    - [bpf_iter_task_next](../kfuncs/bpf_iter_task_next.md)
    - [bpf_iter_task_vma_destroy](../kfuncs/bpf_iter_task_vma_destroy.md)
    - [bpf_iter_task_vma_new](../kfuncs/bpf_iter_task_vma_new.md)
    - [bpf_iter_task_vma_next](../kfuncs/bpf_iter_task_vma_next.md)
    - [bpf_list_pop_back](../kfuncs/bpf_list_pop_back.md)
    - [bpf_list_pop_front](../kfuncs/bpf_list_pop_front.md)
    - [bpf_list_push_back_impl](../kfuncs/bpf_list_push_back_impl.md)
    - [bpf_list_push_front_impl](../kfuncs/bpf_list_push_front_impl.md)
    - [bpf_map_sum_elem_count](../kfuncs/bpf_map_sum_elem_count.md)
    - [bpf_obj_drop_impl](../kfuncs/bpf_obj_drop_impl.md)
    - [bpf_obj_new_impl](../kfuncs/bpf_obj_new_impl.md)
    - [bpf_percpu_obj_drop_impl](../kfuncs/bpf_percpu_obj_drop_impl.md)
    - [bpf_percpu_obj_new_impl](../kfuncs/bpf_percpu_obj_new_impl.md)
    - [bpf_rbtree_add_impl](../kfuncs/bpf_rbtree_add_impl.md)
    - [bpf_rbtree_first](../kfuncs/bpf_rbtree_first.md)
    - [bpf_rbtree_remove](../kfuncs/bpf_rbtree_remove.md)
    - [bpf_rcu_read_lock](../kfuncs/bpf_rcu_read_lock.md)
    - [bpf_rcu_read_unlock](../kfuncs/bpf_rcu_read_unlock.md)
    - [bpf_rdonly_cast](../kfuncs/bpf_rdonly_cast.md)
    - [bpf_refcount_acquire_impl](../kfuncs/bpf_refcount_acquire_impl.md)
    - [bpf_skb_ct_alloc](../kfuncs/bpf_skb_ct_alloc.md)
    - [bpf_skb_ct_lookup](../kfuncs/bpf_skb_ct_lookup.md)
    - [bpf_skb_get_fou_encap](../kfuncs/bpf_skb_get_fou_encap.md)
    - [bpf_skb_get_xfrm_info](../kfuncs/bpf_skb_get_xfrm_info.md)
    - [bpf_skb_set_fou_encap](../kfuncs/bpf_skb_set_fou_encap.md)
    - [bpf_skb_set_xfrm_info](../kfuncs/bpf_skb_set_xfrm_info.md)
    - [bpf_task_acquire](../kfuncs/bpf_task_acquire.md)
    - [bpf_task_from_pid](../kfuncs/bpf_task_from_pid.md)
    - [bpf_task_get_cgroup1](../kfuncs/bpf_task_get_cgroup1.md)
    - [bpf_task_release](../kfuncs/bpf_task_release.md)
    - [bpf_task_under_cgroup](../kfuncs/bpf_task_under_cgroup.md)
    - [bpf_throw](../kfuncs/bpf_throw.md)
    - [bpf_xdp_ct_alloc](../kfuncs/bpf_xdp_ct_alloc.md)
    - [bpf_xdp_ct_lookup](../kfuncs/bpf_xdp_ct_lookup.md)
    - [bpf_xdp_get_xfrm_state](../kfuncs/bpf_xdp_get_xfrm_state.md)
    - [bpf_xdp_xfrm_state_release](../kfuncs/bpf_xdp_xfrm_state_release.md)
    - [crash_kexec](../kfuncs/crash_kexec.md)
<!-- [/PROG_KFUNC_REF] -->
