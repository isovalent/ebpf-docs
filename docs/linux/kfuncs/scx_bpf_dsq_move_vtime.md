---
title: "KFunc 'scx_bpf_dsq_move_vtime'"
description: "This page documents the 'scx_bpf_dsq_move_vtime' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_move_vtime`

<!-- [FEATURE_TAG](scx_bpf_dsq_move_vtime) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/5cbb302880f50f3edf35f8c6a1d38b6948bf4d11)
<!-- [/FEATURE_TAG] -->

This function moves a task from DSQ iteration to a <nospell>PRIQ</nospell> DSQ.

## Definition

Transfer `p` which is on the DSQ currently iterated by `it__iter` to the priority queue of the DSQ specified by `dsq_id`. The destination must be a user DSQ as only user DSQs support priority queue.

`p`'s slice and vtime are kept by default. Use [`scx_bpf_dsq_move_set_slice`](scx_bpf_dsq_move_set_slice.md) and [`scx_bpf_dsq_move_set_vtime`](scx_bpf_dsq_move_set_vtime.md) to update.

All other aspects are identical to [`scx_bpf_dsq_move`](scx_bpf_dsq_move.md). See [`scx_bpf_dsq_insert_vtime`](scx_bpf_dsq_insert_vtime.md) for more information on `vtime`.

**Parameters**

`it__iter`: DSQ iterator in progress

`p:` task to transfer

`dsq_id`: DSQ to move `p` to

`enq_flags`: `SCX_ENQ_*`


**Flags**

`SCX_ENQ_WAKEUP`: Task just became runnable

`SCX_ENQ_HEAD`: Place at front of queue (tail if not specified)

`SCX_ENQ_CPU_SELECTED`: `->select_task_rq()` was called

`SCX_ENQ_PREEMPT`: Set the following to trigger preemption when calling `scx_bpf_dsq_insert` with a local dsq as the target. The slice of the current task is cleared to zero and the CPU is kicked into the scheduling path. Implies `SCX_ENQ_HEAD`.

`SCX_ENQ_REENQ`: The task being enqueued was previously enqueued on the current CPU's `SCX_DSQ_LOCAL`, but was removed from it in a call to the [`scx_bpf_reenqueue_local`](scx_bpf_reenqueue_local.md) kfunc. If [`scx_bpf_reenqueue_local`](scx_bpf_reenqueue_local.md) was invoked in a `->cpu_release()` callback, and the task is again dispatched back to `SCX_LOCAL_DSQ` by this current `->enqueue()`, the task will not be scheduled on the CPU until at least the next invocation of the `->cpu_acquire()` callback.

`SCX_ENQ_LAST`: The task being enqueued is the only task available for the cpu. By default, ext core keeps executing such tasks but when `SCX_OPS_ENQ_LAST` is specified, they are `ops.enqueue()`'d with the `SCX_ENQ_LAST` flag set. The BPF scheduler is responsible for triggering a follow-up scheduling event. Otherwise, Execution may stall.

**Return**

Returns `true` if `p` has been consumed, `false` if `p` had already been consumed or dequeued.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool scx_bpf_dsq_move_vtime(struct bpf_iter_scx_dsq *it__iter, struct task_struct *p, u64 dsq_id, u64 enq_flags)`
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

