---
title: "SCX eBPF macro 'scx_bpf_dsq_move_to_local'"
description: "This page documents the 'scx_bpf_dsq_move_to_local' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `scx_bpf_dsq_move_to_local`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)

The `scx_bpf_dsq_move_to_local` macro handles the renaming of [`scx_bpf_consume`](../../linux/kfuncs/scx_bpf_consume.md) to [`scx_bpf_dsq_move_to_local`](../../linux/kfuncs/scx_bpf_dsq_move_to_local.md) gracefully.

It also handles the introduction of [`scx_bpf_dsq_move_to_local___v2`](../../linux/kfuncs/scx_bpf_dsq_move_to_local___v2.md), using it when available, falling back to v1 when not.

## Definition

```c
#define scx_bpf_dsq_move_to_local(dsq_id, enq_flags)				\
	(bpf_ksym_exists(scx_bpf_dsq_move_to_local___v2) ?			\
	 scx_bpf_dsq_move_to_local___v2((dsq_id), (enq_flags)) :		\
	 (bpf_ksym_exists(scx_bpf_dsq_move_to_local___v1) ?			\
	  scx_bpf_dsq_move_to_local___v1((dsq_id)) :				\
	  scx_bpf_consume___old((dsq_id))))
```

## Usage

This macro has the same name as the [`scx_bpf_dsq_move_to_local`](../../linux/kfuncs/scx_bpf_dsq_move_to_local.md) kfunc, which will cause the pre-processor to emit this macro instead of just the kfunc. It checks at runtime if the kernel has the [`scx_bpf_dsq_move_to_local`](../../linux/kfuncs/scx_bpf_dsq_move_to_local.md) kfunc, and if it does, it calls it. If it doesn't, it calls the [`scx_bpf_consume___compat`](../../linux/kfuncs/scx_bpf_consume.md) kfunc instead.

These two kfuncs are functionally equivalent, but a rename happened since the name `dispatch` was overloaded and confusing.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
