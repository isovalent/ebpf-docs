---
title: "KFunc 'bpf_cpumask_populate'"
description: "This page documents the 'bpf_cpumask_populate' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_populate`

<!-- [FEATURE_TAG](bpf_cpumask_populate) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/950ad93df2fce70445e655ed2e74f5c1a8653ab2)
<!-- [/FEATURE_TAG] -->

Populate the CPU mask from the contents of a BPF memory region.

## Definition

**Parameters**

`cpumask`: The cpumask being populated.

`src`: The BPF memory holding the bit pattern.

`src__sz`: Length of the BPF memory region in bytes.

**Returns**

 * `0` if the `struct cpumask *` instance was populated successfully.
 * `-EACCES` if the memory region is too small to populate the cpumask.
 * `-EINVAL` if the memory region is not aligned to the size of a long and the architecture does not support efficient unaligned accesses.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_cpumask_populate(struct cpumask *cpumask, void *src, size_t src__sz)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

