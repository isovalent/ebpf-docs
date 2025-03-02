---
title: "SCX eBPF macro 'scx_bpf_dsq_insert'"
description: "This page documents the 'scx_bpf_dsq_insert' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `scx_bpf_dsq_insert`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)

The `scx_bpf_dsq_insert` macro handles the renaming of [`scx_bpf_dispatch`](../../linux/kfuncs/scx_bpf_dispatch.md) to [`scx_bpf_dsq_insert`](../../linux/kfuncs/scx_bpf_dsq_insert.md) gracefully.

## Definition

```c
#define scx_bpf_dsq_insert(p, dsq_id, slice, enq_flags)				\
	([bpf_ksym_exists](../libbpf/ebpf/bpf_ksym_exists.md)(scx_bpf_dsq_insert) ?					\
	 [scx_bpf_dsq_insert](../../linux/kfuncs/scx_bpf_dsq_insert.md)((p), (dsq_id), (slice), (enq_flags)) :		\
	 [scx_bpf_dispatch___compat](../../linux/kfuncs/scx_bpf_dispatch.md)((p), (dsq_id), (slice), (enq_flags)))
```

## Usage

This macro has the same name as the [`scx_bpf_dsq_insert`](../../linux/kfuncs/scx_bpf_dsq_insert.md) kfunc, which will cause the pre-processor to emit this macro instead of just the kfunc. It checks at runtime if the kernel has the [`scx_bpf_dsq_insert`](../../linux/kfuncs/scx_bpf_dsq_insert.md) kfunc, and if it does, it calls it. If it doesn't, it calls the [`scx_bpf_dispatch___compat`](../../linux/kfuncs/scx_bpf_dispatch.md) kfunc instead.

These two kfuncs are functionally equivalent, but a rename happened since the name `dispatch` was overloaded and confusing.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
