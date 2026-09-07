---
title: "KFunc 'scx_bpf_nr_online_cids'"
description: "This page documents the 'scx_bpf_nr_online_cids' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_nr_online_cids`

<!-- [FEATURE_TAG](scx_bpf_nr_online_cids) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/5ba0a42423335f76d3e0513df42416c69dc6b742)
<!-- [/FEATURE_TAG] -->

Return current count of online CPUs in cid space.

## Definition

**Returns**

Return [`num_online_cpus`](https://elixir.bootlin.com/linux/v7.2.2/source/include/linux/cpumask.h#L1263). The standard model restarts the scheduler on hot-plug, which lets schedulers treat [0, `nr_online_cids`) as the online range. Schedulers that prefer to handle hotplug without a restart should install a custom mapping via [`scx_bpf_cid_override`](scx_bpf_cid_override.md) and track onlining through the ops.cid_online / `ops.cid_offline` callbacks.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 scx_bpf_nr_online_cids()`
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

