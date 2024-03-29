---
title: "Helper Function 'bpf_for_each_map_elem'"
description: "This page documents the 'bpf_for_each_map_elem' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_for_each_map_elem`

<!-- [FEATURE_TAG](bpf_for_each_map_elem) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/69c087ba6225b574afb6e505b72cb75242a3d844)
<!-- [/FEATURE_TAG] -->

For each element in `map`, call `callback_fn` function with `map`, `callback_ctx` and other map-specific parameters.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


`#!c static long (* const bpf_for_each_map_elem)(void *map, void *callback_fn, void *callback_ctx, __u64 flags) = (void *) 164;`

## Usage

The `callback_fn` should be a static function with the following signature:
`#!c long (*callback_fn)(struct bpf_map *map, const void *key, void *value, void *ctx);`


`callback_ctx` should be a pointer to a variable on the stack, its type can be determined by the caller. The same context is shared between all calls and can so be used to get information back from the callback to the main program.

The `flags` is used to control certain aspects of the helper. Currently, the `flags` must be 0.

For per_cpu maps, the map_value is the value on the cpu where the
bpf_prog is running.

If `callback_fn` return 0, the helper will continue to the next
element. If return value is 1, the helper will skip the rest of
elements and return. Other return values are not used now.

**Returns**
The number of traversed map elements for success, `-EINVAL` for
invalid `flags`.

### Program types

This helper call can be used in the following program types:

<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_DEVICE](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
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

<!-- [HELPER_FUNC_MAP_REF] -->
 * [BPF_MAP_TYPE_ARRAY](../map-type/BPF_MAP_TYPE_ARRAY.md)
 * [BPF_MAP_TYPE_HASH](../map-type/BPF_MAP_TYPE_HASH.md)
 * [BPF_MAP_TYPE_LRU_HASH](../map-type/BPF_MAP_TYPE_LRU_HASH.md)
 * [BPF_MAP_TYPE_LRU_PERCPU_HASH](../map-type/BPF_MAP_TYPE_LRU_PERCPU_HASH.md)
 * [BPF_MAP_TYPE_PERCPU_ARRAY](../map-type/BPF_MAP_TYPE_PERCPU_ARRAY.md)
 * [BPF_MAP_TYPE_PERCPU_HASH](../map-type/BPF_MAP_TYPE_PERCPU_HASH.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
