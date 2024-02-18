# KFunc `bpf_iter_css_new`

<!-- [FEATURE_TAG](bpf_iter_css_new) -->
[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/7251d0905e7518bcb990c8e9a3615b1bb23c78f2)
<!-- [/FEATURE_TAG] -->

Initialize a cgroup iterator.

## Definition

`it` should be a stack allocated `struct bpf_iter_css` that is used to iterate over cgroups. The `start` parameter is the cgroup subsystem state to start the iteration from. The `flags` parameter is a bitmask of flags that control the behavior of the iterator. The following flags are supported:

- `BPF_CGROUP_ITER_DESCENDANTS_PRE`: Walk descendants of the cgroup in pre-order.
- `BPF_CGROUP_ITER_DESCENDANTS_POST`: Walk descendants of the cgroup in post-order.
- `BPF_CGROUP_ITER_ANCESTORS_UP`: Walk ancestors of the cgroup upward.

<!-- [KFUNC_DEF] -->
`#!c int bpf_iter_css_new(struct bpf_iter_css *it, struct cgroup_subsys_state *start, unsigned int flags)`
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
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "bpf_misc.h"
#include "bpf_experimental.h"

char _license[] SEC("license") = "GPL";

pid_t target_pid;
u64 root_cg_id, leaf_cg_id;
u64 first_cg_id, last_cg_id;

int pre_order_cnt, post_order_cnt, tree_high;

struct cgroup *bpf_cgroup_from_id(u64 cgid) __ksym;
void bpf_cgroup_release(struct cgroup *p) __ksym;
void bpf_rcu_read_lock(void) __ksym;
void bpf_rcu_read_unlock(void) __ksym;

SEC("fentry.s/" SYS_PREFIX "sys_getpgid")
int iter_css_for_each(const void *ctx)
{
	struct task_struct *cur_task = bpf_get_current_task_btf();
	struct cgroup_subsys_state *root_css, *leaf_css, *pos;
	struct cgroup *root_cgrp, *leaf_cgrp, *cur_cgrp;

	if (cur_task->pid != target_pid)
		return 0;

	root_cgrp = bpf_cgroup_from_id(root_cg_id);

	if (!root_cgrp)
		return 0;

	leaf_cgrp = bpf_cgroup_from_id(leaf_cg_id);

	if (!leaf_cgrp) {
		bpf_cgroup_release(root_cgrp);
		return 0;
	}
	root_css = &root_cgrp->self;
	leaf_css = &leaf_cgrp->self;
	pre_order_cnt = post_order_cnt = tree_high = 0;
	first_cg_id = last_cg_id = 0;

	bpf_rcu_read_lock();
	bpf_for_each(css, pos, root_css, BPF_CGROUP_ITER_DESCENDANTS_POST) {
		cur_cgrp = pos->cgroup;
		post_order_cnt++;
		last_cg_id = cur_cgrp->kn->id;
	}

	bpf_for_each(css, pos, root_css, BPF_CGROUP_ITER_DESCENDANTS_PRE) {
		cur_cgrp = pos->cgroup;
		pre_order_cnt++;
		if (!first_cg_id)
			first_cg_id = cur_cgrp->kn->id;
	}

	bpf_for_each(css, pos, leaf_css, BPF_CGROUP_ITER_ANCESTORS_UP)
		tree_high++;

	bpf_for_each(css, pos, root_css, BPF_CGROUP_ITER_ANCESTORS_UP)
		tree_high--;
	bpf_rcu_read_unlock();
	bpf_cgroup_release(root_cgrp);
	bpf_cgroup_release(leaf_cgrp);
	return 0;
}
```
