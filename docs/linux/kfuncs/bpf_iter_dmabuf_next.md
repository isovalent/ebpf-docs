---
title: "KFunc 'bpf_iter_dmabuf_next'"
description: "This page documents the 'bpf_iter_dmabuf_next' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_dmabuf_next`

<!-- [FEATURE_TAG](bpf_iter_dmabuf_next) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/6eab7ac7c5eea7628b92cd5f9427bbd963a954ec)
<!-- [/FEATURE_TAG] -->

Advance the iterator to the next DMA buffer.

## Definition

**Parameter**

`it`: The DMA buffer iterator to advance

**Returns**

The current DMA buffer, or `NULL` if at the end of the iterator.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct dma_buf *bpf_iter_dmabuf_next(struct bpf_iter_dmabuf *it)`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

See [`bpf_iter_dmabuf_new`](bpf_iter_dmabuf_new.md#usage) for usage details.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [`bpf_iter_dmabuf_new`](bpf_iter_dmabuf_new.md#example) for examples.


