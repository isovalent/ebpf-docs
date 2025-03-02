---
title: "SCX eBPF macro 'scx_bpf_error'"
description: "This page documents the 'scx_bpf_error' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `scx_bpf_error`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/4c30f5ce4f7af4f639af99e0bdeada8b268b7361)

The `scx_bpf_error` macro wraps the [`scx_bpf_error_bstr`](../../linux/kfuncs/scx_bpf_error_bstr.md) kfunc with variadic arguments instead of an array of u64. Invoking this macro will cause the scheduler to exit in an erroneous state, with diagnostic information being passed to the user.

## Definition

```c
#define scx_bpf_error(fmt, args...)						\
({										\
	[scx_bpf_bstr_preamble](scx_bpf_bstr_preamble.md)(fmt, args)					\
	[scx_bpf_error_bstr](../../linux/kfuncs/scx_bpf_error_bstr.md)(___fmt, ___param, sizeof(___param));			\
	___scx_bpf_bstr_format_checker(fmt, ##args);				\
})
```

## Usage

This macro can be used in an error path of the BPF scheduler that should not be hit under normal circumstances.

### Example

Loop over all tasks in the shared DSQ and move them to bpf scheduler defined `SOME_DSQ`.

```c hl_lines="20 21"
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

	scx_bpf_exit(code, "hotplug event detected (%d going %s)", cpu,
		     onlining ? "online" : "offline");
}

void [BPF_STRUCT_OPS_SLEEPABLE](BPF_STRUCT_OPS_SLEEPABLE.md)(hotplug_cpu_online, s32 cpu)
{
	exit_from_hotplug(cpu, true);
}
```
