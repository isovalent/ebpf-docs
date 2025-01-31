---
title: "KFunc 'scx_bpf_cpuperf_cur'"
description: "This page documents the 'scx_bpf_cpuperf_cur' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cpuperf_cur`

<!-- [FEATURE_TAG](scx_bpf_cpuperf_cur) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/d86adb4fc0655a0867da811d000df75d2a325ef6)
<!-- [/FEATURE_TAG] -->

This function queries the current relative performance of a CPU.

## Definition

The current performance level of a CPU in relation to the maximum performance available in the system can be calculated as follows:

[`scx_bpf_cpuperf_cap()`](scx_bpf_cpuperf_cap.md) * [`scx_bpf_cpuperf_cur()`](scx_bpf_cpuperf_cur.md) / `SCX_CPUPERF_ONE`

The result is in the range [1, `SCX_CPUPERF_ONE`].

`cpu`: CPU of interest

**Returns**

The current relative performance of `cpu` in relation to its maximum. The return value is in the range [`1`, `SCX_CPUPERF_ONE`].

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 scx_bpf_cpuperf_cur(s32 cpu)`
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

