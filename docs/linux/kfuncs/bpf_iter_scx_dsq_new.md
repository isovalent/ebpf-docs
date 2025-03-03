---
title: "KFunc 'bpf_iter_scx_dsq_new'"
description: "This page documents the 'bpf_iter_scx_dsq_new' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_scx_dsq_new`

<!-- [FEATURE_TAG](bpf_iter_scx_dsq_new) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/650ba21b131ed1f8ee57826b2c6295a3be221132)
<!-- [/FEATURE_TAG] -->

This function creates a DSQ iterator.

## Definition

Initialize BPF iterator `it` which can be used with [`bpf_for_each`](../../ebpf-library/libbpf/ebpf/bpf_for_each.md) to walk tasks in the DSQ specified by `dsq_id`. Iteration using `it` only includes tasks which are already queued when this function is invoked.

**Parameters**

`it`: iterator to initialize

`dsq_id`: DSQ to iterate

`flags`: `SCX_DSQ_ITER_*`

**Flags**

`SCX_DSQ_ITER_REV`: iterate in the reverse dispatch order

**Returns**

`0` on success, negative error code on failure

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_iter_scx_dsq_new(struct bpf_iter_scx_dsq *it, u64 dsq_id, u64 flags)`

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

