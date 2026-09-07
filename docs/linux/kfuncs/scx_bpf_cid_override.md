---
title: "KFunc 'scx_bpf_cid_override'"
description: "This page documents the 'scx_bpf_cid_override' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cid_override`

<!-- [FEATURE_TAG](scx_bpf_cid_override) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/df7b5ae038d6f2707ec0f81c257a55b4062f77c7)
<!-- [/FEATURE_TAG] -->

Install an explicit `cpu->cid` mapping.

## Definition

May only be called from [`ops.init`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#init) of the root scheduler. Replace the topology-probed cid mapping with the caller-provided one. Each possible cpu must map to a unique cid in [0, [`num_possible_cpus`](scx_bpf_nr_cids.md)). Topology info is cleared. On invalid input, trigger [`scx_error`](https://elixir.bootlin.com/linux/v7.2.2/source/kernel/sched/ext/internal.h#L1545) to abort the scheduler.

**Parameters**

`cpu_to_cid`: array of `nr_cpu_ids` `s32` entries (cid for each cpu)

`cpu_to_cid__sz`: must be `nr_cpu_ids * sizeof(s32)` bytes

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_cid_override(const s32 *cpu_to_cid, u32 cpu_to_cid__sz)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

