---
title: "SCX eBPF macro 'bpf_rbtree_add'"
description: "This page documents the 'bpf_rbtree_add' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `bpf_rbtree_add`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `bpf_rbtree_add` macro wraps [`bpf_rbtree_add_impl`](../../linux/kfuncs/bpf_rbtree_add_impl.md) to provide a more ergonomic interface.

## Definition

```c
#define bpf_rbtree_add(head, node, less) [bpf_rbtree_add_impl](../../linux/kfuncs/bpf_rbtree_add_impl.md)(head, node, less, NULL, 0)
```

## Usage

The [`bpf_rbtree_add_impl`](../../linux/kfuncs/bpf_rbtree_add_impl.md) kfunc has a quirk where the forth argument is always `NULL`, this wrapper abstracts that quirk away.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
