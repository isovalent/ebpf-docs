---
title: "KFunc 'bpf_rcu_read_lock' - eBPF Docs"
description: "This page documents the 'bpf_rcu_read_lock' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_rcu_read_lock`

<!-- [FEATURE_TAG](bpf_rcu_read_lock) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/9bb00b2895cbfe0ad410457b605d0a72524168c1)
<!-- [/FEATURE_TAG] -->

## Definition

This kfunc is used to define a RCU read lock region in the BPF program.
The end of such a region is marked by [`bpf_rcu_read_unlock`](bpf_rcu_read_unlock.md)

The current implementation does not support nested rcu read lock
region in the prog.

<!-- [KFUNC_DEF] -->
`#!c void bpf_rcu_read_lock()`
<!-- [/KFUNC_DEF] -->

## Usage

```c
  struct task_struct {
    ...
    struct task_struct              *last_wakee;
    struct task_struct __rcu        *real_parent;
    ...
  };
```

Let us say prog does `task = bpf_get_current_task_btf()` to get a
`task` pointer. The basic rules are:

  - '`real_parent = task->real_parent` should be inside `bpf_rcu_read_lock`
    region. This is to simulate `rcu_dereference()` operation. The
    `real_parent` is marked as `MEM_RCU` only if (1). `task->real_parent` is
    inside `bpf_rcu_read_lock` region, and (2). task is a trusted ptr. So
    MEM_RCU marked ptr can be 'trusted' inside the `bpf_rcu_read_lock` region.
  - `last_wakee = real_parent->last_wakee` should be inside `bpf_rcu_read_lock`
    region since it tries to access rcu protected memory.
  - the ptr 'last_wakee' will be marked as PTR_UNTRUSTED since in general
    it is not clear whether the object pointed by 'last_wakee' is valid or
    not even inside `bpf_rcu_read_lock` region.

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

