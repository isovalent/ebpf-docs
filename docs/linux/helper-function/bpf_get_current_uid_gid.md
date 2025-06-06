---
title: "Helper Function 'bpf_get_current_uid_gid'"
description: "This page documents the 'bpf_get_current_uid_gid' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_current_uid_gid`

<!-- [FEATURE_TAG](bpf_get_current_uid_gid) -->
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/ffeedafbf0236f03aeb2e8db273b3e5ae5f5bc89)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get the current uid and gid.

### Returns

A 64-bit integer containing the current GID and UID, and created as such: _current_gid_ **<< 32 \|** _current_uid_.

`#!c static __u64 (* const bpf_get_current_uid_gid)(void) = (void *) 15;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

The `bpf_get_current_uid_gid` helper function returns a 64-bit value containing the current task's UID in the lower 32 bits and GID in the upper 32 bits. This allows eBPF programs to identify the user and group context of the running task. It is useful for enforcing security policies, tracking actions by specific users or groups, and implementing per-UID or per-GID tracing.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example


```c
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>

SEC("tp/syscalls/sys_enter_open")
int sys_open_trace(void *ctx) {
    __u64 uid_gid = bpf_get_current_uid_gid();
    __u32 uid = uid_gid & 0xFFFFFFFF;
    __u32 gid = uid_gid >> 32;
    bpf_printk("Hello from UID %u, GID %u\n", uid, gid);
    return 0;
}
```
