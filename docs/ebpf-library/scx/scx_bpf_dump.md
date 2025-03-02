---
title: "SCX eBPF macro 'scx_bpf_exit'"
description: "This page documents the 'scx_bpf_exit' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `scx_bpf_exit`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/4c30f5ce4f7af4f639af99e0bdeada8b268b7361)

The `scx_bpf_exit` macro wraps the [`scx_bpf_dump_bstr`](../../linux/kfuncs/scx_bpf_dump_bstr.md) kfunc with variadic arguments instead of an array of u64. To be used from [`sched_ext_ops.dump`](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dump) and friends.

## Definition

```c
#define scx_bpf_dump(fmt, args...)						\
({										\
	[scx_bpf_bstr_preamble](scx_bpf_bstr_preamble.md)(fmt, args)					\
	[scx_bpf_dump_bstr](../../linux/kfuncs/scx_bpf_dump_bstr.md)(___fmt, ___param, sizeof(___param));			\
	___scx_bpf_bstr_format_checker(fmt, ##args);				\
})

```

## Usage

This macro can be used to help dump information from the BPF scheduler.

### Example

```c hl_lines="20 24 26"
/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates.
 * Copyright (c) 2022 Tejun Heo <tj@kernel.org>
 * Copyright (c) 2022 David Vernet <dvernet@meta.com>
 */

void [BPF_STRUCT_OPS](BPF_STRUCT_OPS.md)(qmap_dump, struct scx_dump_ctx *dctx)
{
	s32 i, pid;

	if (suppress_dump)
		return;

	[bpf_for](../libbpf/ebpf/bpf_for.md)(i, 0, 5) {
		void *fifo;

		if (!(fifo = [bpf_map_lookup_elem](../../linux/helper-function/bpf_map_lookup_elem.md)(&queue_arr, &i)))
			return;

		scx_bpf_dump("QMAP FIFO[%d]:", i);
		[bpf_repeat](../libbpf/ebpf/bpf_repeat.md)(4096) {
			if ([bpf_map_pop_elem](../../linux/helper-function/bpf_map_pop_elem.md)(fifo, &pid))
				break;
			scx_bpf_dump(" %d", pid);
		}
		scx_bpf_dump("\n");
	}
}
```
