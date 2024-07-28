---
title: "Helper Function 'bpf_get_stackid'"
description: "This page documents the 'bpf_get_stackid' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_stackid`

<!-- [FEATURE_TAG](bpf_get_stackid) -->
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/d5a3b1f691865be576c2bffa708549b8cdccda19)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Walk a user or a kernel stack and return its id. To achieve this, the helper needs _ctx_, which is a pointer to the context on which the tracing program is executed, and a pointer to a _map_ of type **BPF_MAP_TYPE_STACK_TRACE**.

The last argument, _flags_, holds the number of stack frames to skip (from 0 to 255), masked with **BPF_F_SKIP_FIELD_MASK**. The next bits can be used to set a combination of the following flags:

**BPF_F_USER_STACK**

&nbsp;&nbsp;&nbsp;&nbsp;Collect a user space stack instead of a kernel stack.

**BPF_F_FAST_STACK_CMP**

&nbsp;&nbsp;&nbsp;&nbsp;Compare stacks by hash only.

**BPF_F_REUSE_STACKID**

&nbsp;&nbsp;&nbsp;&nbsp;If two different stacks hash into the same _stackid_, discard the old one.

The stack id retrieved is a 32 bit long integer handle which can be further combined with other data (including other stack ids) and used as a key into maps. This can be useful for generating a variety of graphs (such as flame graphs or off-cpu graphs).

For walking a stack, this helper is an improvement over **bpf_probe_read**(), which can be used with unrolled loops but is not efficient and consumes a lot of eBPF instructions. Instead, **bpf_get_stackid**() can collect up to **PERF_MAX_STACK_DEPTH** both kernel and user frames. Note that this limit can be controlled with the **sysctl** program, and that it should be manually increased in order to profile long user stacks (such as stacks for Java programs). To do so, use:

```
# sysctl kernel.perf_event_max_stack=<new value>
```

### Returns

The positive or null stack id on success, or a negative error in case of failure.

`#!c static long (* const bpf_get_stackid)(void *ctx, void *map, __u64 flags) = (void *) 27;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

Call `bpf_get_stackid` to retrieve the stack id of the context in which the program is running, specifying as arguments:

* *ctx*, the pointer to the current context on which the program is executing
* *bpf_map*, the pointer to a map of type [`BPF_MAP_TYPE_STACK_TRACE`](../map-type/BPF_MAP_TYPE_STACK_TRACE.md)
* *flags*, the flags bitmap

```c
long bpf_get_stackid(void *ctx, struct bpf_map *map, u64 flags)
```

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
#include <bpf/bpf_helpers.h>

struct {
	__uint(type, BPF_MAP_TYPE_STACK_TRACE);
	__uint(key_size, sizeof(u32));
	__uint(value_size, PERF_MAX_STACK_DEPTH * sizeof(u64));
	__uint(max_entries, 10000);
} stack_traces SEC(".maps");

SEC("perf_event")
int print_stack_ids(struct bpf_perf_event_data *ctx)
{
	char fmt[] = "kern_stack_id=%d user_stack_id=%d";
	
	kern_stack_id = bpf_get_stackid(ctx, &stack_traces, 0);
	user_stack_id = bpf_get_stackid(ctx, &stack_traces, 0 | BPF_F_USER_STACK);
	
	if kern_stack_id >= 0 && user_stack_id >=0 {
		bpf_trace_printk(fmt, sizeof(fmt), kern_stack_id, user_stack_id);
	}
}

char _license[] SEC("license") = "GPL";
```

Complete examples in the Linux source bpf samples:

  * [`samples/bpf/offwaketime.bpf.c`](https://github.com/torvalds/linux/blob/e8f897f4afef0031fe618a8e94127a0934896aba/samples/bpf/offwaketime.bpf.c)
  * [`samples/bpf/trace_event_kern.c`](https://github.com/torvalds/linux/blob/e8f897f4afef0031fe618a8e94127a0934896aba/samples/bpf/trace_event_kern.c)

