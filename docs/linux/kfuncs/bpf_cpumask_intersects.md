---
title: "KFunc 'bpf_cpumask_intersects'"
description: "This page documents the 'bpf_cpumask_intersects' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_intersects`

<!-- [FEATURE_TAG](bpf_cpumask_intersects) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Check two cpumasks for overlap.

## Definition

`src1`: The first input.
`src2`: The second input.

Return:
* `true`   - `src1` and `src2` have at least one of the same bits set.
* `false`  - `src1` and `src2` don't have any of the same bits set.

`struct bpf_cpumask` pointers may be safely passed to `src1` and `src2`.

<!-- [KFUNC_DEF] -->
`#!c bool bpf_cpumask_intersects(const struct cpumask *src1, const struct cpumask *src2)`
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

