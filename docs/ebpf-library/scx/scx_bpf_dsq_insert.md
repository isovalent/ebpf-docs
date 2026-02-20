---
title: "SCX eBPF function 'scx_bpf_dsq_insert'"
description: "This page documents the 'scx_bpf_dsq_insert' scx eBPF function, including its definition, usage, and examples."
---
# SCX eBPF function `scx_bpf_dsq_insert`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)

The `scx_bpf_dsq_insert` function handles the renaming of [`scx_bpf_dispatch`](../../linux/kfuncs/scx_bpf_dispatch.md) to [`scx_bpf_dsq_insert`](../../linux/kfuncs/scx_bpf_dsq_insert.md) and the migration to [`scx_bpf_dsq_insert___v2`](../../linux/kfuncs/scx_bpf_dsq_insert___v2.md) gracefully.

## Definition

!!! note
    In [:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/cded46d971597ecfe505ba92a54253c0f5e1f2e4) the signature of this function got changed to return a `bool`.

```c
static inline bool
scx_bpf_dsq_insert(struct task_struct *p, u64 dsq_id, u64 slice, u64 enq_flags)
{
	if ([bpf_ksym_exists](../libbpf/ebpf/bpf_ksym_exists.md)(scx_bpf_dsq_insert___v2___compat)) {
		return [scx_bpf_dsq_insert___v2___compat](../../linux/kfuncs/scx_bpf_dsq_insert___v2.md)(p, dsq_id, slice, enq_flags);
	} else if ([bpf_ksym_exists](../libbpf/ebpf/bpf_ksym_exists.md)(scx_bpf_dsq_insert___v1)) {
		[scx_bpf_dsq_insert___v1](../../linux/kfuncs/scx_bpf_dsq_insert.md)(p, dsq_id, slice, enq_flags);
		return true;
	} else {
		[scx_bpf_dispatch___compat](../../linux/kfuncs/scx_bpf_dispatch.md)(p, dsq_id, slice, enq_flags);
		return true;
	}
}
```

## Usage

This function has the same name as the [`scx_bpf_dsq_insert`](../../linux/kfuncs/scx_bpf_dsq_insert.md) kfunc, which will cause the pre-processor to emit this function instead of just the kfunc. It checks at runtime if the kernel has the [`scx_bpf_dsq_insert`](../../linux/kfuncs/scx_bpf_dsq_insert.md) kfunc, and if it does, it calls it. If it doesn't, it calls the [`scx_bpf_dispatch___compat`](../../linux/kfuncs/scx_bpf_dispatch.md) kfunc instead.

These two kfuncs are functionally equivalent, but a rename happened since the name `dispatch` was overloaded and confusing.

Since [:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/cded46d971597ecfe505ba92a54253c0f5e1f2e4), the `scx_bpf_dsq_insert` kfunc is deprecated in favor of the V2 version which returns a boolean. This function was modified to return a boolean even when older kfuncs are in use and to call the v2 kfunc when available.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
