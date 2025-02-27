---
title: "KFunc 'scx_bpf_dsq_insert_vtime'"
description: "This page documents the 'scx_bpf_dsq_insert_vtime' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_insert_vtime`

<!-- [FEATURE_TAG](scx_bpf_dsq_insert_vtime) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)
<!-- [/FEATURE_TAG] -->

This function inserts a task into the `vtime` priority queue of a DSQ.

## Definition

Insert `p` into the `vtime` priority queue of the DSQ identified by `dsq_id`. Tasks queued into the priority queue are ordered by `vtime`. All other aspects are identical to [`scx_bpf_dsq_insert`](scx_bpf_dsq_insert.md).

`vtime` ordering is according to [`time_before64()`](https://elixir.bootlin.com/linux/v6.13/source/include/linux/jiffies.h#L212) which considers wrapping. A numerically larger vtime may indicate an earlier position in the ordering and vice-versa.

A DSQ can only be used as a FIFO or priority queue at any given time and this function must not be called on a DSQ which already has one or more FIFO tasks queued and vice-versa. Also, the built-in DSQs ([`SCX_DSQ_LOCAL`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_dsq_local) and [`SCX_DSQ_GLOBAL`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_dsq_global)) cannot be used as priority queues.

**Parameters**

`p`: task_struct to insert

`dsq_id`: DSQ to insert into

`slice`: duration `p` can run for in nanoseconds, 0 to keep the current value

`vtime`: `p`'s ordering inside the vtime-sorted queue of the target DSQ

`enq_flags`: Bitfield of flags, see [`enum scx_enq_flags`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#enum-scx_enq_flags) for valid values.

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

