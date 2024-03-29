---
title: "Helper Function 'bpf_map_delete_elem'"
description: "This page documents the 'bpf_map_delete_elem' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_map_delete_elem`

<!-- [FEATURE_TAG](bpf_map_delete_elem) -->
[:octicons-tag-24: v3.19](https://github.com/torvalds/linux/commit/d0003ec01c667b731c139e23de3306a8b328ccf5)
<!-- [/FEATURE_TAG] -->

The delete map element helper call is used to delete values from [maps](../index.md#maps).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Delete entry with _key_ from _map_.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_map_delete_elem)(void *map, const void *key) = (void *) 3;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

The `map` argument must be a pointer to a map definition and `key` must be a pointer to the key you
wish to delete.

The return value will be `0` on success or a negative valued error number indicating a failure.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_DEVICE](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LIRC_MODE2](../program-type/BPF_PROG_TYPE_LIRC_MODE2.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [BPF_MAP_TYPE_ARRAY](../map-type/BPF_MAP_TYPE_ARRAY.md)
 * [BPF_MAP_TYPE_HASH](../map-type/BPF_MAP_TYPE_HASH.md)
 * [BPF_MAP_TYPE_LPM_TRIE](../map-type/BPF_MAP_TYPE_LPM_TRIE.md)
 * [BPF_MAP_TYPE_LRU_HASH](../map-type/BPF_MAP_TYPE_LRU_HASH.md)
 * [BPF_MAP_TYPE_LRU_PERCPU_HASH](../map-type/BPF_MAP_TYPE_LRU_PERCPU_HASH.md)
 * [BPF_MAP_TYPE_PERCPU_ARRAY](../map-type/BPF_MAP_TYPE_PERCPU_ARRAY.md)
 * [BPF_MAP_TYPE_PERCPU_HASH](../map-type/BPF_MAP_TYPE_PERCPU_HASH.md)
 * [BPF_MAP_TYPE_SOCKHASH](../map-type/BPF_MAP_TYPE_SOCKHASH.md)
 * [BPF_MAP_TYPE_SOCKMAP](../map-type/BPF_MAP_TYPE_SOCKMAP.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

```c
int key, result;
key = 1;
result = bpf_map_delete_element(&my_map, &key);
if (result == 0)
	bpf_printk("Element deleted from the map\n");
else
	bpf_printk("Failed to delete element from the map: %d\n",result);
```
