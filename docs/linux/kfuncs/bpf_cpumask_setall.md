---
title: "KFunc 'bpf_cpumask_setall' - eBPF Docs"
description: "This page documents the 'bpf_cpumask_setall' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_setall`

<!-- [FEATURE_TAG](bpf_cpumask_setall) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Set all of the bits in a BPF cpumask.

## Definition

`cpumask`: The BPF cpumask having all of its bits set.

<!-- [KFUNC_DEF] -->
`#!c void bpf_cpumask_setall(struct bpf_cpumask *cpumask)`
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

