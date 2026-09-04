---
title: "KFunc 'bpf_arena_alloc_pages'"
description: "This page documents the 'bpf_arena_alloc_pages' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_arena_alloc_pages`

<!-- [FEATURE_TAG](bpf_arena_alloc_pages) -->
[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/317460317a02a1af512697e6e964298dedd8a163)
<!-- [/FEATURE_TAG] -->

Allocate pages of memory for a arena.

## Definition

`p__map`: Pointer to the `BPF_MAP_TYPE_ARENA` map.

`addr__ign`: Address of the start of the page(s) to be allocated, must be a page aligned address.

`page_cnt`: Number of pages to allocate.

`node_id`: NUMA node to allocate memory from.

`flags`: Flags for future use, currently no valid flags exist.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c __arena void *bpf_arena_alloc_pages(void *p__map, __arena void *addr__ign, u32 page_cnt, int node_id, u64 flags)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
<!-- [/KFUNC_DEF] -->

## Usage

A BPF arena is a region of memory that can be shared between BPF programs and userspace programs. This allows for the creation of custom data structures that can be shared between BPF programs and userspace programs.

An arena is created as a map, upon its creation a maximum memory size is specified, but this memory isn't allocated at creation, rather, an arena allows on demand allocation of memory pages.

The kfunc is used to allocate these pages.

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

