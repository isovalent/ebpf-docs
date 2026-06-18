---
title: "KFunc 'bpf_rbtree_add'"
description: "This page documents the 'bpf_rbtree_add' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_rbtree_add`

<!-- [FEATURE_TAG](bpf_rbtree_add) -->
[:octicons-tag-24: 7.1](https://github.com/torvalds/linux/commit/d457072576a6a60ba853b1d815f123da57b48021)
<!-- [/FEATURE_TAG] -->

Add `node` to red-black-tree with root `root` using comparator `less`

## Definition

**Parameters**

`root`: Root node of the red-back-tree.
`node`: Node to add to the tree.
`less`: Callback function used as comparator.

**Returns**

 * `0` if the node was successfully added
 * `-EINVAL` if the node wasn't added because it's already in a tree

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_rbtree_add(struct bpf_rb_root *root, struct bpf_rb_node *node, bool (less)(struct bpf_rb_node * , const struct bpf_rb_node * ))`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

