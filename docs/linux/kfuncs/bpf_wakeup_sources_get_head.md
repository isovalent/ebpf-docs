---
title: "KFunc 'bpf_wakeup_sources_read_unlock'"
description: "This page documents the 'bpf_wakeup_sources_read_unlock' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_wakeup_sources_read_unlock`

<!-- [FEATURE_TAG](bpf_wakeup_sources_read_unlock) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/5ff44955447eb04f77161736ff5729c8c0994f7f)
<!-- [/FEATURE_TAG] -->

Get the head of the wake-up sources list.

## Definition

**Returns**

The head of the wake-up sources list.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void *bpf_wakeup_sources_get_head()`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

