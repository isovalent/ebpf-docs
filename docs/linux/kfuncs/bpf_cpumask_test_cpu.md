---
title: "KFunc 'bpf_cpumask_test_cpu'"
description: "This page documents the 'bpf_cpumask_test_cpu' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_test_cpu`

<!-- [FEATURE_TAG](bpf_cpumask_test_cpu) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Test whether a CPU is set in a cpumask.

## Definition

`cpu`: The CPU being queried for.

`cpumask`: The cpumask being queried for containing a CPU.

Return:
* `true`  - `cpu` is set in the cpumask
* `false` - `cpu` was not set in the cpumask, or `cpu` is an invalid cpu.


<!-- [KFUNC_DEF] -->
`#!c bool bpf_cpumask_test_cpu(u32 cpu, const struct cpumask *cpumask)`
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

