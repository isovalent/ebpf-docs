---
title: "SCX eBPF function 'scx_bpf_dsq_insert_vtime'"
description: "This page documents the 'scx_bpf_dsq_insert_vtime' scx eBPF function, including its definition, usage, and examples."
---
# SCX eBPF function `scx_bpf_dsq_insert_vtime`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429)

The `scx_bpf_dsq_insert_vtime` function handles the renaming of [`scx_bpf_dispatch_vtime`](../../linux/kfuncs/scx_bpf_dispatch_vtime.md) to [`scx_bpf_dsq_insert_vtime`](../../linux/kfuncs/scx_bpf_dsq_insert_vtime.md) and the later transition to [`__scx_bpf_dsq_insert_vtime`](../../linux/kfuncs/__scx_bpf_dsq_insert_vtime.md) gracefully.

## Definition

```c
static inline bool
scx_bpf_dsq_insert_vtime(struct task_struct *p, u64 dsq_id, u64 slice, u64 vtime,
			 u64 enq_flags)
{
	if ([bpf_core_type_exists](../libbpf/ebpf/bpf_core_type_exists.md)(struct scx_bpf_dsq_insert_vtime_args)) {
		struct scx_bpf_dsq_insert_vtime_args args = {
			.dsq_id = dsq_id,
			.slice = slice,
			.vtime = vtime,
			.enq_flags = enq_flags,
		};

		return [__scx_bpf_dsq_insert_vtime](../../linux/kfuncs/__scx_bpf_dsq_insert_vtime.md)(p, &args);
	} else if (bpf_ksym_exists([scx_bpf_dsq_insert_vtime___compat](../../linux/kfuncs/scx_bpf_dsq_insert_vtime.md)))) {
		[scx_bpf_dsq_insert_vtime___compat](../../linux/kfuncs/scx_bpf_dsq_insert_vtime.md)(p, dsq_id, slice, vtime,
						  enq_flags);
		return true;
	} else {
		[scx_bpf_dispatch_vtime___compat](../../linux/kfuncs/scx_bpf_dispatch_vtime.md)(p, dsq_id, slice, vtime,
						enq_flags);
		return true;
	}
}
```

## Usage

This function has the same name as the [`scx_bpf_dsq_insert_vtime`](../../linux/kfuncs/scx_bpf_dsq_insert_vtime.md) kfunc. The kfunc was named [`scx_bpf_dispatch_vtime`](../../linux/kfuncs/scx_bpf_dispatch_vtime.md) in older kernels. And in more recent kernels the [`__scx_bpf_dsq_insert_vtime`](../../linux/kfuncs/__scx_bpf_dsq_insert_vtime.md) kfunc was introduced, which receives the arguments as a structure, allowing for more then 5 function arguments (a BPF instruction set limitation). This compatibility function used CO-RE to pick the correct kfunc to use depending on the version of the kernel the program is being loaded on.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
