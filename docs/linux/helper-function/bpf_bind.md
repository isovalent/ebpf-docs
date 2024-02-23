---
title: "Helper Function 'bpf_bind'"
description: "This page documents the 'bpf_bind' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_bind`

<!-- [FEATURE_TAG](bpf_bind) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/d74bad4e74ee373787a9ae24197c17b7cdc428d5)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.

<!-- [HELPER_FUNC_DEF] -->
Bind the socket associated to _ctx_ to the address pointed by _addr_, of length _addr_len_. This allows for making outgoing connection from the desired IP address, which can be useful for example when all processes inside a cgroup should use one single IP address on a host that has multiple IP configured.

This helper works for IPv4 and IPv6, TCP and UDP sockets. The domain (_addr_**->sa_family**) must be **AF_INET** (or **AF_INET6**). It's advised to pass zero port (**sin_port** or **sin6_port**) which triggers IP_BIND_ADDRESS_NO_PORT-like behavior and lets the kernel efficiently pick up an unused port as long as 4-tuple is unique. Passing non-zero port might lead to degraded performance.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_bind)(struct bpf_sock_addr *ctx, struct sockaddr *addr, int addr_len) = (void *) 64;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
