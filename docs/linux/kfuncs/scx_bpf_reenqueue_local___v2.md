---
title: "KFunc 'scx_bpf_reenqueue_local___v2'"
description: "This page documents the 'scx_bpf_reenqueue_local___v2' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_reenqueue_local___v2`

<!-- [FEATURE_TAG](scx_bpf_reenqueue_local___v2) -->
[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/a3f5d48222532484c1e85ef27cc6893803e4cd17)
<!-- [/FEATURE_TAG] -->

This function re-enqueues tasks on a local DSQ.

## Definition

Iterate over all of the tasks currently enqueued on the local DSQ of the caller's CPU, and re-enqueue them in the BPF scheduler.

!!! warn
    This kfunc is meant as replacement for [`scx_bpf_reenqueue_local`](scx_bpf_reenqueue_local.md). Until kernel v7.3 (4 releases after v6.19) at which time the old implementation is deleted and renamed to `scx_bpf_reenqueue_local`, a breaking change in the function signature.
    Its recommended to use the [`scx_bpf_reenqueue_local`](../../ebpf-library/scx/scx_bpf_reenqueue_local.md) function from the SCX common library instead of defining the kfunc manually to facilitate smooth transition across kernel versions.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_reenqueue_local___v2()`
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

