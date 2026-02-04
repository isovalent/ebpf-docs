---
title: "KFunc 'bpf_list_back'"
description: "This page documents the 'bpf_list_back' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_list_back`

<!-- [FEATURE_TAG](bpf_list_back) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/fb5b480205bad3936b054b86f7c9d2bd7835caac)
<!-- [/FEATURE_TAG] -->

Traverses the linked list backwards.

## Definition

**Returns**
Pointer to bpf_list_node of previous entry, or NULL if list given node has no previous element.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct bpf_list_node *bpf_list_back(struct bpf_list_head *head)`
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

