---
title: "KFunc 'bpf_task_from_pid' - eBPF Docs"
description: "This page documents the 'bpf_task_from_pid' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_task_from_pid`

<!-- [FEATURE_TAG](bpf_task_from_pid) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/3f0e6f2b41d35d4446160c745e8f09037447dd8f)
<!-- [/FEATURE_TAG] -->

Find a `struct task_struct` from its pid by looking it up in the root pid namespace idr.

## Definition

If a task is returned, it must either be stored in a map, or released with [`bpf_task_release()`](bpf_task_release.md).

<!-- [KFUNC_DEF] -->
`#!c struct task_struct *bpf_task_from_pid(s32 pid)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
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

