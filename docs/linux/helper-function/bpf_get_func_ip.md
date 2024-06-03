---
title: "Helper Function 'bpf_get_func_ip'"
description: "This page documents the 'bpf_get_func_ip' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_func_ip`

<!-- [FEATURE_TAG](bpf_get_func_ip) -->
[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/9b99edcae5c80c8fb9f8e7149bae528c9e610a72)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get address of the traced function (for tracing and kprobe programs).

When called for kprobe program attached as uprobe it returns probe address for both entry and return uprobe.



### Returns

Address of the traced function for kprobe. 0 for kprobes placed within the function (not at the entry). Address of the probe for uprobe and return uprobe.

`#!c static __u64 (* const bpf_get_func_ip)(void *ctx) = (void *) 173;`
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
