---
title: "Helper Function 'bpf_skc_lookup_tcp'"
description: "This page documents the 'bpf_skc_lookup_tcp' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_skc_lookup_tcp`

<!-- [FEATURE_TAG](bpf_skc_lookup_tcp) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/edbf8c01de5a104a71ed6df2bf6421ceb2836a8e)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Look for TCP socket matching _tuple_, optionally in a child network namespace _netns_. The return value must be checked, and if non-**NULL**, released via **bpf_sk_release**().

This function is identical to **bpf_sk_lookup_tcp**(), except that it also returns timewait or request sockets. Use **bpf_sk_fullsock**() or **bpf_tcp_sock**() to access the full structure.

This helper is available only if the kernel was compiled with **CONFIG_NET** configuration option.

### Returns

Pointer to **struct bpf_sock**, or **NULL** in case of failure. For sockets with reuseport option, the **struct bpf_sock** result is from _reuse_**->socks**[] using the hash of the tuple.

`#!c static struct bpf_sock *(* const bpf_skc_lookup_tcp)(void *ctx, struct bpf_sock_tuple *tuple, __u32 tuple_size, __u64 netns, __u64 flags) = (void *) 99;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
