---
title: "KFunc '__scx_bpf_select_cpu_and'"
description: "This page documents the '__scx_bpf_select_cpu_and' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `__scx_bpf_select_cpu_and`

<!-- [FEATURE_TAG](__scx_bpf_select_cpu_and) -->
[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/c0d630ba347c7671210e1bab3c79defea19844e9)
<!-- [/FEATURE_TAG] -->

Argument-wrapped CPU selection with cpumask.

## Definition

Wrapper kfunc that takes arguments via struct to work around the 5 argument limit for BPF functions. BPF programs should use [`scx_bpf_select_cpu_and`](../../ebpf-library/scx/scx_bpf_select_cpu_and.md) which is provided as an inline wrapper in common.bpf.h.

Can be called from [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu), [`sched_ext_ops.enqueue`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#enqueue), or from an unlocked context such as a BPF [`test_run`](../syscall/BPF_PROG_TEST_RUN.md) call, as long as built-in CPU selection is enabled: [`sched_ext_ops.update_idle`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#update_idle) is missing or [`SCX_OPS_KEEP_BUILTIN_IDLE`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_ops_keep_builtin_idle) is set.

`p`, `args->prev_cpu` and `args->wake_flag` match [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu).

Returns the selected idle CPU, which will be automatically awakened upon returning from [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu) and can be used for direct dispatch, or a negative value if no idle CPU is available.

**Parameters**

`p`: task_struct to select a CPU for

`cpus_allowed`: cpumask of allowed CPUs

`args`: Structure of additional arguments

```c
struct scx_bpf_select_cpu_and_args {
	/* @p and @cpus_allowed can't be packed together as KF_RCU is not transitive */
	s32			prev_cpu;
	u64			wake_flags;
	u64			flags;
};
```

`args->prev_cpu`: CPU `p` was on previously

`args->wake_flags`: `SCX_WAKE_*`, possible values are:

* `SCX_WAKE_FORK` (`0x02`) - Wakeup after exec
* `SCX_WAKE_TTWU` (`0x04`) - Wakeup after fork
* `SCX_WAKE_SYNC` (`0x08`) - Wakeup

`args->cpus_allowed`: cpumask of allowed CPUs

`args->flags`: `SCX_PICK_IDLE_CPU_*` flags

**Flags**

```c
enum scx_pick_idle_cpu_flags {
	SCX_PICK_IDLE_CORE      = 1LLU << 0,
	SCX_PICK_IDLE_IN_NODE   = 1LLU << 1,
};
```

`SCX_PICK_IDLE_CORE`: pick a CPU whose SMT siblings are also idle

`SCX_PICK_IDLE_IN_NODE`: pick a CPU in the same target NUMA node 


**Returns**

The selected idle CPU, which will be automatically awakened upon returning from [`sched_ext_ops.select_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#select_cpu) and can be used for direct dispatch, or a negative value if no idle CPU is available.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 __scx_bpf_select_cpu_and(struct task_struct *p, const struct cpumask *cpus_allowed, struct scx_bpf_select_cpu_and_args *args)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

