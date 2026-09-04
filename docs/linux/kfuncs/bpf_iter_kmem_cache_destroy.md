---
title: "KFunc 'bpf_iter_kmem_cache_destroy'"
description: "This page documents the 'bpf_iter_kmem_cache_destroy' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_kmem_cache_destroy`

<!-- [FEATURE_TAG](bpf_iter_kmem_cache_destroy) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/2e9a548009c2d804e55cdd5b0e9903756cf7d9b3)
<!-- [/FEATURE_TAG] -->

This function destroys the iterator for slab caches.

## Definition

**Parameters**

`it`: A pointer to a stack allocated `struct bpf_iter_kmem_cache`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_iter_kmem_cache_destroy(struct bpf_iter_kmem_cache *it)`

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

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

