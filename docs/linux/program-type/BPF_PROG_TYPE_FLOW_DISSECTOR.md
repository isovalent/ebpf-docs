---
title: "Program Type 'BPF_PROG_TYPE_FLOW_DISSECTOR'"
description: "This page documents the 'BPF_PROG_TYPE_FLOW_DISSECTOR' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_FLOW_DISSECTOR`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_FLOW_DISSECTOR) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)
<!-- [/FEATURE_TAG] -->

Flow dissector is a program type that parses metadata out of the packets. 

## Usage

BPF flow dissectors can be attached per network namespace. These programs are given a packet, the program should fill out the rest of the `struct bpf_flow_keys` fields located at `__sk_buff->flow_keys`.

Various places in the networking subsystem use these flow keys to aggregate packets for the same "flow" a combination of these fields. By implementing this logic in BPF, it becomes possible to add flow parsing for new or custom protocols.

The return code of the BPF program is either `BPF_OK` to indicate successful dissection, or `BPF_DROP` to indicate parsing error.

## Context

BPF flow dissector programs operate on an [`__sk_buff`](../program-context/__sk_buff.md). 
However, only the limited set of fields is allowed: `data`, `data_end` and `flow_keys`. 

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

[`flow_keys`](../program-context/__sk_buff.md#flow_keys) is `struct bpf_flow_keys` and contains flow dissector input and output arguments. Input arguments `nhoff`/`thoff`/`n_proto` should be also adjusted accordingly.

??? abstract "c structure of `struct bpf_flow_keys`"
    ```c
    struct bpf_flow_keys {
        __u16	nhoff;
        __u16	thoff;
        __u16	addr_proto;			/* ETH_P_* of valid addrs */
        __u8	is_frag;
        __u8	is_first_frag;
        __u8	is_encap;
        __u8	ip_proto;
        __be16	n_proto;
        __be16	sport;
        __be16	dport;
        union {
            struct {
                __be32	ipv4_src;
                __be32	ipv4_dst;
            };
            struct {
                __u32	ipv6_src[4];	/* in6_addr; network order */
                __u32	ipv6_dst[4];	/* in6_addr; network order */
            };
        };
        __u32	flags;
        __be32	flow_label;
    };
    ```

The initial state of the input values can differ based on the type of packet handled and the state of the dissector. For example:

In the VLAN-less case, this is what the initial state of the BPF flow dissector looks like:
```
+------+------+------------+-----------+
| DMAC | SMAC | ETHER_TYPE | L3_HEADER |
+------+------+------------+-----------+
                            ^
                            |
                            +-- flow dissector starts here
```
```
skb->data + flow_keys->nhoff point to the first byte of L3_HEADER
flow_keys->thoff = nhoff
flow_keys->n_proto = ETHER_TYPE
```

In case of VLAN, flow dissector can be called with the two different states.

Pre-VLAN parsing:

```
+------+------+------+-----+-----------+-----------+
| DMAC | SMAC | TPID | TCI |ETHER_TYPE | L3_HEADER |
+------+------+------+-----+-----------+-----------+
                      ^
                      |
                      +-- flow dissector starts here
```
```
skb->data + flow_keys->nhoff point the to first byte of TCI
flow_keys->thoff = nhoff
flow_keys->n_proto = TPID
```

Please note that <nospell>TPID</nospell> can be 802.1AD and, hence, BPF program would have to parse VLAN information twice for double tagged packets.

Post-VLAN parsing:

```
+------+------+------+-----+-----------+-----------+
| DMAC | SMAC | TPID | TCI |ETHER_TYPE | L3_HEADER |
+------+------+------+-----+-----------+-----------+
                                        ^
                                        |
                                        +-- flow dissector starts here
```

```
skb->data + flow_keys->nhoff point the to first byte of L3_HEADER
flow_keys->thoff = nhoff
flow_keys->n_proto = ETHER_TYPE
```

In this case VLAN information has been processed before the flow dissector and BPF flow dissector is not required to handle it.

The takeaway here is as follows: BPF flow dissector program can be called with the optional VLAN header and should gracefully handle both cases: when single or double VLAN is present and when it is not present. The same program can be called for both cases and would have to be written carefully to handle both cases.

## Attachment

Flow dissector programs are attached to network namespaces via the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall or via [BPF link](../syscall/BPF_LINK_CREATE.md). 

This program type must always be loaded with the [`expected_attach_type`](../syscall/BPF_PROG_LOAD.md#expected_attach_type) of `BPF_FLOW_DISSECTOR`.

!!! warning
    `BPF_PROG_ATTACH` and links cannot be combined/used at the same time.

!!! note
    When a flow dissector is added to the root network namespace, it overwrites all other flow dissectors.

### `BPF_PROG_ATTACH` 

When attaching flow dissector programs via [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md), the program will be attached to the network namespace to which the current process is assigned. The specified target FD should be `0`.

### BPF link

To attach flow dissector programs to a network namespace using a link. You [creating the link](../syscall/BPF_LINK_CREATE.md) the `prog_fd` to the file descriptor of the program, `target_fd` should be set to the file descriptor of a network namespace, and the `attach_type` to `BPF_FLOW_DISSECTOR`.

## Example

The kernel maintains a reference example in-tree at [bpf_flow.c](https://elixir.bootlin.com/linux/v6.3/source/tools/testing/selftests/bpf/progs/bpf_flow.c).


## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for socket filter programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_skb_load_bytes`](../helper-function/bpf_skb_load_bytes.md)
    * [`bpf_skc_to_tcp6_sock`](../helper-function/bpf_skc_to_tcp6_sock.md)
    * [`bpf_skc_to_tcp_sock`](../helper-function/bpf_skc_to_tcp_sock.md)
    * [`bpf_skc_to_tcp_timewait_sock`](../helper-function/bpf_skc_to_tcp_timewait_sock.md)
    * [`bpf_skc_to_tcp_request_sock`](../helper-function/bpf_skc_to_tcp_request_sock.md)
    * [`bpf_skc_to_udp6_sock`](../helper-function/bpf_skc_to_udp6_sock.md)
    * [`bpf_skc_to_unix_sock`](../helper-function/bpf_skc_to_unix_sock.md)
    * [`bpf_ktime_get_coarse_ns`](../helper-function/bpf_ktime_get_coarse_ns.md)
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
There are currently no kfuncs supported for this program type
<!-- [/PROG_KFUNC_REF] -->
