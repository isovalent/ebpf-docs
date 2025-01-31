---
title: "KFunc 'scx_bpf_dsq_nr_queued'"
description: "This page documents the 'scx_bpf_dsq_nr_queued' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_nr_queued`

<!-- [FEATURE_TAG](scx_bpf_dsq_nr_queued) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

This function returns the number of queued tasks.

## Definition

**Parameters**

`dsq_id`: id of the DSQ

**Returns**

The number of tasks in the DSQ matching `dsq_id`. If not found, `-ENOENT` is returned.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_dsq_nr_queued(u64 dsq_id)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

