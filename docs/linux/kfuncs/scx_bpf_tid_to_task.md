---
title: "KFunc 'scx_bpf_tid_to_task'"
description: "This page documents the 'scx_bpf_tid_to_task' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_tid_to_task`

<!-- [FEATURE_TAG](scx_bpf_tid_to_task) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/41e3312861eafba171d9620150aaf2e99165d044)
<!-- [/FEATURE_TAG] -->

Look up a task by its SCX task ID.

## Definition

Unavailable if [`sched_ext_ops.update_idle`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#update_idle) is implemented and [`SCX_OPS_KEEP_BUILTIN_IDLE`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_ops_keep_builtin_idle) is not set.

**Parameters**

`tid`: task ID previously read from [`p->scx.tid`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#struct-sched_ext_entity-tid)

**Returns**

The task with the given `tid`, or `NULL` if no such task exists. The returned pointer is valid until the end of the current RCU read section ([`KF_RCU_PROTECTED`](../concepts/kfuncs.md#kf_rcu_protected)). Requires [`SCX_OPS_TID_TO_TASK`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_ops_tid_to_task) to be set on the root scheduler; otherwise an error is raised and `NULL` returned.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct task_struct *scx_bpf_tid_to_task(u64 tid)`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).

!!! note
	This kfunc is RCU protected. This means that the kfunc can be called from RCU read-side critical section.
	If a program isn't called from RCU read-side critical section, such as sleepable programs, the 
	[`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) and 
	[`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) to protect the calls to such KFuncs.
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

