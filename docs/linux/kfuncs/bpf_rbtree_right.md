---
title: "KFunc 'bpf_rbtree_right'"
description: "This page documents the 'bpf_rbtree_right' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_rbtree_right`

<!-- [FEATURE_TAG](bpf_rbtree_right) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/9e3e66c553f705de51707c7ddc7f35ce159a8ef1)
<!-- [/FEATURE_TAG] -->

Traverses a RB tree node on the right.

## Definition

**Returns**
Pointer to bpf_rb_node of the right entry, or NULL if given node has no right element.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct bpf_rb_node *bpf_rbtree_right(struct bpf_rb_root *root, struct bpf_rb_node *node)`
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

