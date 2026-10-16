---
title: "Helper Function 'bpf_strtoul'"
description: "This page documents the 'bpf_strtoul' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_strtoul`

<!-- [FEATURE_TAG](bpf_strtoul) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/d7a4cb9b6705a89937d12c8158a35a3145dc967a)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Convert the initial part of the string from buffer _buf_ of size _buf_len_ to an unsigned long integer according to the given base and save the result in _res_.

The string may begin with an arbitrary amount of white space (as determined by **isspace**(3)).

Five least significant bits of _flags_ encode base, other bits are currently unused.

Base must be either 8, 10, 16 or 0 to detect it automatically similar to user space **strtoul**(3).

### Returns

Number of characters consumed on success. Must be positive but no more than _buf_len_.

**-EINVAL** if no valid digits were found or unsupported base was provided.

**-ERANGE** if resulting value was out of range.

`#!c static long (* const bpf_strtoul)(const char *buf, unsigned long buf_len, __u64 flags, unsigned long *res) = (void *) 106;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md) [:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/8a67f2de9b1dc3cf8b75b4bf589efb1f08e3e9b8)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
