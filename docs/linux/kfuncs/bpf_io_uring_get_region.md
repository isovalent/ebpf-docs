---
title: "KFunc 'bpf_io_uring_get_region'"
description: "This page documents the 'bpf_io_uring_get_region' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_io_uring_get_region`

<!-- [FEATURE_TAG](bpf_io_uring_get_region) -->
[:octicons-tag-24: 7.1](https://github.com/torvalds/linux/commit/890819248a8616558fe12e6c06c918ee1c3a2bc6)
<!-- [/FEATURE_TAG] -->

Get a pointer to a memory region (kernel-userspace shared chunk of memory).

## Definition

**Parameters**

`loop_ctx`: a pointer to the `struct iou_ctx` to act on.
`region_id`: the ID of the region, which can be one of:
   * `IOU_REGION_SQ` returns the submission queue.
   * `IOU_REGION_CQ` stores the Completion Queue, Submission Queue/Completion Queue headers and the `sqarray`. In other words, it gives same memory that would normally be mmap'ed with `IORING_FEAT_SINGLE_MMAP` enabled `IORING_OFF_SQ_RING`.
   * `IOU_REGION_MEM` represents the memory / parameter region. It can be used to store request indirect parameters and for kernel - user communication.
`rdwr_buf_size`: the size as an argument, which should be a load time constant.

**Returns** 

A pointer to the specified region, where `io_uring` regions are kernel-userspace shared chunks of memory.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c __u8 *bpf_io_uring_get_region(struct io_ring_ctx *ctx, __u32 region_id, const size_t rdwr_buf_size)`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
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

