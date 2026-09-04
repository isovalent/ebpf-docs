---
title: "KFunc 'bpf_iter_dmabuf_destroy'"
description: "This page documents the 'bpf_iter_dmabuf_destroy' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_dmabuf_destroy`

<!-- [FEATURE_TAG](bpf_iter_dmabuf_destroy) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/6eab7ac7c5eea7628b92cd5f9427bbd963a954ec)
<!-- [/FEATURE_TAG] -->

This function destroys the iterator for DMA buffers.

## Definition

**Parameters**

`it`: A pointer to a stack allocated `struct bpf_iter_dmabuf`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_iter_dmabuf_destroy(struct bpf_iter_dmabuf *it)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

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


