# KFunc `bpf_cpumask_full`

<!-- [FEATURE_TAG](bpf_cpumask_full) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Check if a cpumask has all bits set.

## Definition

`cpumask`: The cpumask being checked.

Return:
* `true`   - All of the bits in `cpumask` are set.
* `false`  - At least one bit in `cpumask` is cleared.

A `struct bpf_cpumask` pointer may be safely passed to `cpumask`.

<!-- [KFUNC_DEF] -->
`#!c bool bpf_cpumask_full(const struct cpumask *cpumask)`
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

