---
title: "KFunc 'bpf_task_from_vpid'"
description: "This page documents the 'bpf_task_from_vpid' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_task_from_vpid`

<!-- [FEATURE_TAG](bpf_task_from_vpid) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/675c3596ff32c040d1dd2e28dd57e83e634b9f60)
<!-- [/FEATURE_TAG] -->

Find a struct task_struct from its `vpid`.

## Definition

Find a struct `task_struct` from its `vpid` by looking it up in the `pid` namespace of the current task.

**Parameters**

`vpid`: The `vpid` of the task being looked up.

**Returns**

A pointer to a task struct or `NULL`. If a task is returned, it must either be stored in a map, or released with [`bpf_task_release`](bpf_task_release.md).

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct task_struct *bpf_task_from_vpid(s32 vpid)`

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
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

