---
title: "Helper Function 'bpf_tcp_raw_gen_syncookie_ipv4'"
description: "This page documents the 'bpf_tcp_raw_gen_syncookie_ipv4' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_tcp_raw_gen_syncookie_ipv4`

<!-- [FEATURE_TAG](bpf_tcp_raw_gen_syncookie_ipv4) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/33bf9885040c399cf6a95bd33216644126728e14)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Try to issue a SYN cookie for the packet with corresponding IPv4/TCP headers, _iph_ and _th_, without depending on a listening socket.

_iph_ points to the IPv4 header.

_th_ points to the start of the TCP header, while _th_len_ contains the length of the TCP header (at least **sizeof**(**struct tcphdr**)).

### Returns

On success, lower 32 bits hold the generated SYN cookie in followed by 16 bits which hold the MSS value for that cookie, and the top 16 bits are unused.

On failure, the returned value is one of the following:

**-EINVAL** if _th_len_ is invalid.

`#!c static __s64 (* const bpf_tcp_raw_gen_syncookie_ipv4)(struct iphdr *iph, struct tcphdr *th, __u32 th_len) = (void *) 204;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
