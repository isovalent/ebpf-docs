---
title: "KFunc '__scx_bpf_dsq_insert_vtime'"
description: "This page documents the '__scx_bpf_dsq_insert_vtime' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `__scx_bpf_dsq_insert_vtime`

<!-- [FEATURE_TAG](__scx_bpf_dsq_insert_vtime) -->
[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/c0d630ba347c7671210e1bab3c79defea19844e9)
<!-- [/FEATURE_TAG] -->

This function inserts a task into the `vtime` priority queue of a DSQ.

## Definition

Wrapper kfunc that takes arguments via struct to work around the 5 argument limit for BPF functions. BPF programs should use [`scx_bpf_dsq_insert_vtime`](../../ebpf-library/scx/scx_bpf_dsq_insert_vtime.md) which is provided as an inline wrapper in [`common.bpf.h`](../../ebpf-library/scx/index.md#commonbpfh).

Insert `p` into the vtime priority queue of the DSQ identified by `args->dsq_id`. Tasks queued into the priority queue are ordered by `args->vtime`. All other aspects are identical to [`scx_bpf_dsq_insert`](scx_bpf_dsq_insert.md).

`args->vtime` ordering is according to [`time_before64()`](https://elixir.bootlin.com/linux/v6.13/source/include/linux/jiffies.h#L212) which considers wrapping. A numerically larger vtime may indicate an earlier position in the ordering and vice-versa.

A DSQ can only be used as a FIFO or priority queue at any given time and this function must not be called on a DSQ which already has one or more FIFO tasks queued and vice-versa. Also, the built-in DSQs ([`SCX_DSQ_LOCAL`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_dsq_local) and [`SCX_DSQ_GLOBAL`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_dsq_global)) cannot be used as priority queues.

**Parameters**

`p`: task_struct to insert

`args`: struct containing the rest of the arguments

`args->dsq_id`: DSQ to insert into

`args->slice`: duration `p` can run for in nanoseconds, 0 to keep the current value

`args->vtime`: `p`'s ordering inside the vtime-sorted queue of the target DSQ

`args->enq_flags`: Bitfield of flags, see [`enum scx_enq_flags`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#enum-scx_enq_flags) for valid values.

**returns**

Returns `true` on successful insertion, `false` on failure. On the root scheduler, `false` return triggers scheduler abort and the caller doesn't need to check the return value.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool __scx_bpf_dsq_insert_vtime(struct task_struct *p, struct scx_bpf_dsq_insert_vtime_args *args)`
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

