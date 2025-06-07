---
title: "KFunc 'scx_bpf_pick_any_cpu'"
description: "This page documents the 'scx_bpf_pick_any_cpu' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_pick_any_cpu`

<!-- [FEATURE_TAG](scx_bpf_pick_any_cpu) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

Pick and claim an idle CPU if available or pick any CPU

## Definition

Pick and claim an idle CPU in `cpus_allowed`. If none is available, pick any CPU in `cpus_allowed`. Guaranteed to succeed.

If [`sched_ext_ops.update_idle`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#update_idle) is implemented and [`SCX_OPS_KEEP_BUILTIN_IDLE`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_ops_keep_builtin_idle) is not set, this function can't tell which CPUs are idle and will always pick any CPU.

Always returns an error if [`SCX_OPS_BUILTIN_IDLE_PER_NODE`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_ops_builtin_idle_per_node) is set, use [`scx_bpf_pick_idle_cpu_node`](scx_bpf_pick_idle_cpu_node.md) instead.

**Parameters**

`cpus_allowed`: Allowed cpumask

`flags`: `SCX_PICK_IDLE_*` flags

**Flags**

* `SCX_PICK_IDLE_CORE` - pick a CPU whose SMT siblings are also idle

**Returns**

The picked idle CPU number if `cpus_allowed` is not empty. `-EBUSY` is returned if `cpus_allowed` is empty.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_pick_any_cpu(const struct cpumask *cpus_allowed, u64 flags)`
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

