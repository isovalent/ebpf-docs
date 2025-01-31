---
title: "KFunc 'bpf_cpumask_release'"
description: "This page documents the 'bpf_cpumask_release' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_cpumask_release`

<!-- [FEATURE_TAG](bpf_cpumask_release) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/516f4d3397c9e90f4da04f59986c856016269aa1)
<!-- [/FEATURE_TAG] -->

Release a previously acquired BPF CPU-mask.

## Definition

Releases a previously acquired reference to a BPF CPU-mask. When the final
reference of the BPF CPU-mask has been released, it is subsequently freed in
an RCU callback in the BPF memory allocator.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_cpumask_release(struct bpf_cpumask *cpumask)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
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

