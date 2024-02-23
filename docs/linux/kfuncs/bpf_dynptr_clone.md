---
title: "KFunc 'bpf_dynptr_clone' - eBPF Docs"
description: "This page documents the 'bpf_dynptr_clone' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_dynptr_clone`

<!-- [FEATURE_TAG](bpf_dynptr_clone) -->
[:octicons-tag-24: v6.5](https://github.com/torvalds/linux/commit/361f129f3cc185af6667aca0bec0be9a020a8abc)
<!-- [/FEATURE_TAG] -->

Clones a dynptr.

## Definition

The cloned dynptr will point to the same data as its parent dynptr, with the same type, offset, size and read-only properties.

Any writes to a dynptr will be reflected across all instances (by 'instance', this means any dynptrs that point to the same underlying data).

!!! note 
    data slice and dynptr invalidations will affect all instances as well. For example, if bpf_dynptr_write() is called on an skb-type dynptr, all data slices of dynptr instances to that skb will be invalidated as well (eg data slices of any clones, parents, grandparents, ...). Another example is if a ringbuf dynptr is submitted, any instance of that dynptr will be invalidated.

Changing the view of the dynptr (eg advancing the offset or trimming the size) will only affect that dynptr and not affect any other instances.

<!-- [KFUNC_DEF] -->
`#!c int bpf_dynptr_clone(struct bpf_dynptr_kern *ptr, struct bpf_dynptr_kern *clone__uninit)`
<!-- [/KFUNC_DEF] -->

## Usage

One example use case where cloning may be helpful is for hashing or iterating through dynptr data. Cloning will allow the user to maintain the original view of the dynptr for future use, while also allowing views to smaller subsets of the data after the offset is advanced or the size is trimmed.

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

