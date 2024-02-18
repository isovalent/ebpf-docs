# KFunc `bpf_iter_css_task_new`

<!-- [FEATURE_TAG](bpf_iter_css_task_new) -->
[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/9c66dc94b62aef23300f05f63404afb8990920b4)
<!-- [/FEATURE_TAG] -->

Initialize a task iterator for a cgroup.

## Definition

`it` should be a stack allocated `struct bpf_iter_css_task` that is used to iterate over tasks in a cgroup. The `css` parameter is the cgroup subsystem state to iterate over. The `flags` parameter is a bitmask of flags that control the behavior of the iterator. The following flags are supported:

- `0`: Walk all tasks in the domain.
- `CSS_TASK_ITER_PROCS`: Walk only threadgroup leaders.
- `CSS_TASK_ITER_PROCS | CSS_TASK_ITER_THREADED`: Walk all threaded css_sets in the domain.

<!-- [KFUNC_DEF] -->
`#!c int bpf_iter_css_task_new(struct bpf_iter_css_task *it, struct cgroup_subsys_state *css, unsigned int flags)`
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
/* Copyright (C) 2023 Chuyi Zhou <zhouchuyi@bytedance.com> */

#include "vmlinux.h"
#include <errno.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "bpf_misc.h"
#include "bpf_experimental.h"

char _license[] SEC("license") = "GPL";

struct cgroup *bpf_cgroup_from_id(u64 cgid) __ksym;
void bpf_cgroup_release(struct cgroup *p) __ksym;

pid_t target_pid;
int css_task_cnt;
u64 cg_id;

SEC("lsm/file_mprotect")
int BPF_PROG(iter_css_task_for_each, struct vm_area_struct *vma,
	    unsigned long reqprot, unsigned long prot, int ret)
{
	struct task_struct *cur_task = bpf_get_current_task_btf();
	struct cgroup_subsys_state *css;
	struct task_struct *task;
	struct cgroup *cgrp;

	if (cur_task->pid != target_pid)
		return ret;

	cgrp = bpf_cgroup_from_id(cg_id);

	if (!cgrp)
		return -EPERM;

	css = &cgrp->self;
	css_task_cnt = 0;

	bpf_for_each(css_task, task, css, CSS_TASK_ITER_PROCS)
		if (task->pid == target_pid)
			css_task_cnt++;

	bpf_cgroup_release(cgrp);

	return -EPERM;
}
```
