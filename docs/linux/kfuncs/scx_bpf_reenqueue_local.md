---
title: "KFunc 'scx_bpf_reenqueue_local'"
description: "This page documents the 'scx_bpf_reenqueue_local' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_reenqueue_local`

<!-- [FEATURE_TAG](scx_bpf_reenqueue_local) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/245254f7081dbe1c8da54675d0e4ddbe74cee61b)
<!-- [/FEATURE_TAG] -->

This function re-enqueues tasks on a local DSQ.

## Definition

Iterate over all of the tasks currently enqueued on the local DSQ of the caller's CPU, and re-enqueue them in the BPF scheduler.

**Returns**

The number of processed tasks. Can only be called from `ops.cpu_release()`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 scx_bpf_reenqueue_local()`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

