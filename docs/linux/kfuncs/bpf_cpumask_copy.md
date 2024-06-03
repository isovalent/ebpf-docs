---
title: "KFunc 'bpf_cpumask_copy'"
description: "This page documents the 'bpf_cpumask_copy' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_copy`

<!-- [FEATURE_TAG](bpf_cpumask_copy) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Copy the contents of a cpumask into a BPF cpumask.

## Definition

`dst`: The BPF cpumask being copied into.
`src`: The cpumask being copied.

A `struct bpf_cpumask` pointer may be safely passed to `src`.

<!-- [KFUNC_DEF] -->
`#!c void bpf_cpumask_copy(struct bpf_cpumask *dst, const struct cpumask *src)`
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

