---
title: "Helper Function 'bpf_send_signal_thread'"
description: "This page documents the 'bpf_send_signal_thread' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_send_signal_thread`

<!-- [FEATURE_TAG](bpf_send_signal_thread) -->
[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/8482941f09067da42f9c3362e15bfb3f3c19d610)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Send signal _sig_ to the thread corresponding to the current task.

### Returns

0 on success or successfully queued.

**-EBUSY** if work queue under nmi is full.

**-EINVAL** if _sig_ is invalid.

**-EPERM** if no permission to send the _sig_.

**-EAGAIN** if bpf program can try again.

`#!c static long (* const bpf_send_signal_thread)(__u32 sig) = (void *) 117;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/ee971630f20fd421fffcdc4543731ebcb54ed6d0)
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

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
