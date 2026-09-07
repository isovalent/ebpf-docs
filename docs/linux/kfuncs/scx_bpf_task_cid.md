---
title: "KFunc 'scx_bpf_task_cid'"
description: "This page documents the 'scx_bpf_task_cid' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_task_cid`

<!-- [FEATURE_TAG](scx_bpf_task_cid) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/5ba0a42423335f76d3e0513df42416c69dc6b742)
<!-- [/FEATURE_TAG] -->

Get the cid a task is currently associated with.

## Definition

cid-addressed equivalent of [`scx_bpf_task_cpu`](scx_bpf_task_cpu.md). `task_cpu(p)` is always a valid cpu, so this is just a table lookup. 

**Parameters**

`p`: task of interest

**Returns**

the cid a task is currently associated with. Return -EINVAL if called from a non-SCX program before any scheduler has ever been enabled.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_task_cid(const struct task_struct *p)`
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

