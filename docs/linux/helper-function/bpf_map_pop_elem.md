---
title: "Helper Function 'bpf_map_pop_elem'"
description: "This page documents the 'bpf_map_pop_elem' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_map_pop_elem`

<!-- [FEATURE_TAG](bpf_map_pop_elem) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/f1a2e44a3aeccb3ff18d3ccc0b0203e70b95bd92)
<!-- [/FEATURE_TAG] -->


## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Pop an element from _map_.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_map_pop_elem)(void *map, void *value) = (void *) 88;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LIRC_MODE2`](../program-type/BPF_PROG_TYPE_LIRC_MODE2.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [`BPF_MAP_TYPE_QUEUE`](../map-type/BPF_MAP_TYPE_QUEUE.md)
 * [`BPF_MAP_TYPE_STACK`](../map-type/BPF_MAP_TYPE_STACK.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

```c
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>

#define TC_ACT_OK 0

struct {
   __uint(type, BPF_MAP_TYPE_QUEUE);
   __uint(max_entries, 8);
   __uint(value_size, sizeof(__u32));
} ingress SEC(".maps");

SEC("tc")
int tc_ingress(struct __sk_buff *skb)
{
    void* data_end = (void *)(long)skb->data_end;
    void* data = (void *)(long)skb->data;
    u32 value;
    int err;
    
    struct ethhdr* eth = (struct ethhdr *)data;
    if ((void *)(eth + 1) > data_end)
        return TC_ACT_OK;
    
    struct iphdr* iph = (struct iphdr *)(eth + 1);
    if ((void *)(iph + 1) > data_end)
   	    return TC_ACT_OK;  // Check IP header

    err = bpf_map_push_elem(&ingress, &iph->saddr, 0);
    bpf_printk("Pushed something to queue");
    if (err)
        return TC_ACT_OK;
    
    err = bpf_map_peek_elem(&ingress, &value); 
    bpf_printk("Peeked at something in queue");
    if (err)
        return TC_ACT_OK;

    if(value == iph->saddr)
    {
        err = bpf_map_pop_elem(&ingress, &value);
        bpf_printk("Popped something from queue");
   	    if (err)
       	    return TC_ACT_OK;
   }
   return TC_ACT_OK;
}
char _license[] SEC("license") = "GPL";
```
