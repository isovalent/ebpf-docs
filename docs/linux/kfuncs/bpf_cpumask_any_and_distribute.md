---
title: "KFunc 'bpf_cpumask_any_and_distribute'"
description: "This page documents the 'bpf_cpumask_any_and_distribute' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_any_and_distribute`

<!-- [FEATURE_TAG](bpf_cpumask_any_and_distribute) -->
[:octicons-tag-24: v6.5](https://github.com/torvalds/linux/commit/f983be917332ea5e03f689e12c6668be48cb4cfe)
<!-- [/FEATURE_TAG] -->

Return a random set CPU from the AND of two cpumasks.

## Definition

`src1`: The first cpumask.
`src2`: The second cpumask.

Return:
* A random set bit within [0, num_cpus) from the AND of two cpumasks, if at
  least one bit is set.
* >= num_cpus if no bit is set.

`struct bpf_cpumask` pointers may be safely passed to `src1` and `src2`.

<!-- [KFUNC_DEF] -->
`#!c u32 bpf_cpumask_any_and_distribute(const struct cpumask *src1, const struct cpumask *src2)`
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

