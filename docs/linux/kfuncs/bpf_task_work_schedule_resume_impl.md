---
title: "KFunc 'bpf_task_work_schedule_resume_impl'"
description: "This page documents the 'bpf_task_work_schedule_resume_impl' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_task_work_schedule_resume_impl`

<!-- [FEATURE_TAG](bpf_task_work_schedule_resume_impl) -->
[:octicons-tag-24: v6.18](https://github.com/torvalds/linux/commit/38aa7003e369802f81a078f6673d10d97013f04f)
<!-- [/FEATURE_TAG] -->

Schedule BPF callback using `task_work_add` with `TWA_RESUME` mode

## Definition

**Parameters**

`task`: Task struct for which callback should be scheduled

`tw`: Pointer to `struct bpf_task_work` in BPF map value for internal bookkeeping

`map__map`: map that embeds `struct bpf_task_work` in the values

`callback`: pointer to BPF subprogram to call

`aux__prog`: user should pass `NULL`

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_task_work_schedule_resume_impl(struct task_struct *task, struct bpf_task_work *tw, void *map__map, bpf_task_work_callback_t callback, void *aux__prog)`
<!-- [/KFUNC_DEF] -->

`#!c typedef int (*bpf_task_work_callback_t)(struct bpf_map *map, void *key, void *value);`

## Usage

This kfunc allows a BPF program that is being executed in a restricted context such as a Non Mask-able Interrupt (NMI) to schedule a callback on a task. This callback will be executed in a more permissible context (sleepable context) before that task resumes.

This is mostly useful for tools such as profiles. When a program is triggered in an NMI, the program and any helper/kfunc it executes is unable to sleep/wait or page fault. This means that some actions like reading userspace memory or even updating map values may fail. So by scheduling a callback you can do more things in the permissive context, while still passing info from the original execution context via a map value.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
- [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

Example from [`tools/testing/selftests/bpf/progs/task_work.c`](https://github.com/torvalds/linux/blob/b927546677c876e26eba308550207c2ddf812a43/tools/testing/selftests/bpf/progs/task_work.c)

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2025 Meta Platforms, Inc. and affiliates. */

#include <vmlinux.h>
#include <string.h>
#include <stdbool.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "bpf_misc.h"
#include "errno.h"

char _license[] SEC("license") = "GPL";

const void *user_ptr = NULL;

struct elem {
	char data[128];
	struct bpf_task_work tw;
};

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(map_flags, BPF_F_NO_PREALLOC);
	__uint(max_entries, 1);
	__type(key, int);
	__type(value, struct elem);
} hmap SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__uint(max_entries, 1);
	__type(key, int);
	__type(value, struct elem);
} arrmap SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_LRU_HASH);
	__uint(max_entries, 1);
	__type(key, int);
	__type(value, struct elem);
} lrumap SEC(".maps");

static int process_work(struct bpf_map *map, void *key, void *value)
{
	struct elem *work = value;

	bpf_copy_from_user_str(work->data, sizeof(work->data), (const void *)user_ptr, 0);
	return 0;
}

int key = 0;

SEC("perf_event")
int oncpu_hash_map(struct pt_regs *args)
{
	struct elem empty_work = { .data = { 0 } };
	struct elem *work;
	struct task_struct *task;
	int err;

	task = bpf_get_current_task_btf();
	err = bpf_map_update_elem(&hmap, &key, &empty_work, BPF_NOEXIST);
	if (err)
		return 0;
	work = bpf_map_lookup_elem(&hmap, &key);
	if (!work)
		return 0;

	bpf_task_work_schedule_resume_impl(task, &work->tw, &hmap, process_work, NULL);
	return 0;
}

SEC("perf_event")
int oncpu_array_map(struct pt_regs *args)
{
	struct elem *work;
	struct task_struct *task;

	task = bpf_get_current_task_btf();
	work = bpf_map_lookup_elem(&arrmap, &key);
	if (!work)
		return 0;
	bpf_task_work_schedule_signal_impl(task, &work->tw, &arrmap, process_work, NULL);
	return 0;
}

SEC("perf_event")
int oncpu_lru_map(struct pt_regs *args)
{
	struct elem empty_work = { .data = { 0 } };
	struct elem *work;
	struct task_struct *task;
	int err;

	task = bpf_get_current_task_btf();
	work = bpf_map_lookup_elem(&lrumap, &key);
	if (work)
		return 0;
	err = bpf_map_update_elem(&lrumap, &key, &empty_work, BPF_NOEXIST);
	if (err)
		return 0;
	work = bpf_map_lookup_elem(&lrumap, &key);
	if (!work || work->data[0])
		return 0;
	bpf_task_work_schedule_resume_impl(task, &work->tw, &lrumap, process_work, NULL);
	return 0;
}
```
