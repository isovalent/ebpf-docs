---
title: "KFunc 'scx_bpf_dispatch_cancel'"
description: "This page documents the 'scx_bpf_dispatch_cancel' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dispatch_cancel`

<!-- [FEATURE_TAG](scx_bpf_dispatch_cancel) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

This function cancels the latest dispatch.

## Definition

Cancel the latest dispatch. Can be called multiple times to cancel further dispatches. Can only be called from `ops.dispatch()`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dispatch_cancel()`
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

