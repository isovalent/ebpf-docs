---
title: "KFunc 'bpf_cpumask_set_cpu'"
description: "This page documents the 'bpf_cpumask_set_cpu' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_set_cpu`

<!-- [FEATURE_TAG](bpf_cpumask_set_cpu) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Set a bit for a CPU in a BPF CPU-mask.

## Definition

`cpu`: The CPU to be set in the CPU-mask.

`cpumask`: The BPF CPU-mask in which a bit is being set.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_cpumask_set_cpu(u32 cpu, struct bpf_cpumask *cpumask)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

