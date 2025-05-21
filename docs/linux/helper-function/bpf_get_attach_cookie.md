---
title: "Helper Function 'bpf_get_attach_cookie'"
description: "This page documents the 'bpf_get_attach_cookie' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_attach_cookie`

<!-- [FEATURE_TAG](bpf_get_attach_cookie) -->
[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/82e6b1eee6a8875ef4eacfd60711cce6965c6b04)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


Get bpf_cookie value provided (optionally) during the program attachment. It might be different for each individual attachment, even if BPF program itself is the same. Expects BPF program context _ctx_ as a first argument.

Supported for the following program types:

- kprobe/uprobe
- tracepoint
- perf_event
- <nospell>fentry/fexit/fmod_ret</nospell>
- LSM

### Returns

Value specified by user at BPF link creation/attachment time or 0, if it was not specified.

`#!c static __u64 (* const bpf_get_attach_cookie)(void *ctx) = (void *) 174;`

## Usage

This is useful for cases when the same BPF program is used for attaching and processing invocation of different tracepoints/kprobes/uprobes in a generic fashion, but such that each invocation is distinguished from each other (e.g., BPF program can look up additional information associated with a specific kernel function without having to rely on function IP lookups). This enables new use cases to be implemented simply and efficiently that previously were possible only through code generation (and thus multiple instances of almost identical BPF program) or compilation at runtime (BCC-style) on target hosts (even more expensive resource-wise). For uprobes it is not even possible in some cases to know function IP before hand (e.g., when attaching to shared library without process ID filtering, in which case base load address is not known for a library).

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md) [:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md) [:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
SEC("lsm/inode_create")
int BPF_PROG(lsm_inode_create){
	__u64 attach_cookie = bpf_get_attach_cookie(ctx);
    	__u64 cgroup_id = bpf_get_current_cgroup_id();
	if (attach_cookie == cgroup_id)
	{
        	// Only traces events from tasks who belong to a specific cgroup
		struct event event = {};
		event.pid = bpf_get_current_pid_tgid() >> 32;
		bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &event, sizeof(event));
	}
	return 0;
}
```
