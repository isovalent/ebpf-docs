---
title: "KFunc 'bpf_refcount_acquire_impl'"
description: "This page documents the 'bpf_refcount_acquire_impl' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_refcount_acquire_impl`

<!-- [FEATURE_TAG](bpf_refcount_acquire_impl) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/7c50b1cb76aca4540aa917db5f2a302acddcadff)
<!-- [/FEATURE_TAG] -->

Increment the refcount on a refcounted local kptr, turning the non-owning reference input into an owning reference in the process.

## Definition

The `meta` parameter is rewritten by the verifier, no need for BPF program to set it.

**Returns**

An owning reference to the object pointed to by `kptr`

<!-- [KFUNC_DEF] -->
`#!c void *bpf_refcount_acquire_impl(void *p__refcounted_kptr, void *meta__ign)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.

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
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

