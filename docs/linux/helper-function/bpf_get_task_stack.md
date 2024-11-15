---
title: "Helper Function 'bpf_get_task_stack'"
description: "This page documents the 'bpf_get_task_stack' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_task_stack`

<!-- [FEATURE_TAG](bpf_get_task_stack) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/fa28dcb82a38f8e3993b0fae9106b1a80b59e4f0)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return a user or a kernel stack in bpf program provided buffer. Note: the user stack will only be populated if the _task_ is the current task; all other tasks will return -EOPNOTSUPP. To achieve this, the helper needs _task_, which is a valid BTF pointer to **struct task_struct**, see [bpf_get_current_task_btf](../helper-function/bpf_get_current_task_btf.md) for more information. To store the stacktrace, the bpf program provides _buf_ with a nonnegative _size_.

The last argument, _flags_, holds the number of stack frames to skip (from 0 to 255), masked with **BPF_F_SKIP_FIELD_MASK**. The next bits can be used to set the following flags:

**BPF_F_USER_STACK**

&nbsp;&nbsp;&nbsp;&nbsp;Collect a user space stack instead of a kernel stack. The _task_ must be the current task.

**BPF_F_USER_BUILD_ID**

&nbsp;&nbsp;&nbsp;&nbsp;Collect buildid+offset instead of ips for user stack, only valid if **BPF_F_USER_STACK** is also specified.

**bpf_get_task_stack**() can collect up to **PERF_MAX_STACK_DEPTH** both kernel and user frames, subject to sufficient large buffer size. Note that this limit can be controlled with the **sysctl** program, and that it should be manually increased in order to profile long user stacks (such as stacks for Java programs). To do so, use:

```
# sysctl kernel.perf_event_max_stack=<new value>
```

### Returns

The non-negative copied _buf_ length equal to or less than _size_ on success, or a negative error in case of failure.

`#!c static long (* const bpf_get_task_stack)(struct task_struct *task, void *buf, __u32 size, __u64 flags) = (void *) 141;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
