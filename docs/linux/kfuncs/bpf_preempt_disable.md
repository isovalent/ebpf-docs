---
title: "KFunc 'bpf_preempt_disable'"
description: "This page documents the 'bpf_preempt_disable' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_preempt_disable`

<!-- [FEATURE_TAG](bpf_preempt_disable) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/fc7566ad0a826cdc8886c5dbbb39ce72a0dc6333)
<!-- [/FEATURE_TAG] -->

Disable preemption in eBPF programs.

## Definition

<!-- [KFUNC_DEF] -->
`#!c void bpf_preempt_disable()`
<!-- [/KFUNC_DEF] -->

## Usage

This kfuncs allow disabling preemption in BPF programs. Nesting is allowed, since the intended use cases includes building native BPF spin locks without kernel helper involvement. Apart from that, this can be used to per-CPU data structures for cases where programs (or userspace) may preempt one or the other. Currently, while per-CPU access is stable, whether it will be consistent is not guaranteed, as only migration is disabled for BPF programs.

Global functions are disallowed from being called in non-preemptible regions. Static subprog calls are permitted. Sleepable helpers and kfuncs are disallowed in non-preemptible regions.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [BPF_PROG_TYPE_NETFILTER](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

