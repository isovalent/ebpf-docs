---
title: "SCX eBPF macro 'RESIZABLE_ARRAY'"
description: "This page documents the 'RESIZABLE_ARRAY' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `RESIZABLE_ARRAY`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `RESIZABLE_ARRAY` macro generates annotations for an array that may be resized.

## Definition

```c
#define RESIZABLE_ARRAY(elfsec, arr) [SEC](../libbpf/ebpf/SEC.md)("."#elfsec"."#arr) arr[1]
```

## Usage

libbpf has an API for setting map value sizes. Since data sections (i.e. `bss`, `data`, `rodata`) themselves are maps, a data section can be resized. If a data section has an array as its last element, the BTF info for that array will be adjusted so that length of the array is extended to meet the new length of the data section. This macro annotates an array to have an element count of one with the assumption that this array can be resized within the userspace program. It also annotates the section specifier so this array exists in a custom sub data section which can be resized independently.

The [`ARRAY_ELEM_PTR`](ARRAY_ELEM_PTR.md) macro can be used to get a pointer (or `NULL`) into the array without triping the verifier.

**Parameters**

`elfsec`: the data section of the BPF program in which to place the array

`arr`: the name of the array

### Example

```c hl_lines="7"
/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates.
 * Copyright (c) 2022 Tejun Heo <tj@kernel.org>
 * Copyright (c) 2022 David Vernet <dvernet@meta.com>
 */

u64 RESIZABLE_ARRAY(data, cpu_started_at);

void [BPF_STRUCT_OPS](BPF_STRUCT_OPS.md)(central_running, struct task_struct *p)
{
	s32 cpu = [scx_bpf_task_cpu](../../linux/kfuncs/scx_bpf_task_cpu.md)(p);
	u64 *started_at = [ARRAY_ELEM_PTR](ARRAY_ELEM_PTR.md)(cpu_started_at, cpu, nr_cpu_ids);
	if (started_at)
		*started_at = [scx_bpf_now](../../linux/kfuncs/scx_bpf_now.md)() ?: 1;	/* 0 indicates idle */
}
```
