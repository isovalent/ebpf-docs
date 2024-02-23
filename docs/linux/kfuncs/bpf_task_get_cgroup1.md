---
title: "KFunc 'bpf_task_get_cgroup1'"
description: "This page documents the 'bpf_task_get_cgroup1' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_task_get_cgroup1`

<!-- [FEATURE_TAG](bpf_task_get_cgroup1) -->
[:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/fe977716b40cb98cf9c91a66454adf3dc2f8c59a)
<!-- [/FEATURE_TAG] -->

Acquires the associated cgroup of a task within a specific cgroup1 hierarchy. 

## Definition

The cgroup1 hierarchy is identified by its hierarchy ID.

**Returns**

On success, the cgroup is returen. On failure, NULL is returned.

<!-- [KFUNC_DEF] -->
`#!c struct cgroup *bpf_task_get_cgroup1(struct task_struct *task, int hierarchy_id)`

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

