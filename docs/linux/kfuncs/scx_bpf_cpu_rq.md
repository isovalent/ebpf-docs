---
title: "KFunc 'scx_bpf_cpu_rq'"
description: "This page documents the 'scx_bpf_cpu_rq' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cpu_rq`

<!-- [FEATURE_TAG](scx_bpf_cpu_rq) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/6203ef73fa5c0358f7960b038628259be1448724)
<!-- [/FEATURE_TAG] -->

This function fetches the <nospell>rq</nospell> (run queue) of a CPU.

!!! warning
    This kfunc was deprecated in [:octicons-tag-24: v6.18](https://github.com/torvalds/linux/commit/5c48d88fe004988ec508923dbdd00549e65f4055) and removed in [:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/37cca059804ee35782a1cc9d7b087e5d3c855eae), it is no longer available. Use [`scx_bpf_locked_rq`](scx_bpf_locked_rq.md) to access the locked run queue, or [`scx_bpf_cpu_curr`](scx_bpf_cpu_curr.md) to look up the current task of a remote CPU instead.

## Definition

**Parameters**

`cpu`: CPU of the <nospell>rq</nospell>

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct rq *scx_bpf_cpu_rq(s32 cpu)`
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

