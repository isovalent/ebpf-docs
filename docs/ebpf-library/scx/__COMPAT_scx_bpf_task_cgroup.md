---
title: "SCX eBPF macro '__COMPAT_scx_bpf_task_cgroup'"
description: "This page documents the '__COMPAT_scx_bpf_task_cgroup' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `__COMPAT_scx_bpf_task_cgroup`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/1e123fd73deb16cb362ecefb55c90c9196f4a6c2)

The `__COMPAT_scx_bpf_task_cgroup` macro executes [`scx_bpf_task_cgroup`](../../linux/kfuncs/scx_bpf_task_cgroup.md) or returns NULL if the function is not available.

## Definition

```c
#define __COMPAT_scx_bpf_task_cgroup(p)						\
	([bpf_ksym_exists](../libbpf/ebpf/bpf_ksym_exists.md)(scx_bpf_task_cgroup) ?					\
	 [scx_bpf_task_cgroup](../../linux/kfuncs/scx_bpf_task_cgroup.md)((p)) : NULL)
```

## Usage

This macro handles checking for the existence of [`scx_bpf_task_cgroup`](../../linux/kfuncs/scx_bpf_task_cgroup.md) at runtime, and calling it if it exists. If you were to have a program that called [`scx_bpf_task_cgroup`](../../linux/kfuncs/scx_bpf_task_cgroup.md) directly, without this check, it would refuse to load on kernels before the kfunc was added.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
