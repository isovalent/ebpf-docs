---
title: "Helper Function 'bpf_probe_read_str'"
description: "This page documents the 'bpf_probe_read_str' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_probe_read_str`

<!-- [FEATURE_TAG](bpf_probe_read_str) -->
[:octicons-tag-24: v4.11](https://github.com/torvalds/linux/commit/a5e8c07059d0f0b31737408711d44794928ac218)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Copy a NUL terminated string from an unsafe kernel address _unsafe_ptr_ to _dst_. See **bpf_probe_read_kernel_str**() for more details.

Generally, use **bpf_probe_read_user_str**() or **bpf_probe_read_kernel_str**() instead.

### Returns

On success, the strictly positive length of the string, including the trailing NUL character. On error, a negative value.

`#!c static long (* const bpf_probe_read_str)(void *dst, __u32 size, const void *unsafe_ptr) = (void *) 45;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
