# KFunc `bpf_cpumask_weight`

<!-- [FEATURE_TAG](bpf_cpumask_weight) -->
[:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/a6de18f310a511278c1ff16b96eb2d500eada725)
<!-- [/FEATURE_TAG] -->

Return the number of bits in cpumask.

## Definition

`cpumask`: The cpumask being queried.

Count the number of set bits in the given cpumask.

<!-- [KFUNC_DEF] -->
`#!c u32 bpf_cpumask_weight(const struct cpumask *cpumask)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../../program-types/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../../program-types/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_TRACING](../../program-types/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

