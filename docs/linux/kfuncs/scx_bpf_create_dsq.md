---
title: "KFunc 'scx_bpf_create_dsq'"
description: "This page documents the 'scx_bpf_create_dsq' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_create_dsq`

<!-- [FEATURE_TAG](scx_bpf_create_dsq) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

This function creates a custom DSQ.

## Definition

Create a custom DSQ identified by `dsq_id`. Can be called from any sleepable scx callback, and any [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md) prog.

**Parameters**

`dsq_id`: DSQ to create

`node`: NUMA node to allocate from

**Returns**

`0` on success, negative error code on failure

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_create_dsq(u64 dsq_id, s32 node)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

