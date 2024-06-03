---
title: "Helper Function 'bpf_sys_close'"
description: "This page documents the 'bpf_sys_close' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_sys_close`

<!-- [FEATURE_TAG](bpf_sys_close) -->
[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/3abea089246f76c1517b054ddb5946f3f1dbd2c0)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Execute close syscall for given FD.

### Returns

A syscall result.

`#!c static long (* const bpf_sys_close)(__u32 fd) = (void *) 168;`
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
