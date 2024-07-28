---
title: "KFunc 'bpf_preempt_enable'"
description: "This page documents the 'bpf_preempt_enable' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_preempt_enable`

<!-- [FEATURE_TAG](bpf_preempt_enable) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/fc7566ad0a826cdc8886c5dbbb39ce72a0dc6333)
<!-- [/FEATURE_TAG] -->

Enable preemption in eBPF programs.

## Definition

<!-- [KFUNC_DEF] -->
`#!c void bpf_preempt_enable()`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc re-enables preemption in BPF programs. Nesting is allowed. See [`bpf_preempt_disable`](bpf_preempt_disable.md) for more information.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

