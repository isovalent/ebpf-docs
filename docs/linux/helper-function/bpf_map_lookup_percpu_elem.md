---
title: "Helper Function 'bpf_map_lookup_percpu_elem'"
description: "This page documents the 'bpf_map_lookup_percpu_elem' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_map_lookup_percpu_elem`

<!-- [FEATURE_TAG](bpf_map_lookup_percpu_elem) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/07343110b293456d30393e89b86c4dee1ac051c8)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Perform a lookup in _percpu map_ for an entry associated to _key_ on _cpu_.

### Returns

Map value associated to _key_ on _cpu_, or **NULL** if no entry was found or _cpu_ is invalid.

`#!c static void *(* const bpf_map_lookup_percpu_elem)(void *map, const void *key, __u32 cpu) = (void *) 195;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage
The `map` argument must be a pointer to a map definition, `key` must be a pointer to the key you wish to lookup and `cpu` should be an integer starting from 0.

The return value will be a pointer to the map value in the hash slot for a specific CPU or NULL if no entry was found or cpu is invalid.

### Program types

This helper call can be used in the following program types:

<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
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

<!-- [HELPER_FUNC_MAP_REF] -->
 * [`BPF_MAP_TYPE_LRU_PERCPU_HASH`](../map-type/BPF_MAP_TYPE_LRU_PERCPU_HASH.md)
 * [`BPF_MAP_TYPE_PERCPU_ARRAY`](../map-type/BPF_MAP_TYPE_PERCPU_ARRAY.md)
 * [`BPF_MAP_TYPE_PERCPU_HASH`](../map-type/BPF_MAP_TYPE_PERCPU_HASH.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

```c
int key, *value, cpuid;
key=0;
cpuid=bpf_get_smp_processor_id();
value = bpf_map_lookup_percpu_elem(&percpu_map, &key, cpuid);
if (value)
	bpf_printk("Read value '%d' from the map on CPU '%d'\n", *value, cpuid);
else
	bpf_printk("Failed to read value from the map\n");
```
