---
title: "Helper Function 'bpf_copy_from_user_task'"
description: "This page documents the 'bpf_copy_from_user_task' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_copy_from_user_task`

<!-- [FEATURE_TAG](bpf_copy_from_user_task) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/376040e47334c6dc6a939a32197acceb00fe4acf)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Read _size_ bytes from user space address _user_ptr_ in _tsk_'s address space, and stores the data in _dst_. _flags_ is not used yet and is provided for future extensibility. This helper can only be used by sleepable programs.

### Returns

0 on success, or a negative error in case of failure. On error _dst_ buffer is zeroed out.

`#!c static long (* const bpf_copy_from_user_task)(void *dst, __u32 size, const void *user_ptr, struct task_struct *tsk, __u64 flags) = (void *) 191;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
