# Helper function `bpf_find_vma`

<!-- [FEATURE_TAG](bpf_find_vma) -->
[:octicons-tag-24: v5.17](https://github.com/torvalds/linux/commit/7c7e3d31e7856a8260a254f8c71db416f7f9f5a1)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Find vma of _task_ that contains _addr_, call _callback_fn_ function with _task_, _vma_, and _callback_ctx_. The _callback_fn_ should be a static function and the _callback_ctx_ should be a pointer to the stack. The _flags_ is used to control certain aspects of the helper. Currently, the _flags_ must be 0.

The expected callback signature is

long (\_callback_fn)(struct task_struct \_task, struct vm_area_struct \_vma, void \_callback_ctx);



### Returns

0 on success. **-ENOENT** if _task->mm_ is NULL, or no vma contains _addr_. **-EBUSY** if failed to try lock mmap_lock. **-EINVAL** for invalid **flags**.

`#!c static long (*bpf_find_vma)(struct task_struct *task, __u64 addr, void *callback_fn, void *callback_ctx, __u64 flags) = (void *) 180;`
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
