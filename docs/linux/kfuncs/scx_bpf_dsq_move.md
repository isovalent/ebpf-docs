---
title: "KFunc 'scx_bpf_dsq_move'"
description: "This page documents the 'scx_bpf_dsq_move' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_move`

<!-- [FEATURE_TAG](scx_bpf_dsq_move) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/5cbb302880f50f3edf35f8c6a1d38b6948bf4d11)
<!-- [/FEATURE_TAG] -->

This function moves a task from DSQ iteration to a DSQ.

## Definition

Transfer `p` which is on the DSQ currently iterated by `it__iter` to the DSQ specified by `dsq_id`. All DSQs - local DSQs, global DSQ and user DSQs - can be the destination.

For the transfer to be successful, `p` must still be on the DSQ and have been queued before the DSQ iteration started. This function doesn't care whether `p` was obtained from the DSQ iteration. `p` just has to be on the DSQ and have been queued before the iteration started.

`p`'s slice is kept by default. Use [`scx_bpf_dsq_move_set_slice`](scx_bpf_dsq_move_set_slice.md) to update.

Can be called from [`sched_ext_ops.dispatch`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dispatch) or any BPF context which doesn't hold a <nospell>rq</nospell> lock (e.g. [BPF timers](../concepts/timers.md) or [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md) programs).

**Parameters**

`it__iter`: DSQ iterator in progress

`p:` task to transfer

`dsq_id`: DSQ to move `p` to

`enq_flags`: Bitfield of flags, see [`enum scx_enq_flags`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#enum-scx_enq_flags) for valid values.

**Return**

Returns `true` if `p` has been consumed, `false` if `p` had already been consumed or dequeued.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool scx_bpf_dsq_move(struct bpf_iter_scx_dsq *it__iter, struct task_struct *p, u64 dsq_id, u64 enq_flags)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

