---
title: "KFunc 'bpf_cpumask_first_zero'"
description: "This page documents the 'bpf_cpumask_first_zero' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_first_zero`

<!-- [FEATURE_TAG](bpf_cpumask_first_zero) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Get the index of the first unset bit in the cpumask.

## Definition

Find the index of the first unset bit of the cpumask. A `struct bpf_cpumask`
pointer may be safely passed to this function.

<!-- [KFUNC_DEF] -->
`#!c u32 bpf_cpumask_first_zero(const struct cpumask *cpumask)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

