---
title: "SCX eBPF macro 'scx_bpf_reenqueue_local'"
description: "This page documents the 'scx_bpf_reenqueue_local' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `scx_bpf_reenqueue_local`

[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/a3f5d48222532484c1e85ef27cc6893803e4cd17)

The `scx_bpf_reenqueue_local` macro wraps the [`scx_bpf_reenqueue_local`](../../linux/kfuncs/scx_bpf_reenqueue_local.md) and [`scx_bpf_reenqueue_local___v2`](../../linux/kfuncs/scx_bpf_reenqueue_local___v2.md) kfuncs. It implements the CO-RE logic needed to select the proper kfunc to use depending on the kernel being used.

## Definition

```c
static inline void scx_bpf_reenqueue_local(void)
{
	if (__COMPAT_scx_bpf_reenqueue_local_from_anywhere())
		scx_bpf_reenqueue_local___v2___compat();
	else
		scx_bpf_reenqueue_local___v1();
}
```

## Usage

See [`scx_bpf_reenqueue_local`](../../linux/kfuncs/scx_bpf_reenqueue_local.md)
