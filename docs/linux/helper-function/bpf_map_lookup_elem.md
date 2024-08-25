---
title: "Helper Function 'bpf_map_lookup_elem'"
description: "This page documents the 'bpf_map_lookup_elem' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_map_lookup_elem`

<!-- [FEATURE_TAG](bpf_map_lookup_elem) -->
[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/0a542a86d73b1577e7d4f55fc95dcffd3fe62643)
<!-- [/FEATURE_TAG] -->

The lookup map element helper call is used to read values from [maps](../index.md#maps).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Perform a lookup in _map_ for an entry associated to _key_.

### Returns

Map value associated to _key_, or **NULL** if no entry was found.

`#!c static void *(* const bpf_map_lookup_elem)(void *map, const void *key) = (void *) 1;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

The `map` argument must be a pointer to a map definition and `key` must be a pointer to the key you
wish to lookup.

The return value will be a pointer to the map value or `NULL`. The value is a direct reference to the kernel memory where this map value is stored, not a copy. Therefor any modifications made to the value are automatically persisted without the need to call any additional helpers.

!!! warning
    modifying map values of non per-CPU maps is subject to race conditions, atomic instructions or spinlocks must be utilized to prevent race conditions if they are detrimental to your use case.
    <!-- TODO link to guide on memory access serialization -->


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
 * [`BPF_MAP_TYPE_ARRAY`](../map-type/BPF_MAP_TYPE_ARRAY.md)
 * [`BPF_MAP_TYPE_ARRAY_OF_MAPS`](../map-type/BPF_MAP_TYPE_ARRAY_OF_MAPS.md)
 * [`BPF_MAP_TYPE_HASH`](../map-type/BPF_MAP_TYPE_HASH.md)
 * [`BPF_MAP_TYPE_HASH_OF_MAPS`](../map-type/BPF_MAP_TYPE_HASH_OF_MAPS.md)
 * [`BPF_MAP_TYPE_LPM_TRIE`](../map-type/BPF_MAP_TYPE_LPM_TRIE.md)
 * [`BPF_MAP_TYPE_LRU_HASH`](../map-type/BPF_MAP_TYPE_LRU_HASH.md)
 * [`BPF_MAP_TYPE_LRU_PERCPU_HASH`](../map-type/BPF_MAP_TYPE_LRU_PERCPU_HASH.md)
 * [`BPF_MAP_TYPE_PERCPU_ARRAY`](../map-type/BPF_MAP_TYPE_PERCPU_ARRAY.md)
 * [`BPF_MAP_TYPE_PERCPU_HASH`](../map-type/BPF_MAP_TYPE_PERCPU_HASH.md)
 * [`BPF_MAP_TYPE_SOCKHASH`](../map-type/BPF_MAP_TYPE_SOCKHASH.md)
 * [`BPF_MAP_TYPE_SOCKMAP`](../map-type/BPF_MAP_TYPE_SOCKMAP.md)
 * [`BPF_MAP_TYPE_XSKMAP`](../map-type/BPF_MAP_TYPE_XSKMAP.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

```c
int key, value;
key = 1;
value = bpf_map_lookup_elem(&my_map, &key);
if (value)
	bpf_printk("Value read from the map: '%d'\n", value);
else
	bpf_printk("Failed to read value from the map\n");
```
