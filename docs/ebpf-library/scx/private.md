---
title: "SCX eBPF macro 'private'"
description: "This page documents the 'private' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `private`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `private` macro is used to mark global variables such that they will not be used from userspace.

## Definition

```c
#define private(name) [SEC](../libbpf/ebpf/SEC.md)(".data." #name) [__hidden](../libbpf/ebpf/__hidden.md) __attribute__((aligned(8)))
```

## Usage

This macro places a global variable in a section named `.data.<name>` and marks it as hidden. When all global variables in the same section are hidden, loaders will avoid mmap'ing the section into userspace and will make the section as `static` when loading. This allows the verifier to make stronger assumptions about the section and the variables in it.

### Example

```c hl_lines="18"
/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (c) 2022 Tejun Heo <tj@kernel.org> */

struct bpf_rb_root {
	__u64 __opaque[2];
} __attribute__((aligned(8)));

struct bpf_rb_node {
	__u64 __opaque[4];
} __attribute__((aligned(8)));

struct cgv_node {
	struct bpf_rb_node  rb_node;
	__u64               cvtime;
	__u64               cgid;
};

private.md(CGV_TREE) struct bpf_rb_root cgv_tree [__contains](__contains.md)(cgv_node, rb_node);
```
