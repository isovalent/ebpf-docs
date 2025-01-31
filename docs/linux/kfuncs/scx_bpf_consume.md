---
title: "KFunc 'scx_bpf_consume'"
description: "This page documents the 'scx_bpf_consume' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_consume`

<!-- [FEATURE_TAG](scx_bpf_consume) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

This function was renamed to [`scx_bpf_dsq_move_to_local`](scx_bpf_dsq_move_to_local.md) in [:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/5209c03c8ed215357a4827496a71fd32167d83ef). But will be aliased until v6.15.

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool scx_bpf_consume(u64 dsq_id)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

