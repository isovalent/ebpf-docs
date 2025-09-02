---
title: "Helper Function 'bpf_get_current_ancestor_cgroup_id'"
description: "This page documents the 'bpf_get_current_ancestor_cgroup_id' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_current_ancestor_cgroup_id`

<!-- [FEATURE_TAG](bpf_get_current_ancestor_cgroup_id) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/0f09abd105da6c37713d2b253730a86cb45e127a)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return id of cgroup v2 that is ancestor of the cgroup associated with the current task at the _ancestor_level_. The root cgroup is at _ancestor_level_ zero and each step down the hierarchy increments the level. If _ancestor_level_ == level of cgroup associated with the current task, then return value will be the same as that of **bpf_get_current_cgroup_id**().

The helper is useful to implement policies based on cgroups that are upper in hierarchy than immediate cgroup associated with the current task.

The format of returned id and helper limitations are same as in **bpf_get_current_cgroup_id**().

### Returns

The id is returned or 0 in case the id could not be retrieved.

`#!c static __u64 (* const bpf_get_current_ancestor_cgroup_id)(int ancestor_level) = (void *) 123;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c501bf55c88b834adefda870c7c092ec9052a437)
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

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
