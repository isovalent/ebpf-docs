---
title: "SCX eBPF macro 'BPF_STRUCT_OPS'"
description: "This page documents the 'BPF_STRUCT_OPS' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `BPF_STRUCT_OPS`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `BPF_STRUCT_OPS` macro makes it easier to define [struct ops](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS.md) programs correctly.

## Definition

```c
#define BPF_STRUCT_OPS(name, args...)   \
    [SEC](../libbpf/ebpf/SEC.md)("struct_ops/"#name)             \
    [BPF_PROG](../libbpf/ebpf/BPF_PROG.md)(name, ##args)
```

## Usage

This macro can be used to shorted the definition of [struct ops](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS.md) programs. It places the program in an ELF section starting with `struct_ops/` and names the program with the given name, this signals the loader of the program type. It also unpacks the program context into the arguments specified via `args`. See [`BPF_PROG`](../libbpf/ebpf/BPF_PROG.md) for details.

### Example

```c hl_lines="7"
/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates.
 * Copyright (c) 2022 Tejun Heo <tj@kernel.org>
 * Copyright (c) 2022 David Vernet <dvernet@meta.com>
 */

void BPF_STRUCT_OPS(qmap_dump, struct scx_dump_ctx *dctx)
{
	s32 i, pid;

	if (suppress_dump)
		return;

	[bpf_for](../libbpf/ebpf/bpf_for.md)(i, 0, 5) {
		void *fifo;

		if (!(fifo = [bpf_map_lookup_elem](../../linux/helper-function/bpf_map_lookup_elem.md)(&queue_arr, &i)))
			return;

		[scx_bpf_dump](scx_bpf_dump.md)("QMAP FIFO[%d]:", i);
		[bpf_repeat](../libbpf/ebpf/bpf_repeat.md)(4096) {
			if ([bpf_map_pop_elem](../../linux/helper-function/bpf_map_pop_elem.md)(fifo, &pid))
				break;
			[scx_bpf_dump](scx_bpf_dump.md)(" %d", pid);
		}
		[scx_bpf_dump](scx_bpf_dump.md)("\n");
	}
}
```
