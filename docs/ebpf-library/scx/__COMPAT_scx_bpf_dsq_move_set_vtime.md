---
title: "SCX eBPF macro '__COMPAT_scx_bpf_dsq_move_set_vtime'"
description: "This page documents the '__COMPAT_scx_bpf_dsq_move_set_vtime' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `__COMPAT_scx_bpf_dsq_move_set_vtime`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)

The `__COMPAT_scx_bpf_dsq_move_set_vtime` macro handles both the renaming of [`scx_bpf_dispatch_from_dsq_set_vtime`](../../linux/kfuncs/scx_bpf_dispatch_from_dsq_set_vtime.md) to 
[`scx_bpf_dsq_move_set_vtime`](../../linux/kfuncs/scx_bpf_dsq_move_set_vtime.md), and the case where the kernel has neither of these kfuncs.
## Definition

```c
#define __COMPAT_scx_bpf_dsq_move_set_vtime(it__iter, vtime)			\
	([bpf_ksym_exists](../libbpf/ebpf/bpf_ksym_exists.md)(scx_bpf_dsq_move_set_vtime) ?				\
	 [scx_bpf_dsq_move_set_vtime](../../linux/kfuncs/scx_bpf_dsq_move_set_vtime.md)((it__iter), (vtime)) :			\
	 ([bpf_ksym_exists](../libbpf/ebpf/bpf_ksym_exists.md)(scx_bpf_dispatch_from_dsq_set_vtime___compat) ?	\
	  [scx_bpf_dispatch_from_dsq_set_vtime___compat](../../linux/kfuncs/scx_bpf_dispatch_from_dsq_set_vtime.md)((it__iter), (vtime)) :	\
	  (void)0))

```

## Usage

This macro checks at runtime if the kernel has the `scx_bpf_dsq_move_set_vtime`(../../linux/kfuncs/scx_bpf_dsq_move_set_vtime.md) kfunc, and if it does, it calls it. If it doesn't, it checks if the kernel has the [`scx_bpf_dispatch_from_dsq_set_vtime`](../../linux/kfuncs/scx_bpf_dispatch_from_dsq_set_vtime.md) kfunc, and if it does, it calls it. If neither of these kfuncs are available, it return zero.

These two kfuncs are functionally equivalent, but a rename happened since the name `dispatch` was overloaded and confusing.

All of this logic makes sure your program can run on any kernel, no matter changes made to the kfuncs.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
