---
title: "Helper Function 'bpf_get_func_arg_cnt'"
description: "This page documents the 'bpf_get_func_arg_cnt' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_func_arg_cnt`

<!-- [FEATURE_TAG](bpf_get_func_arg_cnt) -->
[:octicons-tag-24: v5.17](https://github.com/torvalds/linux/commit/f92c1e183604c20ce00eb889315fdaa8f2d9e509)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get number of registers of the traced function (for tracing programs) where function arguments are stored in these registers.



### Returns

The number of argument registers of the traced function.

`#!c static long (* const bpf_get_func_arg_cnt)(void *ctx) = (void *) 185;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
