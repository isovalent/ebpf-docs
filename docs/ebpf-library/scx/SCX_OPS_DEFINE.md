---
title: "SCX eBPF macro 'SCX_OPS_DEFINE'"
description: "This page documents the 'SCX_OPS_DEFINE' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `SCX_OPS_DEFINE`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `SCX_OPS_DEFINE` macro is used to define a full [`sched_ext_ops`](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md) implementation.

## Definition

```c
#define SCX_OPS_DEFINE(__name, ...)		\
	SEC(".struct_ops.link")			    \
	struct sched_ext_ops __name = {	    \
		__VA_ARGS__,					\
	};
```

## Usage

This macro can be used to define a full [`sched_ext_ops`](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md) implementation, associating the separate BPF programs with the corresponding [`sched_ext_ops`](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md) functions.

### Example

```c hl_lines="7"
/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates.
 * Copyright (c) 2022 Tejun Heo <tj@kernel.org>
 * Copyright (c) 2022 David Vernet <dvernet@meta.com>
 */

SCX_OPS_DEFINE(simple_ops,
	       .select_cpu		= (void *)simple_select_cpu,
	       .enqueue			= (void *)simple_enqueue,
	       .dispatch		= (void *)simple_dispatch,
	       .running			= (void *)simple_running,
	       .stopping		= (void *)simple_stopping,
	       .enable			= (void *)simple_enable,
	       .init			= (void *)simple_init,
	       .exit			= (void *)simple_exit,
	       .name			= "simple");
```
