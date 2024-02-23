---
title: "Helper Function 'bpf_tcp_raw_check_syncookie_ipv6'"
description: "This page documents the 'bpf_tcp_raw_check_syncookie_ipv6' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_tcp_raw_check_syncookie_ipv6`

<!-- [FEATURE_TAG](bpf_tcp_raw_check_syncookie_ipv6) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/33bf9885040c399cf6a95bd33216644126728e14)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Check whether _iph_ and _th_ contain a valid SYN cookie ACK without depending on a listening socket.

_iph_ points to the IPv6 header.

_th_ points to the TCP header.

### Returns

0 if _iph_ and _th_ are a valid SYN cookie ACK.

On failure, the returned value is one of the following:

**-EACCES** if the SYN cookie is not valid.

**-EPROTONOSUPPORT** if CONFIG_IPV6 is not builtin.

`#!c static long (*bpf_tcp_raw_check_syncookie_ipv6)(struct ipv6hdr *iph, struct tcphdr *th) = (void *) 207;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
