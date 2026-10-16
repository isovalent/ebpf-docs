---
title: "Helper Function 'bpf_get_current_comm'"
description: "This page documents the 'bpf_get_current_comm' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_current_comm`

<!-- [FEATURE_TAG](bpf_get_current_comm) -->
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/ffeedafbf0236f03aeb2e8db273b3e5ae5f5bc89)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Copy the **comm** attribute of the current task into _buf_ of _size_of_buf_. The **comm** attribute contains the name of the executable (excluding the path) for the current task. The _size_of_buf_ must be strictly positive. On success, the helper makes sure that the _buf_ is NUL-terminated. On failure, it is filled with zeroes.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_get_current_comm)(void *buf, __u32 size_of_buf) = (void *) 16;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

The `bpf_get_current_comm` helper function retrieves the name of the executable associated with the current task. This is useful for identifying the process context in which the eBPF program is executing, enabling per-process tracing. It can help trace specific applications, enforce process-level policies, or monitor system behavior tied to particular commands.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>

SEC("tp/syscalls/sys_enter_open")
int sys_open_trace(void *ctx) {
  // TASK_COMM_LEN is defined in vmlinux.h
  char comm[TASK_COMM_LEN];
  if (bpf_get_current_comm(comm, TASK_COMM_LEN)) {
    bpf_printk("Failed to get comm\n");
    return 0;
  }
  bpf_printk("Hello from %s\n", comm);
  return 0;
}
```
