---
title: "KFunc 'scx_bpf_cpu_curr'"
description: "This page documents the 'scx_bpf_cpu_curr' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cpu_curr`

<!-- [FEATURE_TAG](scx_bpf_cpu_curr) -->
[:octicons-tag-24: v6.18](https://github.com/torvalds/linux/commit/20b158094a1adc9bbfdcc41780059b5cd8866ad8)
<!-- [/FEATURE_TAG] -->

Return remote CPU's current task

## Definition

Callers must hold RCU read lock ([`KF_RCU`](../concepts/kfuncs.md#kf_rcu)).

**Parameters**

`cpu`: CPU of interest

**Returns**

The remote CPU's current task

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct task_struct *scx_bpf_cpu_curr(s32 cpu)`

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

`scx_bpf_cpu_curr` Provides a way for scx schedulers to check the current task of a remote run queue without assuming its lock is held.

Before its introduction, many scx schedulers make use of [`scx_bpf_cpu_rq`](scx_bpf_cpu_rq.md) to check a remote current cpu (for example to see if it should be preempted). This is problematic because [`scx_bpf_cpu_rq`](scx_bpf_cpu_rq.md) provides access to all fields of struct run queue, most of which aren't safe to use without holding the associated run queue lock.

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

