---
title: "Helper Function 'bpf_timer_start'"
description: "This page documents the 'bpf_timer_start' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_timer_start`

<!-- [FEATURE_TAG](bpf_timer_start) -->
[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/b00628b1c7d595ae5b544e059c27b1f5828314b4)
<!-- [/FEATURE_TAG] -->

This helper starts a [timer](../concepts/timers.md).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Set timer expiration N nanoseconds from the current time. The configured callback will be invoked in soft irq context on some cpu and will not repeat unless another bpf_timer_start() is made. In such case the next invocation can migrate to a different cpu. Since struct bpf_timer is a field inside map element the map owns the timer. The bpf_timer_set_callback() will increment refcnt of BPF program to make sure that callback_fn code stays valid. When user space reference to a map reaches zero all timers in a map are cancelled and corresponding program's refcnts are decremented. This is done to make sure that Ctrl-C of a user process doesn't leave any timers running. If map is pinned in bpffs the callback_fn can re-arm itself indefinitely. bpf_map_update/delete_elem() helpers and user space sys_bpf commands cancel and free the timer in the given map element. The map can contain timers that invoke callback_fn-s from different programs. The same callback_fn can serve different timers from different maps if key/value layout matches across maps. Every bpf_timer_set_callback() can have different callback_fn.

_flags_ can be one of:

**BPF_F_TIMER_ABS**

&nbsp;&nbsp;&nbsp;&nbsp;Start the timer in absolute expire value instead of the default relative one.

**BPF_F_TIMER_CPU_PIN**

&nbsp;&nbsp;&nbsp;&nbsp;Timer will be pinned to the CPU of the caller.

&nbsp;&nbsp;&nbsp;&nbsp;

### Returns

0 on success. **-EINVAL** if _timer_ was not initialized with bpf_timer_init() earlier or invalid _flags_ are passed.

`#!c static long (* const bpf_timer_start)(struct bpf_timer *timer, __u64 nsecs, __u64 flags) = (void *) 171;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- TODO use generated list as soon as we can exclude functions from inherited groups -->
<!-- verifier.c excludes tracing programs from using timers -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)

### Example

```c
#include <linux/bpf.h>
#include <time.h>
#include <stdbool.h>
#include <errno.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

struct elem {
	struct bpf_timer t;
};

struct {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__uint(max_entries, 1);
	__type(key, int);
	__type(value, struct elem);
} hmap SEC(".maps");

static int timer_callback(void* hmap, int* key, struct bpf_timer *timer)
{
    	bpf_printk("Callback was invoked do something useful");
	return 0;
}

SEC("cgroup_skb/egress")
int bpf_prog1(void *ctx)
{
	struct bpf_timer *timer;
	int err, key = 0;
	struct elem init;
	struct elem* ele;

	__builtin_memset(&init, 0, sizeof(struct elem));
	bpf_map_update_elem(&hmap, &key, &init, BPF_ANY);

	ele = bpf_map_lookup_elem(&hmap, &key);
	if (!ele)
    	return 1;

	timer = &ele->t;
	err = bpf_timer_init(timer, &hmap, CLOCK_MONOTONIC);
	if (err && err != -EBUSY)
    	return 1;

	bpf_timer_set_callback(timer, timer_callback);
	bpf_timer_start(timer, 0, 0);
	bpf_timer_cancel(timer);

	return 0;
}

char _license[] SEC("license") = "GPL";
```
