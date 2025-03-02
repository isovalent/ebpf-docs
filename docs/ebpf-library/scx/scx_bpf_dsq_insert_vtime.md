---
title: "SCX eBPF macro 'scx_bpf_dsq_insert_vtime'"
description: "This page documents the 'scx_bpf_dsq_insert_vtime' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `scx_bpf_dsq_insert_vtime`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)

The `scx_bpf_dsq_insert_vtime` macro handles the renaming of [`scx_bpf_dispatch_vtime`](../../linux/kfuncs/scx_bpf_dispatch_vtime.md) to [`scx_bpf_dsq_insert_vtime`](../../linux/kfuncs/scx_bpf_dsq_insert_vtime.md) gracefully.

## Definition

```c
#define scx_bpf_dsq_insert_vtime(p, dsq_id, slice, vtime, enq_flags)		\
	([bpf_ksym_exists](../libbpf/ebpf/bpf_ksym_exists.md)(scx_bpf_dsq_insert_vtime) ?				\
	 [scx_bpf_dsq_insert_vtime](../../linux/kfuncs/scx_bpf_dsq_insert_vtime.md)((p), (dsq_id), (slice), (vtime), (enq_flags)) : \
	 [scx_bpf_dispatch_vtime___compat](../../linux/kfuncs/scx_bpf_dispatch_vtime.md)((p), (dsq_id), (slice), (vtime), (enq_flags)))

```

## Usage

This macro has the same name as the [`scx_bpf_dsq_insert_vtime`](../../linux/kfuncs/scx_bpf_dsq_insert_vtime.md) kfunc, which will cause the pre-processor to emit this macro instead of just the kfunc. It checks at runtime if the kernel has the [`scx_bpf_dsq_insert_vtime`](../../linux/kfuncs/scx_bpf_dsq_insert_vtime.md) kfunc, and if it does, it calls it. If it doesn't, it calls the [`scx_bpf_dispatch_vtime___compat`](../../linux/kfuncs/scx_bpf_dispatch_vtime.md) kfunc instead.

These two kfuncs are functionally equivalent, but a rename happened since the name `dispatch` was overloaded and confusing.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
