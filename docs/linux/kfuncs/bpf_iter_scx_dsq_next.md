---
title: "KFunc 'bpf_iter_scx_dsq_next'"
description: "This page documents the 'bpf_iter_scx_dsq_next' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_scx_dsq_next`

<!-- [FEATURE_TAG](bpf_iter_scx_dsq_next) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/650ba21b131ed1f8ee57826b2c6295a3be221132)
<!-- [/FEATURE_TAG] -->

This function progresses a DSQ iterator.

## Definition

`it`: iterator to progress

**Returns**

The next task. See [`bpf_iter_scx_dsq_new`](bpf_iter_scx_dsq_new.md).

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct task_struct *bpf_iter_scx_dsq_next(struct bpf_iter_scx_dsq *it)`

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

