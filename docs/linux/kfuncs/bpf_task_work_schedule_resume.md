---
title: "KFunc 'bpf_task_work_schedule_resume'"
description: "This page documents the 'bpf_task_work_schedule_resume' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_task_work_schedule_resume`

<!-- [FEATURE_TAG](bpf_task_work_schedule_resume) -->
[:octicons-tag-24: v7.0](https://github.com/torvalds/linux/commit/6e663ffdf7600168338fdfa2fd1eed83395d58a3)
<!-- [/FEATURE_TAG] -->

Schedule BPF callback using `task_work_add` with `TWA_RESUME` mode.

!!! note
    This kfunc supersedes `bpf_task_work_schedule_resume_impl`, migrated to use implicit arguments in [:octicons-tag-24: v7.0](https://github.com/torvalds/linux/commit/6e663ffdf7600168338fdfa2fd1eed83395d58a3).

## Definition

**Parameters**

`task`: Task struct for which callback should be scheduled

`tw`: Pointer to `struct bpf_task_work` in BPF map value for internal bookkeeping

`map__map`: map that embeds `struct bpf_task_work` in the values

`callback`: pointer to BPF subprogram to call


**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_task_work_schedule_resume(struct task_struct *task, struct bpf_task_work *tw, void *map__map, bpf_task_work_callback_t callback)`
<!-- [/KFUNC_DEF] -->

`#!c typedef int (*bpf_task_work_callback_t)(struct bpf_map *map, void *key, void *value);`

## Usage

This kfunc allows a BPF program that is being executed in a restricted context such as a Non Mask-able Interrupt (NMI) to schedule a callback on a task. This callback will be executed in a more permissible context (sleepable context) before that task resumes.

This is mostly useful for tools such as profilers. When a program is triggered in an NMI, the program and any helper/kfunc it executes is unable to sleep/wait or page fault. This means that some actions like reading userspace memory or even updating map values may fail. So by scheduling a callback you can do more things in the permissive context, while still passing info from the original execution context via a map value.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
- [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
