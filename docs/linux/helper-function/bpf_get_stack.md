# Helper function `bpf_get_stack`

<!-- [FEATURE_TAG](bpf_get_stack) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/c195651e565ae7f41a68acb7d4aa7390ad215de1)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return a user or a kernel stack in bpf program provided buffer. To achieve this, the helper needs _ctx_, which is a pointer to the context on which the tracing program is executed. To store the stacktrace, the bpf program provides _buf_ with a nonnegative _size_.

The last argument, _flags_, holds the number of stack frames to skip (from 0 to 255), masked with **BPF_F_SKIP_FIELD_MASK**. The next bits can be used to set the following flags:

**BPF_F_USER_STACK**

&nbsp;&nbsp;&nbsp;&nbsp;Collect a user space stack instead of a kernel stack.

**BPF_F_USER_BUILD_ID**

&nbsp;&nbsp;&nbsp;&nbsp;Collect (build_id, file_offset) instead of ips for user stack, only valid if **BPF_F_USER_STACK** is also specified.

&nbsp;&nbsp;&nbsp;&nbsp;_file_offset_ is an offset relative to the beginning of the executable or shared object file backing the vma which the _ip_ falls in. It is _not_ an offset relative to that object's base address. Accordingly, it must be adjusted by adding (sh_addr - sh_offset), where sh_{addr,offset} correspond to the executable section containing _file_offset_ in the object, for comparisons to symbols' st_value to be valid.

**bpf_get_stack**() can collect up to **PERF_MAX_STACK_DEPTH** both kernel and user frames, subject to sufficient large buffer size. Note that this limit can be controlled with the **sysctl** program, and that it should be manually increased in order to profile long user stacks (such as stacks for Java programs). To do so, use:

```
# sysctl kernel.perf_event_max_stack=<new value>
```

### Returns

The non-negative copied _buf_ length equal to or less than _size_ on success, or a negative error in case of failure.

`#!c static long (*bpf_get_stack)(void *ctx, void *buf, __u32 size, __u64 flags) = (void *) 67;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
