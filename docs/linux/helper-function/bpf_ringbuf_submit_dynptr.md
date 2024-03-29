---
title: "Helper Function 'bpf_ringbuf_submit_dynptr'"
description: "This page documents the 'bpf_ringbuf_submit_dynptr' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_ringbuf_submit_dynptr`

<!-- [FEATURE_TAG](bpf_ringbuf_submit_dynptr) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/bc34dee65a65e9c920c420005b8a43f2a721a458)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Submit reserved ring buffer sample, pointed to by _data_, through the dynptr interface. This is a no-op if the dynptr is invalid/null.

For more information on _flags_, please see 'bpf_ringbuf_submit'.

### Returns

Nothing. Always succeeds.

`#!c static void (* const bpf_ringbuf_submit_dynptr)(struct bpf_dynptr *ptr, __u64 flags) = (void *) 199;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
