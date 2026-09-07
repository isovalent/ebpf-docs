---
title: "KFunc 'scx_bpf_cidperf_cur'"
description: "This page documents the 'scx_bpf_cidperf_cur' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cidperf_cur`

<!-- [FEATURE_TAG](scx_bpf_cidperf_cur) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/5ba0a42423335f76d3e0513df42416c69dc6b742)
<!-- [/FEATURE_TAG] -->

Query the current performance of the CPU at `cid`.

## Definition

cid-addressed equivalent of [`scx_bpf_cpuperf_cur`](scx_bpf_cpuperf_cur.md).

**Parameters**

`cid`: cid of the CPU to query

**Returns**

Current performance of the CPU.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 scx_bpf_cidperf_cur(s32 cid)`
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

