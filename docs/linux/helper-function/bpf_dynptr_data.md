---
title: "Helper Function 'bpf_dynptr_data'"
description: "This page documents the 'bpf_dynptr_data' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_dynptr_data`

<!-- [FEATURE_TAG](bpf_dynptr_data) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/34d4ef5775f776ec4b0d53a02d588bf3195cada6)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get a pointer to the underlying dynptr data.

_len_ must be a statically known value. The returned data slice is invalidated whenever the dynptr is invalidated.

skb and xdp type dynptrs may not use bpf_dynptr_data. They should instead use bpf_dynptr_slice and bpf_dynptr_slice_rdwr.

### Returns

Pointer to the underlying dynptr data, NULL if the dynptr is read-only, if the dynptr is invalid, or if the offset and length is out of bounds.

`#!c static void *(* const bpf_dynptr_data)(const struct bpf_dynptr *ptr, __u32 offset, __u32 len) = (void *) 203;`
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
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
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

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
