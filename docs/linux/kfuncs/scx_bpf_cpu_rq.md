---
title: "KFunc 'scx_bpf_cpu_rq'"
description: "This page documents the 'scx_bpf_cpu_rq' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cpu_rq`

<!-- [FEATURE_TAG](scx_bpf_cpu_rq) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/6203ef73fa5c0358f7960b038628259be1448724)
<!-- [/FEATURE_TAG] -->

This function fetches the rq (runqueue) of a CPU.

## Definition

**Parameters**

`cpu`: CPU of the rq

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct rq *scx_bpf_cpu_rq(s32 cpu)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

