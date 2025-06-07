---
title: "KFunc 'scx_bpf_pick_idle_cpu_node'"
description: "This page documents the 'scx_bpf_pick_idle_cpu_node' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_pick_idle_cpu_node`

<!-- [FEATURE_TAG](scx_bpf_pick_idle_cpu_node) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/01059219b0cfdb9fc0d5bd60458e614a3135e6e7)
<!-- [/FEATURE_TAG] -->

Pick and claim an idle cpu from `node`

## Definition

Pick and claim an idle cpu in `cpus_allowed` from the NUMA node @node.

The search starts from @node and proceeds to other online NUMA nodes in order of increasing distance (unless `SCX_PICK_IDLE_IN_NODE` is specified, in which case the search is limited to the target `node`).

Always returns an error if [`ops.update_idle`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#update_idle) is implemented and [`SCX_OPS_KEEP_BUILTIN_IDLE`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_ops_keep_builtin_idle) is not set, or if [`SCX_OPS_BUILTIN_IDLE_PER_NODE`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#scx_ops_builtin_idle_per_node) is not set.

**Parameters**

`cpus_allowed`: Allowed cpumask

`node`: target NUMA node

`flags`: `SCX_PICK_IDLE_*` flags

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

The picked idle cpu number on success, or `-EBUSY` if no matching cpu was found.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_pick_idle_cpu_node(const struct cpumask *cpus_allowed, int node, u64 flags)`
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

