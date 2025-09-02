---
title: "Helper Function 'bpf_get_current_cgroup_id'"
description: "This page documents the 'bpf_get_current_cgroup_id' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_current_cgroup_id`

<!-- [FEATURE_TAG](bpf_get_current_cgroup_id) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/bf6fa2c893c5237b48569a13fa3c673041430b6c)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get the current cgroup id based on the cgroup within which the current task is running.

### Returns

A 64-bit integer containing the current cgroup id based on the cgroup within which the current task is running.

`#!c static __u64 (* const bpf_get_current_cgroup_id)(void) = (void *) 80;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

The `bpf_get_current_cgroup_id` helper function retrieves the cGroup ID of the cGroup in which the current task is running. This ID corresponds to the cGroup's file descriptor in the cGroup filesystem (`/sys/fs/cgroup`) and uniquely identifies a cGroup. It may be used to distinguish between containers, as container runtimes rely on cGroups for resource isolation and attribute a unique cGroup to each container. This helper function also enables enforcing cGroup-specific policies.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>

SEC("lsm_cgroup/inode_create")
int BPF_PROG(lsm_pre_bpf_file) {
    __u64 cgroup_id = bpf_get_current_cgroup_id();
    if (cgroup_id == 12092) {
        bpf_printk("Task from the target cgroup has created an inode!\n");
    }
    return 0;
}
```
