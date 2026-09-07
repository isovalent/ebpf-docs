---
title: "KFunc 'scx_bpf_kick_cid'"
description: "This page documents the 'scx_bpf_kick_cid' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_kick_cid`

<!-- [FEATURE_TAG](scx_bpf_kick_cid) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/5ba0a42423335f76d3e0513df42416c69dc6b742)
<!-- [/FEATURE_TAG] -->

Trigger reschedule on the CPU mapped to `cid`.

## Definition

cid-addressed equivalent of [`scx_bpf_kick_cpu`](scx_bpf_kick_cpu.md).

**Parameters**

`cid`: cid to kick

`flags`: `SCX_KICK_*` flags

**Flags**

`SCX_KICK_IDLE`: Kick the target CPU if idle. Guarantees that the target CPU goes through at least one full scheduling cycle before going idle. If the target CPU can be determined to be currently not idle and going to go through a scheduling cycle before going idle, noop.

<code id="scx_kick_preempt">SCX_KICK_PREEMPT</code>: Preempt the current task and execute the dispatch path. If the current task of the target CPU is an SCX task, its [`->scx.slice`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#struct-sched_ext_entity-slice) is cleared to zero before the scheduling path is invoked so that the task expires and the dispatch path is invoked.

`SCX_KICK_WAIT`: Wait for the CPU to be rescheduled. The `scx_bpf_kick_cpu` call will return after the target CPU finishes picking the next task.

**Returns**

`0` on success, `-errno` otherwise.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_kick_cid(s32 cid, u64 flags)`
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

