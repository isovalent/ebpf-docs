---
title: "KFunc 'scx_bpf_pick_idle_cpu'"
description: "This page documents the 'scx_bpf_pick_idle_cpu' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_pick_idle_cpu`

<!-- [FEATURE_TAG](scx_bpf_pick_idle_cpu) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

This function picks and claims an idle cpu.

## Definition

Pick and claim an idle cpu in `cpus_allowed`. 

Idle CPU tracking may race against CPU scheduling state transitions. For example, this function may return `-EBUSY` as CPUs are transitioning into the idle state. If the caller then assumes that there will be dispatch events on the CPUs as they were all busy, the scheduler may end up stalling with CPUs idling while there are pending tasks. Use [`scx_bpf_pick_any_cpu`](scx_bpf_pick_any_cpu.md) and [`scx_bpf_kick_cpu`](scx_bpf_kick_cpu.md) to guarantee that there will be at least one dispatch event in the near future.

Unavailable if `ops.update_idle()` is implemented and `SCX_OPS_KEEP_BUILTIN_IDLE` is not set.

**Parameters**

`cpus_allowed`: Allowed cpumask

`flags`: `SCX_PICK_IDLE_*` flags

**Flags**

* `SCX_PICK_IDLE_CORE` - pick a CPU whose SMT siblings are also idle

**Returns**

The picked idle cpu number on success. `-EBUSY` if no matching cpu was found.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_pick_idle_cpu(const struct cpumask *cpus_allowed, u64 flags)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

