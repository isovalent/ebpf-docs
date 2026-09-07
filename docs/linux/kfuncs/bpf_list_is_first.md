---
title: "KFunc 'bpf_list_is_first'"
description: "This page documents the 'bpf_list_is_first' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_list_is_first`

<!-- [FEATURE_TAG](bpf_list_is_first) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/745515d386eb5e6891d9f91a92ad15dace3a33ef)
<!-- [/FEATURE_TAG] -->

Check if a node is the first node of a list.

## Definition

**Parameters**

`head`: The list in which to check for `node__nonown_allowed`

`node__nonown_allowed`: The node we want to know is first or not.

**Returns**

`true` if `node__nonown_allowed` is the first node of `head`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool bpf_list_is_first(struct bpf_list_head *head, struct bpf_list_node *node__nonown_allowed)`
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

