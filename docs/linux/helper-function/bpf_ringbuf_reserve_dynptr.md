---
title: "Helper Function 'bpf_ringbuf_reserve_dynptr'"
description: "This page documents the 'bpf_ringbuf_reserve_dynptr' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_ringbuf_reserve_dynptr`

<!-- [FEATURE_TAG](bpf_ringbuf_reserve_dynptr) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/bc34dee65a65e9c920c420005b8a43f2a721a458)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Reserve _size_ bytes of payload in a ring buffer _ringbuf_ through the dynptr interface. _flags_ must be 0.

Please note that a corresponding bpf_ringbuf_submit_dynptr or bpf_ringbuf_discard_dynptr must be called on _ptr_, even if the reservation fails. This is enforced by the verifier.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_ringbuf_reserve_dynptr)(void *ringbuf, __u32 size, __u64 flags, struct bpf_dynptr *ptr) = (void *) 198;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

struct {
	__uint(type, BPF_MAP_TYPE_RINGBUF);
	__uint(max_entries, 256 * 1024);
} ringbuf SEC(".maps");

SEC("tp/syscalls/sys_enter_openat")
int bpf_prog1(void *ctx)
{
	char write_data[64] = "hello there, world!!";
	char read_data[64] = {};
	struct bpf_dynptr ptr;
	int i;
	int local_err = 0;

	local_err = bpf_ringbuf_reserve_dynptr(&ringbuf, sizeof(write_data), 0, &ptr);
	if (local_err < 0)
    	goto discard;

	/* Write data into the dynptr */
	local_err = bpf_dynptr_write(&ptr, 0, write_data, sizeof(write_data), 0);
	if (local_err) {
    	goto discard;
	}

	/* Read the data that was written into the dynptr */
	local_err = bpf_dynptr_read(read_data, sizeof(read_data), &ptr, 0, 0);
	if (local_err) {
    	goto discard;
	}

	/* Ensure the data we read matches the data we wrote */
	for (i = 0; i < sizeof(read_data); i++) {
    	if (read_data[i] != write_data[i]) {
        	break;
    	}
	}
	bpf_printk("Read matches write, dynptr API works");
	bpf_ringbuf_submit_dynptr(&ptr, 0);
	return 0;

discard:
	bpf_ringbuf_discard_dynptr(&ptr, 0);
	return 0;
}
```
