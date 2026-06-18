---
title: "KFunc 'bpf_io_uring_submit_sqes'"
description: "This page documents the 'bpf_io_uring_submit_sqes' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_io_uring_submit_sqes`

<!-- [FEATURE_TAG](bpf_io_uring_submit_sqes) -->
[:octicons-tag-24: 7.1](https://github.com/torvalds/linux/commit/890819248a8616558fe12e6c06c918ee1c3a2bc6)
<!-- [/FEATURE_TAG] -->

Submit a number of `SQE`'s (Submission Queue Entry).

## Definition

**Parameters**

`loop_ctx`: pointer to the `struct iou_ctx` to act on.
`nr`: number of entries to submit.

**Returns** 

The number of entries submitted or a negative error code.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_io_uring_submit_sqes(struct io_ring_ctx *ctx, u32 nr)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
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

