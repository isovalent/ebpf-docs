---
title: "Helper Function 'bpf_tcp_gen_syncookie'"
description: "This page documents the 'bpf_tcp_gen_syncookie' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_tcp_gen_syncookie`

<!-- [FEATURE_TAG](bpf_tcp_gen_syncookie) -->
[:octicons-tag-24: v5.4](https://github.com/torvalds/linux/commit/70d66244317e958092e9c971b08dd5b7fd29d9cb)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Try to issue a SYN cookie for the packet with corresponding IP/TCP headers, _iph_ and _th_, on the listening socket in _sk_.

_iph_ points to the start of the IPv4 or IPv6 header, while _iph_len_ contains **sizeof**(**struct iphdr**) or **sizeof**(**struct ipv6hdr**).

_th_ points to the start of the TCP header, while _th_len_ contains the length of the TCP header with options (at least **sizeof**(**struct tcphdr**)).

### Returns

On success, lower 32 bits hold the generated SYN cookie in followed by 16 bits which hold the MSS value for that cookie, and the top 16 bits are unused.

On failure, the returned value is one of the following:

**-EINVAL** SYN cookie cannot be issued due to error

**-ENOENT** SYN cookie should not be issued (no SYN flood)

**-EOPNOTSUPP** kernel configuration does not enable SYN cookies

**-EPROTONOSUPPORT** IP packet version is not 4 or 6

`#!c static __s64 (*bpf_tcp_gen_syncookie)(void *sk, void *iph, __u32 iph_len, struct tcphdr *th, __u32 th_len) = (void *) 110;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
