---
title: "KFunc 'bpf_obj_drop_impl'"
description: "This page documents the 'bpf_obj_drop_impl' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_obj_drop_impl`

<!-- [FEATURE_TAG](bpf_obj_drop_impl) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/ac9f06050a3580cf4076a57a470cd71f12a81171)
<!-- [/FEATURE_TAG] -->

Free an allocated object.

## Definition

All fields of the object that require destruction will be destructed before the storage is freed.

The `meta` parameter is rewritten by the verifier, no need for BPF
program to set it.

<!-- [KFUNC_DEF] -->
`#!c void bpf_obj_drop_impl(void *p__alloc, void *meta__ign)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
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

