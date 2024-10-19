---
title: "Libbpf eBPF macro '__long'"
description: "This page documents the '__long' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__long`

[:octicons-tag-24: v1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)

The `__long` macros is used to define unsigned long properties of BTF maps.

## Definition

`#!c #define __ulong(name, val) enum { ___bpf_concat(__unique_value, __COUNTER__) = val } name`

## Usage

This macro is used to encode unsigned long properties in BTF map definitions. BTF does not have a notion of literal values, we encode them as enum with value `X`, were `X` is the actual number we want to communicate.

The `__ulong` supports up to 64-bit values, unlike the `__uint` macro which only supports 32-bit values.

### Example

```c hl_lines="9 11"
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2024 Meta Platforms, Inc. and affiliates. */

struct {
	__uint(type, BPF_MAP_TYPE_ARENA);
	__uint(map_flags, BPF_F_MMAPABLE);
	__uint(max_entries, 10); /* number of pages */
#ifdef __TARGET_ARCH_arm64
	__ulong(map_extra, 0x1ull << 32); /* start of mmap() region */
#else
	__ulong(map_extra, 0x1ull << 44); /* start of mmap() region */
#endif
} arena SEC(".maps");
```

[Source](https://github.com/torvalds/linux/blob/3d5ad2d4eca337e80f38df77de89614aa5aaceb9/tools/testing/selftests/bpf/progs/arena_atomics.c)
