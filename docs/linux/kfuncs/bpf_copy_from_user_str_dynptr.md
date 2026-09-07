---
title: "KFunc 'bpf_copy_from_user_str_dynptr'"
description: "This page documents the 'bpf_copy_from_user_str_dynptr' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_copy_from_user_str_dynptr`

<!-- [FEATURE_TAG](bpf_copy_from_user_str_dynptr) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/a498ee7576de24b4b0916ce56cf2686e261a29f7)
<!-- [/FEATURE_TAG] -->

Sleepable, copies user-space string into a [dynptr](../concepts/dynptrs.md) for the current task

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_copy_from_user_str_dynptr(const struct bpf_dynptr *dptr, u64 off, u64 size, const void *unsafe_ptr__ign)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

!!! note
    In [:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/531b87d865eb9e625c2e46ec8f06a65a6157ee45) the signature of this kfunc changed from `u32` to `u64` types for `off` and `size`. This may require CO-RE logic to select the correct kfunc.

## Usage

This kfunc can be used to probe a string from kernel memory directly into a [dynptr](../concepts/dynptrs.md) buffer. The max size of which can be dynamic (not known at verification time).

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

