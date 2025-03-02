---
title: "SCX eBPF macro '__contains'"
description: "This page documents the '__contains' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `__contains`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `__contains` macro is used during the definition of graph data structures such as [linked lists](https://github.com/torvalds/linux/commit/f0c5941ff5b255413d31425bb327c2aec3625673) or [red-black trees](https://github.com/torvalds/linux/commit/9c395c1b99bd23f74bc628fa000480c49593d17f) to inform the verifier of the types that will make up the graph.

## Definition

```c
#define __contains(name, node) __attribute__((btf_decl_tag("contains:" #name ":" #node)))
```

## Usage

This macro is used when declaring the root/head of a graph. These graphs, in eBPF must always be made up of of nodes that embed known types that actually implement the graph. Followed by the data the user wishes to store in the graph node.

For example:

```c
struct my_linked_list_node {
    struct bpf_list_node list_node;
    __u64 data;
}
```

or 

```c
struct my_rb_tree_node {
    struct bpf_rb_node rb_node;
    __u64 data;
}
```

Each known type has a corresponding root / head type which is what would be part of a map value or global variable. It is on this root type that the `__contains` macro is used.

The `name` parameter is the name of the node type. The `node` parameter of the macro informs the kernel of the field name of the embedded known type.

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

[private](private.md)(CGV_TREE) struct bpf_rb_root cgv_tree __contains(cgv_node, rb_node);
```
