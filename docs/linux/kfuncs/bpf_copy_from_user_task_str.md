---
title: "KFunc 'bpf_copy_from_user_task_str'"
description: "This page documents the 'bpf_copy_from_user_task_str' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_copy_from_user_task_str`

<!-- [FEATURE_TAG](bpf_copy_from_user_task_str) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/f0f8a5b58f78f42ed665f81375a9e99a24853b03)
<!-- [/FEATURE_TAG] -->

Copy a string from an task's address space

## Definition

Copies a `NUL` terminated string from a task's address space to `dst__sz` buffer. If user string is too long this will still ensure zero termination in the `dst__sz` buffer unless buffer size is `0`.

**Parameters**

`dst`: Destination address, in kernel space. This buffer must be at least `dst__sz` bytes long.

`dst__sz`: Maximum number of bytes to copy, includes the trailing `NUL`.

`unsafe_ptr__ign`: Source address in the task's address space.

`tsk`: The task whose address space will be used

`flags`: The only supported flag is `BPF_F_PAD_ZEROS`

**Flags**

`BPF_F_PAD_ZEROS`:  If is set, memset the tail of `dst__sz` to 0 on success and memset all of `dst__sz` on failure.

**Returns**

The number of copied bytes on success including the `NUL` terminator. A negative error code on failure.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_copy_from_user_task_str(void *dst, u32 dst__sz, const void *unsafe_ptr__ign, struct task_struct *tsk, u64 flags)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

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

