---
title: "Helper Function 'bpf_dynptr_read' - eBPF Docs"
description: "This page documents the 'bpf_dynptr_read' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_dynptr_read`

<!-- [FEATURE_TAG](bpf_dynptr_read) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/13bbbfbea7598ea9f8d9c3d73bf053bb57f9c4b2)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Read _len_ bytes from _src_ into _dst_, starting from _offset_ into _src_. _flags_ is currently unused.

### Returns

0 on success, -E2BIG if _offset_ + _len_ exceeds the length of _src_'s data, -EINVAL if _src_ is an invalid dynptr or if _flags_ is not 0.

`#!c static long (*bpf_dynptr_read)(void *dst, __u32 len, const struct bpf_dynptr *src, __u32 offset, __u64 flags) = (void *) 201;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
