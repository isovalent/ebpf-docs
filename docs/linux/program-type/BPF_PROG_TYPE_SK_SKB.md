---
title: "Program Type 'BPF_PROG_TYPE_SK_SKB'"
description: "This page documents the 'BPF_PROG_TYPE_SK_SKB' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_SK_SKB`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_SK_SKB) -->
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/b005fd189cec9407b700599e1e80e0552446ee79)
<!-- [/FEATURE_TAG] -->

Socket SKB programs are called on L4 data streams to parse L7 messages and/or to determine if the L4/L7 messages should be allowed, blocked or redirected.

## Usage

Socket SKB programs are attached to [`BPF_MAP_TYPE_SOCKMAP`](../map-type/BPF_MAP_TYPE_SOCKMAP.md) or [`BPF_MAP_TYPE_SOCKHASH`](../map-type/BPF_MAP_TYPE_SOCKHASH.md) maps and will be invoked when messages get received on the sockets which are part of the map the program is attached to. The exact purpose of the program differs depending on its attach type.

### As `BPF_SK_SKB_STREAM_PARSER` program

When this attach type is used the program acts as a [stream parser](https://www.kernel.org/doc/Documentation/networking/strparser.txt). The idea behind a stream parser is to parse message based application layer protocols (OSI Layer 7) which are implemented on top of data streams such as TCP.

The job of the program is to parse the L7 data/packet and to tell the kernel how long the L7 message is. This will allow the kernel to combine multiple data stream packets and return complete L7 messages for every [`recv`](https://man7.org/linux/man-pages/man2/recv.2.html) instead of returning the TCP messages which might only contain part of the L7 message.

The return value is interpreted as follows:

* `>0` - indicates length of successfully parsed message
* `0` - indicates more data must be received to parse the message
* `-ESTRPIPE` - current message should not be processed by the kernel, return control of the socket to userspace which can proceed to read the messages itself
* `other < 0` - Error in parsing, give control back to userspace assuming that synchronization is lost and the stream is unrecoverable (application expected to close TCP socket)

!!! note
    Before [:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/ef5659280eb13e8ac31c296f58cfdfa1684ac06b) it was required to have a stream parser attached to a `BPF_MAP_TYPE_SOCKMAP` if you wanted to use the stream verdict as well. On newer versions this is no longer required.

    On the older kernels, a no-op program can be used to just return the length of the current skb to retain default behavior and pass verdict per TCP packet.

    ```c
    SEC("sk_skb/stream_parser")
    int noop_parser(struct __sk_buff *skb)
    {
        return skb->len;
    }
    ```

### As `BPF_SK_SKB_STREAM_VERDICT` program

When this attach type is used the program acts as a filter, comparable to [TC](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) or [XDP](../program-type/BPF_PROG_TYPE_XDP.md) programs. The program gets called for every message indicated by the parser (or TCP packet if no parser is specified) and returns a verdict.

The return value is interpreted as follows:

* `SK_PASS` - The message may pass to the socket or it has been redirected with a helper.
* `SK_DROP` - The message should be dropped.

Unlike TC or XDP programs, there is no special redirect return code, helpers such as [`bpf_sk_redirect_map`](../helper-function/bpf_sk_redirect_map.md) will return `SK_PASS` on success.

### As `BPF_SK_SKB_VERDICT` program
<!-- [FEATURE_TAG](BPF_SK_SKB_VERDICT) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/a7ba4558e69a3c2ae4ca521f015832ef44799538)
<!-- [/FEATURE_TAG] -->

The non-stream verdict attach type is a replacement for the `BPF_SK_SKB_STREAM_VERDICT` attach type. The program type has the same job and uses the same return values. The difference is that this the stream verdict variant only supports TCP data streams while `BPF_SK_SKB_VERDICT` also supports UDP.

## Context

Socket SKB programs are called by the kernel with a [`__sk_buff`](../program-context/__sk_buff.md) context.

This program type isn't allowed to read from and write to all fields of the context since doing so might break assumptions in the kernel or because data isn't available at the point where the program is hooked into the kernel.

<!-- Information based on `sk_skb_is_valid_access` and `bpf_skb_is_valid_access` functions in the kernel sources -->

??? abstract "Context fields"
    | Field                                                                | Read             | Write            |
    | -------------------------------------------------------------------- | ---------------- | ---------------- |
    | [`len`](../program-context/__sk_buff.md#len)                         | :material-check: | :material-close: |
    | [`pkt_type`](../program-context/__sk_buff.md#pkt_type)               | :material-check: | :material-close: |
    | [`mark`](../program-context/__sk_buff.md#mark)                       | :material-close: | :material-close: |
    | [`queue_mapping`](../program-context/__sk_buff.md#queue_mapping)     | :material-check: | :material-close: |
    | [`protocol`](../program-context/__sk_buff.md#protocol)               | :material-check: | :material-close: |
    | [`vlan_present`](../program-context/__sk_buff.md#vlan_present)       | :material-check: | :material-close: |
    | [`vlan_tci`](../program-context/__sk_buff.md#vlan_tci)               | :material-check: | :material-close: |
    | [`vlan_proto`](../program-context/__sk_buff.md#vlan_proto)           | :material-check: | :material-close: |
    | [`priority`](../program-context/__sk_buff.md#priority)               | :material-check: | :material-check: |
    | [`ingress_ifindex`](../program-context/__sk_buff.md#ingress_ifindex) | :material-check: | :material-close: |
    | [`ifindex`](../program-context/__sk_buff.md#ifindex)                 | :material-check: | :material-close: |
    | [`tc_index`](../program-context/__sk_buff.md#tc_index)               | :material-check: | :material-check: |
    | [`cb`](../program-context/__sk_buff.md#cb)                           | :material-check: | :material-close: |
    | [`hash`](../program-context/__sk_buff.md#hash)                       | :material-check: | :material-close: |
    | [`tc_classid`](../program-context/__sk_buff.md#tc_classid)           | :material-close: | :material-close: |
    | [`data`](../program-context/__sk_buff.md#data)                       | :material-check: | :material-close: |
    | [`data_end`](../program-context/__sk_buff.md#data_end)               | :material-check: | :material-close: |
    | [`napi_id`](../program-context/__sk_buff.md#napi_id)                 | :material-check: | :material-close: |
    | [`family`](../program-context/__sk_buff.md#family)                   | :material-check: | :material-close: |
    | [`remote_ip4`](../program-context/__sk_buff.md#remote_ip4)           | :material-check: | :material-close: |
    | [`local_ip4`](../program-context/__sk_buff.md#local_ip4)             | :material-check: | :material-close: |
    | [`remote_ip4`](../program-context/__sk_buff.md#remote_ip4)           | :material-check: | :material-close: |
    | [`remote_ip6`](../program-context/__sk_buff.md#remote_ip6)           | :material-check: | :material-close: |
    | [`local_ip6`](../program-context/__sk_buff.md#local_ip6)             | :material-check: | :material-close: |
    | [`remote_port`](../program-context/__sk_buff.md#remote_port)         | :material-check: | :material-close: |
    | [`local_port`](../program-context/__sk_buff.md#local_port)           | :material-check: | :material-close: |
    | [`data_meta`](../program-context/__sk_buff.md#data_meta)             | :material-close: | :material-close: |
    | [`flow_keys`](../program-context/__sk_buff.md#flow_keys)             | :material-close: | :material-close: |
    | [`tstamp`](../program-context/__sk_buff.md#tstamp)                   | :material-check: | :material-close: |
    | [`wire_len`](../program-context/__sk_buff.md#wire_len)               | :material-close: | :material-close: |
    | [`tstamp`](../program-context/__sk_buff.md#tstamp)                   | :material-close: | :material-close: |
    | [`gso_segs`](../program-context/__sk_buff.md#gso_segs)               | :material-check: | :material-close: |
    | [`sk`](../program-context/__sk_buff.md#sk)                           | :material-check: | :material-close: |
    | [`gso_size`](../program-context/__sk_buff.md#gso_size)               | :material-check: | :material-close: |
    | [`tstamp_type`](../program-context/__sk_buff.md#tstamp_type)         | :material-close: | :material-close: |
    | [`hwtstamp`](../program-context/__sk_buff.md#hwtstamp)               | :material-close: | :material-close: |

## Attachment

Socket SKB programs are attached to [`BPF_MAP_TYPE_SOCKMAP`](../map-type/BPF_MAP_TYPE_SOCKMAP.md) or [`BPF_MAP_TYPE_SOCKHASH`](../map-type/BPF_MAP_TYPE_SOCKHASH.md) using the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall (`bpf_prog_attach` libbpf function).

The programs should be loaded with the same expected attach type as used during the attaching.

!!! note
    Before `BPF_SK_SKB_STREAM_VERDICT` and `BPF_SK_SKB_VERDICT` are mutually exclusive per map, only one or the other program type can be used.

## Example

Example BPF program:

```c
// Copyright Red Hat
SEC("sk_skb/stream_verdict")
int bpf_prog_verdict(struct __sk_buff *skb)
{
        __u32 lport = skb->local_port;
        __u32 idx = 0;

        if (lport == 10000)
                return bpf_sk_redirect_map(skb, &sock_map_rx, idx, 0);

        return SK_PASS;
}
```

Example userspace loader code:

```c
// Copyright Red Hat
int create_sample_sockmap(int sock, int parse_prog_fd, int verdict_prog_fd)
{
        int index = 0;
        int map, err;

        map = bpf_map_create(BPF_MAP_TYPE_SOCKMAP, NULL, sizeof(int), sizeof(int), 1, NULL);
        if (map < 0) {
                fprintf(stderr, "Failed to create sockmap: %s\n", strerror(errno));
                return -1;
        }

        err = bpf_prog_attach(parse_prog_fd, map, BPF_SK_SKB_STREAM_PARSER, 0);
        if (err){
                fprintf(stderr, "Failed to attach_parser_prog_to_map: %s\n", strerror(errno));
                goto out;
        }

        err = bpf_prog_attach(verdict_prog_fd, map, BPF_SK_SKB_STREAM_VERDICT, 0);
        if (err){
                fprintf(stderr, "Failed to attach_verdict_prog_to_map: %s\n", strerror(errno));
                goto out;
        }

        err = bpf_map_update_elem(map, &index, &sock, BPF_NOEXIST);
        if (err) {
                fprintf(stderr, "Failed to update sockmap: %s\n", strerror(errno));
                goto out;
        }

out:
        close(map);
        return err;
}
```

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_skb_store_bytes`](../helper-function/bpf_skb_store_bytes.md)
    * [`bpf_skb_load_bytes`](../helper-function/bpf_skb_load_bytes.md)
    * [`bpf_skb_pull_data`](../helper-function/bpf_skb_pull_data.md)
    * [`bpf_skb_change_tail`](../helper-function/bpf_skb_change_tail.md)
    * [`bpf_skb_change_head`](../helper-function/bpf_skb_change_head.md)
    * [`bpf_skb_adjust_room`](../helper-function/bpf_skb_adjust_room.md)
    * [`bpf_get_socket_cookie`](../helper-function/bpf_get_socket_cookie.md)
    * [`bpf_get_socket_uid`](../helper-function/bpf_get_socket_uid.md)
    * [`bpf_sk_redirect_map`](../helper-function/bpf_sk_redirect_map.md)
    * [`bpf_sk_redirect_hash`](../helper-function/bpf_sk_redirect_hash.md)
    * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
    * [`bpf_sk_lookup_tcp`](../helper-function/bpf_sk_lookup_tcp.md)
    * [`bpf_sk_lookup_udp`](../helper-function/bpf_sk_lookup_udp.md)
    * [`bpf_sk_release`](../helper-function/bpf_sk_release.md)
    * [`bpf_skc_lookup_tcp`](../helper-function/bpf_skc_lookup_tcp.md)
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
