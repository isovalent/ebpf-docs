# KFunc `bpf_cpumask_subset`

<!-- [FEATURE_TAG](bpf_cpumask_subset) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Check if a cpumask is a subset of another.

## Definition

`src1`: The first cpumask being checked as a subset.
`src2`: The second cpumask being checked as a superset.

Return:
* `true`   - All of the bits of `src1` are set in `src2`.
* `false`  - At least one bit in `src1` is not set in `src2`.

`struct bpf_cpumask` pointers may be safely passed to `src1` and `src2`.

<!-- [KFUNC_DEF] -->
`#!c bool bpf_cpumask_subset(const struct cpumask *src1, const struct cpumask *src2)`
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

