---
title: "KFunc 'scx_bpf_dsq_move_to_local'"
description: "This page documents the 'scx_bpf_dsq_move_to_local' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_move_to_local`

<!-- [FEATURE_TAG](scx_bpf_dsq_move_to_local) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/5209c03c8ed215357a4827496a71fd32167d83ef)
<!-- [/FEATURE_TAG] -->

This function moves a task from a DSQ to the current CPU's local DSQ.

## Definition

Move a task from the non-local DSQ identified by `dsq_id` to the current CPU's local DSQ for execution. Can only be called from `ops.dispatch()`.

This function flushes the in-flight dispatches from [`scx_bpf_dsq_insert`](scx_bpf_dsq_insert.md) before trying to move from the specified DSQ. It may also grab rq locks and thus can't be called under any BPF locks.

**Parameters**

`dsq_id`: DSQ to move task from

**Returns**

Returns `true` if a task has been moved, `false` if there isn't any task to move.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool scx_bpf_dsq_move_to_local(u64 dsq_id)`
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

