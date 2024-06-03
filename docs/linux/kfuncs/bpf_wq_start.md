---
title: "KFunc 'bpf_wq_start'"
description: "This page documents the 'bpf_wq_start' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_wq_start`

<!-- [FEATURE_TAG](bpf_wq_start) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/8e83da9732d91c60fdc651b2486c8e5935eb0ca2)
<!-- [/FEATURE_TAG] -->

Start a workqueue.

## Definition

This kfunc starts a workqueue which allows eBPF programs to schedule work to be executed asynchronously.

`wq`: A pointer to a `struct bpf_wq` which must reside in a map value.

`flags`: Flags to allow for future extensions.

**Returns**

Return `0` on success, or a negative error code on failure.

<!-- [KFUNC_DEF] -->
`#!c int bpf_wq_start(struct bpf_wq *wq, unsigned int flags)`
<!-- [/KFUNC_DEF] -->

## Usage

Once a workqueue has been initialized with the [`bpf_wq_init`](bpf_wq_init.md) kfunc and a callback function has been associated with it using the [`bpf_wq_set_callback_impl`](bpf_wq_set_callback_impl.md) kfunc, work can be scheduled using this function.

The callback will be called asynchronously sometime after the current eBPF program has finished executing whenever the scheduler decides to run the workqueue.

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

