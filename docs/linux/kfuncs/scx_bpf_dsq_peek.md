---
title: "KFunc 'scx_bpf_dsq_peek'"
description: "This page documents the 'scx_bpf_dsq_peek' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_peek`

<!-- [FEATURE_TAG](scx_bpf_dsq_peek) -->
[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/44f5c8ec5b9ad8ed4ade08d727f803b2bb07f1c3)
<!-- [/FEATURE_TAG] -->

Lock-less peek at the first element.

## Definition

Read the first element in the DSQ. This is semantically equivalent to using the DSQ iterator, but is lock-free. Of course, like any lock-less operation, this provides only a point-in-time snapshot, and the contents may change by the time any subsequent locking operation reads the queue.

**Parameters**

`dsq_id`: DSQ to examine.

**Returns**

The pointer, or `NULL` indicates an empty queue OR internal error.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct task_struct *scx_bpf_dsq_peek(u64 dsq_id)`

!!! note
	This kfunc is RCU protected. This means that the kfunc can be called from RCU read-side critical section.
	If a program isn't called from RCU read-side critical section, such as sleepable programs, the 
	[`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) and 
	[`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) to protect the calls to such KFuncs.

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
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

