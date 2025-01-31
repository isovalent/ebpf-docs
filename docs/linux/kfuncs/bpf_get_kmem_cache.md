---
title: "KFunc 'bpf_get_kmem_cache'"
description: "This page documents the 'bpf_get_kmem_cache' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_get_kmem_cache`

<!-- [FEATURE_TAG](bpf_get_kmem_cache) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/a992d7a3979120fbd7c13435d27b3da8d9ed095a)
<!-- [/FEATURE_TAG] -->

This function returns slab cache information from a virtual address of a slab object.

## Definition

This function returns slab cache information from a virtual address of a slab object.

It doesn't grab a reference count of the kmem_cache so the caller is responsible to manage the access. The returned point is marked as `PTR_UNTRUSTED`.

**Parameters**

`addr`: virtual address of the slab object

**Returns**

A valid `kmem_cache` pointer, otherwise `NULL`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct kmem_cache *bpf_get_kmem_cache(u64 addr)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc has many possible use cases. One example is its usage by perf to resolve names of slab objects, when used in combination with a slab allocator iterator(`iter/kmem_cache`). see: [commit](https://github.com/torvalds/linux/commit/0c631ef07c96536a66d8168dc7e176de5fa82878)

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
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2024 Google */
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "bpf_experimental.h"

char _license[] SEC("license") = "GPL";

#define SLAB_NAME_MAX  32

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(key_size, sizeof(void *));
	__uint(value_size, SLAB_NAME_MAX);
	__uint(max_entries, 1);
} slab_hash SEC(".maps");

extern struct kmem_cache *bpf_get_kmem_cache(u64 addr) __ksym;

/* Result, will be checked by userspace */
int task_struct_found;

SEC("raw_tp/bpf_test_finish")
int BPF_PROG(check_task_struct)
{
	u64 curr = bpf_get_current_task();
	struct kmem_cache *s;
	char *name;

	s = bpf_get_kmem_cache(curr);
	if (s == NULL) {
		task_struct_found = -1;
		return 0;
	}
	name = bpf_map_lookup_elem(&slab_hash, &s);
	if (name && !bpf_strncmp(name, 11, "task_struct"))
		task_struct_found = 1;
	else
		task_struct_found = -2;
	return 0;
}
```
