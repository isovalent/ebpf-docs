---
title: "SCX eBPF function 'scx_bpf_select_cpu_and'"
description: "This page documents the 'scx_bpf_select_cpu_and' scx eBPF function, including its definition, usage, and examples."
---
# SCX eBPF function `scx_bpf_select_cpu_and`

[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/c0d630ba347c7671210e1bab3c79defea19844e9)

The `scx_bpf_select_cpu_and` function wraps the [`scx_bpf_select_cpu_and`](../../linux/kfuncs/scx_bpf_select_cpu_and.md) and [`__scx_bpf_select_cpu_and`](../../linux/kfuncs/__scx_bpf_select_cpu_and.md) kfuncs. It implements the CO-RE logic needed to select the proper kfunc to use depending on the kernel being used.

## Definition

```c
static inline s32
scx_bpf_select_cpu_and(struct task_struct *p, s32 prev_cpu, u64 wake_flags,
		       const struct cpumask *cpus_allowed, u64 flags)
{
	if ([bpf_core_type_exists](../libbpf/ebpf/bpf_core_type_exists.md)(struct scx_bpf_select_cpu_and_args)) {
		struct scx_bpf_select_cpu_and_args args = {
			.prev_cpu = prev_cpu,
			.wake_flags = wake_flags,
			.flags = flags,
		};

		return [__scx_bpf_select_cpu_and](../../linux/kfuncs/__scx_bpf_select_cpu_and.md)(p, cpus_allowed, &args);
	} else {
		return [scx_bpf_select_cpu_and___compat](../../linux/kfuncs/scx_bpf_select_cpu_and.md)(p, prev_cpu, wake_flags,
						       cpus_allowed, flags);
	}
}
```

## Usage

See [`scx_bpf_select_cpu_and`](../../linux/kfuncs/scx_bpf_select_cpu_and.md)
