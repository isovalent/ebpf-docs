---
title: "KFunc 'bpf_cpumask_first_and'"
description: "This page documents the 'bpf_cpumask_first_and' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_first_and`

<!-- [FEATURE_TAG](bpf_cpumask_first_and) -->
[:octicons-tag-24: v6.5](https://github.com/torvalds/linux/commit/5ba3a7a851e3ebffc4cb8f052a4581c4d8af3ae3)
<!-- [/FEATURE_TAG] -->

Return the index of the first nonzero bit from the AND of two CPU-masks.

## Definition

Find the index of the first nonzero bit of the AND of two CPU-masks.
`struct bpf_cpumask` pointers may be safely passed to `src1` and `src2`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 bpf_cpumask_first_and(const struct cpumask *src1, const struct cpumask *src2)`
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

