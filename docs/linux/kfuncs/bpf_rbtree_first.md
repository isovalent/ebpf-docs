---
title: "KFunc 'bpf_rbtree_first'"
description: "This page documents the 'bpf_rbtree_first' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_rbtree_first`

<!-- [FEATURE_TAG](bpf_rbtree_first) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/bd1279ae8a691d7ec75852c6d0a22139afb034a4)
<!-- [/FEATURE_TAG] -->

Return the first (leftmost) node in input tree

## Definition

**Returns**

Pointer to the node, which is _not_ removed from the tree. If the tree contains no nodes, returns `NULL`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct bpf_rb_node *bpf_rbtree_first(struct bpf_rb_root *root)`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
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

