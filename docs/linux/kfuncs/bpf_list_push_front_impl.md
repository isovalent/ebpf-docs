---
title: "KFunc 'bpf_list_push_front_impl'"
description: "This page documents the 'bpf_list_push_front_impl' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_list_push_front_impl`

<!-- [FEATURE_TAG](bpf_list_push_front_impl) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/d2dcc67df910dd85253a701b6a5b747f955d28f5)
<!-- [/FEATURE_TAG] -->

Add a new entry to the beginning of the BPF linked list.

## Definition

The `meta` and `off` parameters are rewritten by the verifier, no need for BPF programs to set them

**Returns**

* `0` if the node was successfully added
* `-EINVAL` if the node wasn't added because it's already in a list

<!-- [KFUNC_DEF] -->
`#!c int bpf_list_push_front_impl(struct bpf_list_head *head, struct bpf_list_node *node, void *meta__ign, u64 off)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

