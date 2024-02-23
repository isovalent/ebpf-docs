---
title: "KFunc 'bpf_iter_task_vma_new' - eBPF Docs"
description: "This page documents the 'bpf_iter_task_vma_new' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_task_vma_new`

<!-- [FEATURE_TAG](bpf_iter_task_vma_new) -->
[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/4ac4546821584736798aaa9e97da9f6eaf689ea3)
<!-- [/FEATURE_TAG] -->

Create a new task vma iterator.

## Definition

<!-- [KFUNC_DEF] -->
`#!c int bpf_iter_task_vma_new(struct bpf_iter_task_vma *it, struct task_struct *task, u64 addr)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [BPF_PROG_TYPE_NETFILTER](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2023 Meta Platforms, Inc. and affiliates. */

#include "vmlinux.h"
#include "bpf_experimental.h"
#include <bpf/bpf_helpers.h>
#include "bpf_misc.h"

pid_t target_pid = 0;
unsigned int vmas_seen = 0;

struct {
	__u64 vm_start;
	__u64 vm_end;
} vm_ranges[1000];

SEC("raw_tp/sys_enter")
int iter_task_vma_for_each(const void *ctx)
{
	struct task_struct *task = bpf_get_current_task_btf();
	struct vm_area_struct *vma;
	unsigned int seen = 0;

	if (task->pid != target_pid)
		return 0;

	if (vmas_seen)
		return 0;

	bpf_for_each(task_vma, vma, task, 0) {
		if (seen >= 1000)
			break;

		vm_ranges[seen].vm_start = vma->vm_start;
		vm_ranges[seen].vm_end = vma->vm_end;
		seen++;
	}

	vmas_seen = seen;
	return 0;
}

char _license[] SEC("license") = "GPL";
```
