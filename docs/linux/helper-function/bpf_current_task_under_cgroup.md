---
title: "Helper Function 'bpf_current_task_under_cgroup'"
description: "This page documents the 'bpf_current_task_under_cgroup' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_current_task_under_cgroup`

<!-- [FEATURE_TAG](bpf_current_task_under_cgroup) -->
[:octicons-tag-24: v4.9](https://github.com/torvalds/linux/commit/60d20f9195b260bdf0ac10c275ae9f6016f9c069)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Check whether the probe is being run is the context of a given subset of the cgroup2 hierarchy. The cgroup2 to test is held by _map_ of type **BPF_MAP_TYPE_CGROUP_ARRAY**, at _index_.

### Returns

The return value depends on the result of the test, and can be:

* 1, if current task belongs to the cgroup2.
* 0, if current task does not belong to the cgroup2.
* A negative error code, if an error occurred.

`#!c static long (* const bpf_current_task_under_cgroup)(void *map, __u32 index) = (void *) 37;`
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
