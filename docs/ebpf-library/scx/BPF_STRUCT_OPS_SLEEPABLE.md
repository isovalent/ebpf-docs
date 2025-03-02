---
title: "SCX eBPF macro 'BPF_STRUCT_OPS'"
description: "This page documents the 'BPF_STRUCT_OPS' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `BPF_STRUCT_OPS`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `BPF_STRUCT_OPS` macro makes it easier to define [sleepable](../../linux/syscall/BPF_PROG_LOAD.md#bpf_f_sleepable) [struct ops](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS.md) programs correctly.

## Definition

```c
#define BPF_STRUCT_OPS_SLEEPABLE(name, args...)   \
    [SEC](../libbpf/ebpf/SEC.md)("struct_ops.s/"#name)             \
    [BPF_PROG](../libbpf/ebpf/BPF_PROG.md)(name, ##args)
```

## Usage

This macro can be used to shorted the definition of a [sleepable](../../linux/syscall/BPF_PROG_LOAD.md#bpf_f_sleepable) [struct ops](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS.md) programs. It places the program in an ELF section starting with `struct_ops.s/` and names the program with the given name, this signals the loader of the program type. It also unpacks the program context into the arguments specified via `args`. See [`BPF_PROG`](../libbpf/ebpf/BPF_PROG.md) for details.

### Example

```c hl_lines="24"
/* SPDX-License-Identifier: GPL-2.0 */
/*
 * Copyright (c) 2024 Meta Platforms, Inc. and affiliates.
 * Copyright (c) 2024 David Vernet <dvernet@meta.com>
 */

static void exit_from_hotplug(s32 cpu, bool onlining)
{
	/*
	 * Ignored, just used to verify that we can invoke blocking kfuncs
	 * from the hotplug path.
	 */
	[scx_bpf_create_dsq](../../linux/kfuncs/scx_bpf_create_dsq.md)(0, -1);

	s64 code = SCX_ECODE_ACT_RESTART | HOTPLUG_EXIT_RSN;

	if (onlining)
		code |= HOTPLUG_ONLINING;

	[scx_bpf_exit](scx_bpf_exit.md)(code, "hotplug event detected (%d going %s)", cpu,
		     onlining ? "online" : "offline");
}

void BPF_STRUCT_OPS_SLEEPABLE(hotplug_cpu_online, s32 cpu)
{
	exit_from_hotplug(cpu, true);
}
```
