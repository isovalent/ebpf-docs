---
title: "KFunc 'scx_bpf_dsq_insert_vtime'"
description: "This page documents the 'scx_bpf_dsq_insert_vtime' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_insert_vtime`

<!-- [FEATURE_TAG](scx_bpf_dsq_insert_vtime) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)
<!-- [/FEATURE_TAG] -->

This function inserts a task into the vtime priority queue of a DSQ.

## Definition

Insert `p` into the vtime priority queue of the DSQ identified by `dsq_id`. Tasks queued into the priority queue are ordered by `vtime`. All other aspects are identical to [`scx_bpf_dsq_insert`](scx_bpf_dsq_insert.md).

`vtime` ordering is according to [`time_before64()`](https://elixir.bootlin.com/linux/v6.13/source/include/linux/jiffies.h#L212) which considers wrapping. A numerically larger vtime may indicate an earlier position in the ordering and vice-versa.

A DSQ can only be used as a FIFO or priority queue at any given time and this function must not be called on a DSQ which already has one or more FIFO tasks queued and vice-versa. Also, the built-in DSQs (`SCX_DSQ_LOCAL` and `SCX_DSQ_GLOBAL`) cannot be used as priority queues.

**Parameters**

`p`: task_struct to insert

`dsq_id`: DSQ to insert into

`slice`: duration `p` can run for in nanoseconds, 0 to keep the current value

`vtime`: `p`'s ordering inside the vtime-sorted queue of the target DSQ

`enq_flags`: `SCX_ENQ_*`

**Flags**

`SCX_ENQ_WAKEUP`: Task just became runnable

`SCX_ENQ_HEAD`: Place at front of queue (tail if not specified)

`SCX_ENQ_CPU_SELECTED`: `->select_task_rq()` was called

`SCX_ENQ_PREEMPT`: Set the following to trigger preemption when calling `scx_bpf_dsq_insert` with a local dsq as the target. The slice of the current task is cleared to zero and the CPU is kicked into the scheduling path. Implies `SCX_ENQ_HEAD`.

`SCX_ENQ_REENQ`: The task being enqueued was previously enqueued on the current CPU's `SCX_DSQ_LOCAL`, but was removed from it in a call to the [`scx_bpf_reenqueue_local`](scx_bpf_reenqueue_local.md) kfunc. If [`scx_bpf_reenqueue_local`](scx_bpf_reenqueue_local.md) was invoked in a `->cpu_release()` callback, and the task is again dispatched back to `SCX_LOCAL_DSQ` by this current `->enqueue()`, the task will not be scheduled on the CPU until at least the next invocation of the `->cpu_acquire()` callback.

`SCX_ENQ_LAST`: The task being enqueued is the only task available for the cpu. By default, ext core keeps executing such tasks but when `SCX_OPS_ENQ_LAST` is specified, they are `ops.enqueue()`'d with the `SCX_ENQ_LAST` flag set. The BPF scheduler is responsible for triggering a follow-up scheduling event. Otherwise, Execution may stall.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dsq_insert_vtime(struct task_struct *p, u64 dsq_id, u64 slice, u64 vtime, u64 enq_flags)`
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

