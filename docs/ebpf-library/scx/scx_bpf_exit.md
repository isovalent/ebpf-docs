---
title: "SCX eBPF macro 'scx_bpf_exit'"
description: "This page documents the 'scx_bpf_exit' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `scx_bpf_exit`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/4c30f5ce4f7af4f639af99e0bdeada8b268b7361)

The `scx_bpf_exit` macro wraps the [`scx_bpf_exit_bstr`](../../linux/kfuncs/scx_bpf_exit_bstr.md) kfunc with variadic arguments instead of an array of u64. Using this macro will cause the scheduler to exit cleanly with the specified exit code being passed to user space.

## Definition

```c
#define scx_bpf_exit(code, fmt, args...)                            \
({                                                                  \
	[scx_bpf_bstr_preamble](scx_bpf_bstr_preamble.md)(fmt, args)                                \
	[scx_bpf_exit_bstr](../../linux/kfuncs/scx_bpf_exit_bstr.md)(code, ___fmt, ___param, sizeof(___param));    \
	___scx_bpf_bstr_format_checker(fmt, ##args);                    \
})
```

## Usage

This macro can be used in an error path of the BPF scheduler that you expect to be hit under normal circumstances.

### Example

```c hl_lines="16 17"
/* SPDX-License-Identifier: GPL-2.0 */
/*
 * Create and destroy DSQs in a loop.
 *
 * Copyright (c) 2024 Meta Platforms, Inc. and affiliates.
 * Copyright (c) 2024 David Vernet <dvernet@meta.com>
 */

s32 [BPF_STRUCT_OPS_SLEEPABLE](BPF_STRUCT_OPS_SLEEPABLE.md)(create_dsq_init_task, struct task_struct *p,
			     struct scx_init_task_args *args)
{
	s32 err;

	err = [scx_bpf_create_dsq](../../linux/kfuncs/scx_bpf_create_dsq.md)(p->pid, -1);
	if (err)
		scx_bpf_error("Failed to create DSQ for %s[%d]",
			      p->comm, p->pid);

	return err;
}
```
