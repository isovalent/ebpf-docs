---
title: "KFunc 'scx_bpf_destroy_dsq'"
description: "This page documents the 'scx_bpf_destroy_dsq' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_destroy_dsq`

<!-- [FEATURE_TAG](scx_bpf_destroy_dsq) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

This function destroys a custom DSQ.

## Definition

Destroy the custom DSQ identified by `dsq_id`. Only DSQs created with [`scx_bpf_create_dsq`](scx_bpf_create_dsq.md) can be destroyed. The caller must ensure that the DSQ is empty and no further tasks are dispatched to it. Ignored if called on a DSQ which doesn't exist. Can be called from any online [`sched_ext_ops`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md) operations.

**Parameters**

`dsq_id`: DSQ to destroy

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_destroy_dsq(u64 dsq_id)`
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

