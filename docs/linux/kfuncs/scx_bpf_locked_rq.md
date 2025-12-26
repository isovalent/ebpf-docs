---
title: "KFunc 'scx_bpf_locked_rq'"
description: "This page documents the 'scx_bpf_locked_rq' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_locked_rq`

<!-- [FEATURE_TAG](scx_bpf_locked_rq) -->
[:octicons-tag-24: v6.18](https://github.com/torvalds/linux/commit/e0ca169638be12a0a861e3439e6117c58972cd08)
<!-- [/FEATURE_TAG] -->

Return the run queue currently locked by SCX

## Definition

**Returns**

Returns the run queue if a run queue lock is currently held by SCX. Otherwise emits an error and returns `NULL`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct rq *scx_bpf_locked_rq()`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
<!-- [/KFUNC_DEF] -->

## Usage

Most fields in [`scx_bpf_cpu_rq`](scx_bpf_cpu_rq.md) assume that its `rq_lock` is held. Furthermore they become meaningless without run queue lock, too. Make a safer version of [`scx_bpf_cpu_rq`](scx_bpf_cpu_rq.md) that only returns a run queue if we hold run queue lock of that run queue.

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

