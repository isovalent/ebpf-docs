---
title: "KFunc 'scx_bpf_dsq_insert'"
description: "This page documents the 'scx_bpf_dsq_insert' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_insert`

<!-- [FEATURE_TAG](scx_bpf_dsq_insert) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)
<!-- [/FEATURE_TAG] -->

This function inserts a task into the First In First Out(FIFO) queue of a Dispatch Queue(DSQ)

## Definition

Insert `p` into the FIFO queue of the DSQ identified by `dsq_id`. It is safe to call this function spuriously. Can be called from [`sched_ext_ops.enqueue`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#enqueue), [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu), and [`sched_ext_ops.dispatch`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dispatch).

When called from [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu) or [`sched_ext_ops.enqueue`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#enqueue), it's for direct dispatch and `p` must match the task being enqueued.

When called from [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu), `enq_flags` and `dsp_id` are stored, and `p` will be directly inserted into the corresponding dispatch queue after [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu) returns. If `p` is inserted into `SCX_DSQ_LOCAL`, it will be inserted into the local DSQ of the CPU returned by [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu). `enq_flags` are OR'd with the enqueue flags on the enqueue path before the task is inserted.

When called from [`sched_ext_ops.dispatch`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dispatch), there are no restrictions on `p` or `dsq_id` and this function can be called up to [`sched_ext_ops.dispatch_max_batch`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dispatch_max_batch) times to insert multiple tasks. [`scx_bpf_dispatch_nr_slots`](scx_bpf_dispatch_nr_slots.md) returns the number of the remaining slots. [`scx_bpf_consume`](scx_bpf_consume.md) flushes the batch and resets the counter.

This function doesn't have any locking restrictions and may be called under BPF locks (in the future when BPF introduces more flexible locking).

`p` is allowed to run for `slice`. The scheduling path is triggered on slice exhaustion. If zero, the current residual slice is maintained. If [`SCX_SLICE_INF`](https://elixir.bootlin.com/linux/v6.13.4/source/include/linux/sched/ext.h#L21), `p` never expires and the BPF scheduler must kick the CPU with [`scx_bpf_kick_cpu`](scx_bpf_kick_cpu.md) to trigger scheduling.

**Parameters**

`p`: task_struct to insert

`dsq_id`: DSQ to insert into

`slice`: duration `p` can run for in nanoseconds, 0 to keep the current value

`enq_flags`: Bitfield of flags, see [`enum scx_enq_flags`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#enum-scx_enq_flags) for valid values.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dsq_insert(struct task_struct *p, u64 dsq_id, u64 slice, u64 enq_flags)`
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

