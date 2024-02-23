---
title: "KFunc 'bpf_dynptr_adjust'"
description: "This page documents the 'bpf_dynptr_adjust' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_dynptr_adjust`

<!-- [FEATURE_TAG](bpf_dynptr_adjust) -->
[:octicons-tag-24: v6.5](https://github.com/torvalds/linux/commit/987d0242d189661f78b77cc4d77f843b15600fed)
<!-- [/FEATURE_TAG] -->

Adjusts the dynptr to reflect the new [start, end) interval.

## Definition

It advances the offset of the dynptr by `start` bytes, and if end is less than the size of the dynptr, then this will trim the dynptr accordingly.

<!-- [KFUNC_DEF] -->
`#!c int bpf_dynptr_adjust(struct bpf_dynptr_kern *ptr, u32 start, u32 end)`
<!-- [/KFUNC_DEF] -->

## Usage

Adjusting the dynptr interval may be useful in certain situations. For example, when hashing which takes in generic dynptrs, if the dynptr points to a struct but only a certain memory region inside the struct should be hashed, adjust can be used to narrow in on the specific region to hash.

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

