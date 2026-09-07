---
title: "KFunc 'bpf_list_add'"
description: "This page documents the 'bpf_list_add' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_list_add`

<!-- [FEATURE_TAG](bpf_list_add) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/a3493ca504f16877bf29a123f27835c3f841a05f)
<!-- [/FEATURE_TAG] -->

Inserts `new` after `prev` in the BPF linked list.

## Definition

**Parameters**

`head`: The list head to which `new` should be added.

`new`: The node to be added to the list.

`prev__nonown_allowed`: The node after which `new` should be added.

**Returns**

Pointer to bpf_list_node of that was added, or `NULL` if it was not added to the list.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_list_add(struct bpf_list_head *head, struct bpf_list_node *new, struct bpf_list_node *prev__nonown_allowed)`
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

