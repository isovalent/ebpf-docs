---
title: "SCX eBPF macro 'bpf_obj_drop'"
description: "This page documents the 'bpf_obj_drop' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `bpf_obj_drop`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `bpf_obj_drop` macro wraps [`bpf_obj_drop_impl`](../../linux/kfuncs/bpf_obj_drop_impl.md) to provide a more ergonomic interface.

## Definition

```c
#define bpf_obj_drop(kptr) [bpf_obj_drop_impl](../../linux/kfuncs/bpf_obj_drop_impl.md)(kptr, NULL)
```

## Usage

The [`bpf_obj_drop_impl`](../../linux/kfuncs/bpf_obj_drop_impl.md) kfunc has a quirk where the second argument is always `NULL`, this wrapper abstracts that quirk away.

### Example

```c hl_lines="67"
/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (c) 2022 Tejun Heo <tj@kernel.org> */

int [BPF_STRUCT_OPS_SLEEPABLE](BPF_STRUCT_OPS_SLEEPABLE.md)(fcg_cgroup_init, struct cgroup *cgrp,
			     struct scx_cgroup_init_args *args)
{
	struct fcg_cgrp_ctx *cgc;
	struct cgv_node *cgv_node;
	struct cgv_node_stash empty_stash = {}, *stash;
	u64 cgid = cgrp->kn->id;
	int ret;

	/*
	 * Technically incorrect as cgroup ID is full 64bit while dsq ID is
	 * 63bit. Should not be a problem in practice and easy to spot in the
	 * unlikely case that it breaks.
	 */
	ret = [scx_bpf_create_dsq](../../linux/kfuncs/scx_bpf_create_dsq.md)(cgid, -1);
	if (ret)
		return ret;

	cgc = [bpf_cgrp_storage_get](../../linux/helper-function/bpf_cgrp_storage_get.md)(&cgrp_ctx, cgrp, 0,
				   BPF_LOCAL_STORAGE_GET_F_CREATE);
	if (!cgc) {
		ret = -ENOMEM;
		goto err_destroy_dsq;
	}

	cgc->weight = args->weight;
	cgc->hweight = FCG_HWEIGHT_ONE;

	ret = [bpf_map_update_elem](../../linux/helper-function/bpf_map_update_elem.md)(&cgv_node_stash, &cgid, &empty_stash,
				  BPF_NOEXIST);
	if (ret) {
		if (ret != -ENOMEM)
			[scx_bpf_error](scx_bpf_error.md)("unexpected stash creation error (%d)",
				      ret);
		goto err_destroy_dsq;
	}

	stash = [bpf_map_lookup_elem](../../linux/helper-function/bpf_map_lookup_elem.md)(&cgv_node_stash, &cgid);
	if (!stash) {
		[scx_bpf_error](scx_bpf_error.md)("unexpected cgv_node stash lookup failure");
		ret = -ENOENT;
		goto err_destroy_dsq;
	}

	cgv_node = [bpf_obj_new](bpf_obj_new.md)(struct cgv_node);
	if (!cgv_node) {
		ret = -ENOMEM;
		goto err_del_cgv_node;
	}

	cgv_node->cgid = cgid;
	cgv_node->cvtime = cvtime_now;

	cgv_node = [bpf_kptr_xchg](../../linux/helper-function/bpf_kptr_xchg.md)(&stash->node, cgv_node);
	if (cgv_node) {
		[scx_bpf_error](scx_bpf_error.md)("unexpected !NULL cgv_node stash");
		ret = -EBUSY;
		goto err_drop;
	}

	return 0;

err_drop:
	bpf_obj_drop(cgv_node);
err_del_cgv_node:
	[bpf_map_delete_elem](../../linux/helper-function/bpf_map_delete_elem.md)(&cgv_node_stash, &cgid);
err_destroy_dsq:
	[scx_bpf_destroy_dsq](../../linux/kfuncs/scx_bpf_destroy_dsq.md)(cgid);
	return ret;
}
```
