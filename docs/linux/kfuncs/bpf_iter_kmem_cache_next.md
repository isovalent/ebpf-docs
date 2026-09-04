---
title: "KFunc 'bpf_iter_kmem_cache_next'"
description: "This page documents the 'bpf_iter_kmem_cache_next' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_kmem_cache_next`

<!-- [FEATURE_TAG](bpf_iter_kmem_cache_next) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/2e9a548009c2d804e55cdd5b0e9903756cf7d9b3)
<!-- [/FEATURE_TAG] -->

This function returns the current slab cache and advances the iterator to the next slab cache.

## Definition

**Parameters**

`it`: A pointer to a stack allocated `struct bpf_iter_kmem_cache` that is used to iterate over slab caches.

**Returns**

A pointer to the next slab cache, or `NULL` if there are no more slab caches to iterate over.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct kmem_cache *bpf_iter_kmem_cache_next(struct bpf_iter_kmem_cache *it)`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).

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

