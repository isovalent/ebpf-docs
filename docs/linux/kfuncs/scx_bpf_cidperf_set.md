---
title: "KFunc 'scx_bpf_cidperf_set'"
description: "This page documents the 'scx_bpf_cidperf_set' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cidperf_set`

<!-- [FEATURE_TAG](scx_bpf_cidperf_set) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/5ba0a42423335f76d3e0513df42416c69dc6b742)
<!-- [/FEATURE_TAG] -->

Set the relative performance target of a CPU.

## Definition

Set the target performance level of `cpu` to `perf`. `perf` is in linear relative scale between 0 and [`SCX_CPUPERF_ONE`](https://elixir.bootlin.com/linux/v7.2.2/source/kernel/sched/ext/types.h#L29). This determines how the schedutil cpu frequency governor chooses the target frequency.

The actual performance level chosen, CPU grouping, and the overhead and latency of the operations are dependent on the hardware and cpu frequency driver in use. Consult hardware and cpu frequency documentation for more information. The current performance level can be monitored using [`scx_bpf_cpuperf_cur`](scx_bpf_cpuperf_cur.md).

**Parameters**

`cpu`: CPU of interest

`perf`: target performance level [0, [`SCX_CPUPERF_ONE`](https://elixir.bootlin.com/linux/v7.2.2/source/kernel/sched/ext/types.h#L29)]

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_cidperf_set(s32 cid, u32 perf)`
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

