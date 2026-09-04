---
title: "KFunc 'bpf_arena_reserve_pages'"
description: "This page documents the 'bpf_arena_reserve_pages' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_arena_reserve_pages`

<!-- [FEATURE_TAG](bpf_arena_reserve_pages) -->
[:octicons-tag-24: v6.17](https://github.com/torvalds/linux/commit/8fc3d2d8b5016adf63a3a6d21c189677fa653a4a)
<!-- [/FEATURE_TAG] -->

Reserve pages of memory for a arena.

## Definition

`p__map`: Pointer to the `BPF_MAP_TYPE_ARENA` map.

`ptr__ign`: Address of the start of the page(s) to be reserved, must be a page aligned address.

`page_cnt`: Number of pages to reserve.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_arena_reserve_pages(void *p__map, __arena void *ptr__ign, u32 page_cnt)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

Reserves a region of the mapping to prevent it from being mapped. This prevents the range from
being populated using [`bpf_arena_alloc_pages`](bpf_arena_alloc_pages.md). These regions serve as guards against out-of-bounds accesses and are useful for debugging arena-related code.

Reserved pages can be unreserved using [`bpf_arena_free_pages`](bpf_arena_free_pages.md). They can also be allocated from userspace through minor faults. It is up to the user to prevent erroneous frees and/or use the `BPF_F_SEGV_ON_FAULT` flag to catch stray userspace accesses.

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

