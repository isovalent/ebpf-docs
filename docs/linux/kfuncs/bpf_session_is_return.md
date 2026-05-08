---
title: "KFunc 'bpf_session_is_return'"
description: "This page documents the 'bpf_session_is_return' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_session_is_return`

<!-- [FEATURE_TAG](bpf_session_is_return) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/adf46d88ae4b2557f7e2e02547a25fb866935711)
<!-- [/FEATURE_TAG] -->

Check if the bpf program is executed from the exit probe of the kprobe multi link attached in wrapper mode.

## Definition

**Returns**

Returns `true` if the bpf program is executed from the exit probe of the kprobe multi link attached in wrapper mode. It returns `false` otherwise.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool bpf_session_is_return(void *ctx)`
<!-- [/KFUNC_DEF] -->

!!! note
    In [:octicons-tag-24: v7.0](https://github.com/torvalds/linux/commit/8fe4dc4f6456b3d2c9e6f8aeb1f978b7bff0f6c8) the signature of this kfunc has been changed. It used to be `#!c bool bpf_session_is_return()`. [Weak symbols](../../ebpf-library/libbpf/ebpf/__weak.md) and [`bpf_ksym_exists`](../../ebpf-library/libbpf/ebpf/bpf_ksym_exists.md) can be used to write fallback logic for older kernels.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

