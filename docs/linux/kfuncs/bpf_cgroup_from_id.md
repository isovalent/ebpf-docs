---
title: "KFunc 'bpf_cgroup_from_id'"
description: "This page documents the 'bpf_cgroup_from_id' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_cgroup_from_id`

<!-- [FEATURE_TAG](bpf_cgroup_from_id) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/332ea1f697be148bd5e66475d82b5ecc5084da65)
<!-- [/FEATURE_TAG] -->

Find a cGroup from its ID.

## Definition

cGroup returned by this kfunc which is not subsequently stored in a map, must be released by calling [`bpf_cgroup_release()`](bpf_cgroup_release.md).

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct cgroup *bpf_cgroup_from_id(u64 cgid)`

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
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

