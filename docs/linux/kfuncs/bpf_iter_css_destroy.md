---
title: "KFunc 'bpf_iter_css_destroy'"
description: "This page documents the 'bpf_iter_css_destroy' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_css_destroy`

<!-- [FEATURE_TAG](bpf_iter_css_destroy) -->
[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/7251d0905e7518bcb990c8e9a3615b1bb23c78f2)
<!-- [/FEATURE_TAG] -->

Destroy a cGroup iterator.

## Definition

<!-- [KFUNC_DEF] -->
`#!c void bpf_iter_css_destroy(struct bpf_iter_css *it)`
<!-- [/KFUNC_DEF] -->

## Usage

See [`bpf_iter_css_new`](bpf_iter_css_new.md#usage) for usage details.

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

See [`bpf_iter_css_new`](bpf_iter_css_new.md#example) an example.

