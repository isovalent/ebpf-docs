# Helper function `bpf_get_task_stack`

<!-- [FEATURE_TAG](bpf_get_task_stack) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/fa28dcb82a38f8e3993b0fae9106b1a80b59e4f0)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Return a user or a kernel stack in bpf program provided buffer.
To achieve this, the helper needs `task`, which is a valid
pointer to `struct task_struct`. To store the stacktrace, the
bpf program provides `buf` with a nonnegative `size`.

The last argument, `flags`, holds the number of stack frames to
skip (from 0 to 255), masked with
`BPF_F_SKIP_FIELD_MASK`. The next bits can be used to set
the following flags:

`BPF_F_USER_STACK`
Collect a user space stack instead of a kernel stack.
`BPF_F_USER_BUILD_ID`
Collect buildid+offset instead of ips for user stack,
only valid if `BPF_F_USER_STACK` is also specified.

`bpf_get_task_stack`\ () can collect up to
`PERF_MAX_STACK_DEPTH` both kernel and user frames, subject
to sufficient large buffer size. Note that
this limit can be controlled with the `sysctl` program, and
that it should be manually increased in order to profile long
user stacks (such as stacks for Java programs). To do so, use:

::

# sysctl kernel.perf_event_max_stack=<new value>


**Returns**
The non-negative copied `buf` length equal to or less than
`size` on success, or a negative error in case of failure.

`#!c static long (*bpf_get_task_stack)(struct task_struct *task, void *buf, __u32 size, __u64 flags) = (void *) 141;`
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
