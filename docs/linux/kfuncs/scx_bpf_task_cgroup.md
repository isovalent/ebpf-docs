---
title: "KFunc 'scx_bpf_task_cgroup'"
description: "This page documents the 'scx_bpf_task_cgroup' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_task_cgroup`

<!-- [FEATURE_TAG](scx_bpf_task_cgroup) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)
<!-- [/FEATURE_TAG] -->

This function returns the sched cGroup of a task.

## Definition

`p->sched_task_group->css.cgroup` represents the cgroup `p` is associated with from the scheduler's POV. SCX operations should use this function to determine `p`'s current cgroup as, unlike following `p->cgroups`, `p->sched_task_group` is protected by `p`'s rq lock and thus atomic w.r.t. all rq-locked operations. Can be called on the parameter tasks of rq-locked operations. The restriction guarantees that `p`'s rq is locked by the caller.

**Parameters**

`p`: task of interest

**Returns**

The sched cGroup of a task

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct cgroup *scx_bpf_task_cgroup(struct task_struct *p)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

