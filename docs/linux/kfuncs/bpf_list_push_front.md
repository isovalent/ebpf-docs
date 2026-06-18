---
title: "KFunc 'bpf_list_push_front'"
description: "This page documents the 'bpf_list_push_front' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_list_push_front`

<!-- [FEATURE_TAG](bpf_list_push_front) -->
[:octicons-tag-24: 7.1](https://github.com/torvalds/linux/commit/d457072576a6a60ba853b1d815f123da57b48021)
<!-- [/FEATURE_TAG] -->

Add a new entry to the beginning of the BPF linked list.

## Definition

**Parameters**

`head`: Head of the linked list.
`node`: The node to add to the list.

**Returns**

* `0` if the node was successfully added
* `-EINVAL` if the node wasn't added because it's already in a list

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_list_push_front(struct bpf_list_head *head, struct bpf_list_node *node)`
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

