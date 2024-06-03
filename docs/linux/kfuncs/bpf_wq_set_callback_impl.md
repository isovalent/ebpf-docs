---
title: "KFunc 'bpf_wq_set_callback_impl'"
description: "This page documents the 'bpf_wq_set_callback_impl' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_wq_set_callback_impl`

<!-- [FEATURE_TAG](bpf_wq_set_callback_impl) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/81f1d7a583fa1fa14f0c4e6140d34b5e3d08d227)
<!-- [/FEATURE_TAG] -->

Set a callback function for a workqueue.

## Definition

This kfunc associates a callback function with a workqueue. The workqueue must be initialized with the [`bpf_wq_init`](bpf_wq_init.md) kfunc before calling this function. After the callback function is set, work can be scheduled using the [`bpf_wq_start`](bpf_wq_start.md) kfunc.

The callback will be called asynchronously sometime after the current eBPF program has finished executing whenever the scheduler decides to run the workqueue.

`wq`: A pointer to a `struct bpf_wq` which must reside in a map value.

`callback_fn`: The callback function to be called when the workqueue is run. The callback function must have the following signature: `#!c int (callback_fn)(void *map, int *key, struct bpf_wq *wq)`

**Returns**

Return `0` on success, or a negative error code on failure.

<!-- [KFUNC_DEF] -->
`#!c int bpf_wq_set_callback_impl(struct bpf_wq *wq, int (callback_fn)(void * , int * , struct bpf_wq * ), unsigned int flags, void *aux__ign)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

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

