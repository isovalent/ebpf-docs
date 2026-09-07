---
title: "KFunc 'bpf_list_empty'"
description: "This page documents the 'bpf_list_empty' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_list_empty`

<!-- [FEATURE_TAG](bpf_list_empty) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/745515d386eb5e6891d9f91a92ad15dace3a33ef)
<!-- [/FEATURE_TAG] -->

Check if a list is empty.

## Definition

**Parameters**

`head`: The list to check.

**Returns**

`true` if `head` is an empty list.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool bpf_list_empty(struct bpf_list_head *head)`
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

