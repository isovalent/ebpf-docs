---
title: "Helper Function 'bpf_timer_set_callback'"
description: "This page documents the 'bpf_timer_set_callback' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_timer_set_callback`

<!-- [FEATURE_TAG](bpf_timer_set_callback) -->
[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/b00628b1c7d595ae5b544e059c27b1f5828314b4)
<!-- [/FEATURE_TAG] -->

This helper sets the callback function for a [timer](../concepts/timers.md).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.

<!-- [HELPER_FUNC_DEF] -->
Configure the timer to call _callback_fn_ static function.

### Returns

0 on success. **-EINVAL** if _timer_ was not initialized with bpf_timer_init() earlier. **-EPERM** if _timer_ is in a map that doesn't have any user references. The user space should either hold a file descriptor to a map with timers or pin such map in bpffs. When map is unpinned or file descriptor is closed all timers in the map will be cancelled and freed.

`#!c static long (* const bpf_timer_set_callback)(struct bpf_timer *timer, void *callback_fn) = (void *) 170;`
<!-- [/HELPER_FUNC_DEF] -->

The function passed to `callback_fn` should have the following signature:

`#!c static int callback_fn(void *map, {map key type} *key, {map value type} *value)`

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
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
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
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

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
