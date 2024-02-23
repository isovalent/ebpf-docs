---
title: "Helper Function 'bpf_sk_lookup_tcp'"
description: "This page documents the 'bpf_sk_lookup_tcp' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_lookup_tcp`

<!-- [FEATURE_TAG](bpf_sk_lookup_tcp) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/6acc9b432e6714d72d7d77ec7c27f6f8358d0c71)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Look for TCP socket matching _tuple_, optionally in a child network namespace _netns_. The return value must be checked, and if non-**NULL**, released via **bpf_sk_release**().

The _ctx_ should point to the context of the program, such as the skb or socket (depending on the hook in use). This is used to determine the base network namespace for the lookup.

_tuple_size_ must be one of:

**sizeof**(_tuple_**->ipv4**)

&nbsp;&nbsp;&nbsp;&nbsp;Look for an IPv4 socket.

**sizeof**(_tuple_**->ipv6**)

&nbsp;&nbsp;&nbsp;&nbsp;Look for an IPv6 socket.

If the _netns_ is a negative signed 32-bit integer, then the socket lookup table in the netns associated with the _ctx_ will be used. For the TC hooks, this is the netns of the device in the skb. For socket hooks, this is the netns of the socket. If _netns_ is any other signed 32-bit value greater than or equal to zero then it specifies the ID of the netns relative to the netns associated with the _ctx_. _netns_ values beyond the range of 32-bit integers are reserved for future use.

All values for _flags_ are reserved for future usage, and must be left at zero.

This helper is available only if the kernel was compiled with **CONFIG_NET** configuration option.

### Returns

Pointer to **struct bpf_sock**, or **NULL** in case of failure. For sockets with reuseport option, the **struct bpf_sock** result is from _reuse_**->socks**[] using the hash of the tuple.

`#!c static struct bpf_sock *(*bpf_sk_lookup_tcp)(void *ctx, struct bpf_sock_tuple *tuple, __u32 tuple_size, __u64 netns, __u64 flags) = (void *) 84;`
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
