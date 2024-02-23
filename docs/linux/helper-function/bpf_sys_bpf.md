---
title: "Helper Function 'bpf_sys_bpf' - eBPF Docs"
description: "This page documents the 'bpf_sys_bpf' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sys_bpf`

<!-- [FEATURE_TAG](bpf_sys_bpf) -->
[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/79a7f8bdb159d9914b58740f3d31d602a6e4aca8)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Execute bpf syscall with given arguments.

### Returns

A syscall result.

`#!c static long (*bpf_sys_bpf)(__u32 cmd, void *attr, __u32 attr_size) = (void *) 166;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
