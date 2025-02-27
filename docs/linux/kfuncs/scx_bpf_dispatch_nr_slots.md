---
title: "KFunc 'scx_bpf_dispatch_nr_slots'"
description: "This page documents the 'scx_bpf_dispatch_nr_slots' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dispatch_nr_slots`

<!-- [FEATURE_TAG](scx_bpf_dispatch_nr_slots) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

This function returns the number of remaining dispatch slots.

## Definition

Can only be called from [`sched_ext_ops.dispatch`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dispatch).

**Returns**

The number of remaining dispatch slots

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 scx_bpf_dispatch_nr_slots()`
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

