---
title: "Helper Function 'bpf_get_ns_current_pid_tgid'"
description: "This page documents the 'bpf_get_ns_current_pid_tgid' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_ns_current_pid_tgid`

<!-- [FEATURE_TAG](bpf_get_ns_current_pid_tgid) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/b4490c5c4e023f09b7d27c9a9d3e7ad7d09ea6bf)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Returns 0 on success, values for _pid_ and _tgid_ as seen from the current _namespace_ will be returned in _nsdata_.

### Returns

0 on success, or one of the following in case of failure:

**-EINVAL** if dev and inum supplied don't match dev_t and inode number with nsfs of current task, or if dev conversion to dev_t lost high bits.

**-ENOENT** if pidns does not exists for the current task.

`#!c static long (* const bpf_get_ns_current_pid_tgid)(__u64 dev, __u64 ino, struct bpf_pidns_info *nsdata, __u32 size) = (void *) 120;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
