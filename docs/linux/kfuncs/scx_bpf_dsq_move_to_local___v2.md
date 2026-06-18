---
title: "KFunc 'scx_bpf_dsq_move_to_local___v2'"
description: "This page documents the 'scx_bpf_dsq_move_to_local___v2' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_move_to_local___v2`

<!-- [FEATURE_TAG](scx_bpf_dsq_move_to_local___v2) -->
[:octicons-tag-24: 7.1](https://github.com/torvalds/linux/commit/860683763ebf4662cb72a312279334e02718308f)
<!-- [/FEATURE_TAG] -->

This function moves a task from a DSQ to the current CPU's local DSQ.

## Definition

Move a task from the non-local DSQ identified by `dsq_id` to the current CPU's local DSQ for execution. Can only be called from [`sched_ext_ops.dispatch`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dispatch).

This function flushes the in-flight dispatches from [`scx_bpf_dsq_insert`](scx_bpf_dsq_insert.md) before trying to move from the specified DSQ. It may also grab <nospell>rq</nospell> locks and thus can't be called under any BPF locks.

**Parameters**

`dsq_id`: DSQ to move task from
`enq_flags`: [`SCX_ENQ_*`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#enum-scx_enq_flags)

**Returns**

Returns `true` if a task has been moved, `false` if there isn't any task to move.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool scx_bpf_dsq_move_to_local___v2(u64 dsq_id, u64 enq_flags)`
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

