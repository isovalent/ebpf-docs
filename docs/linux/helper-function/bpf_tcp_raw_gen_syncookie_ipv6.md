# Helper function `bpf_tcp_raw_gen_syncookie_ipv6`

<!-- [FEATURE_TAG](bpf_tcp_raw_gen_syncookie_ipv6) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/33bf9885040c399cf6a95bd33216644126728e14)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Try to issue a SYN cookie for the packet with corresponding IPv6/TCP headers, _iph_ and _th_, without depending on a listening socket.

_iph_ points to the IPv6 header.

_th_ points to the start of the TCP header, while _th_len_ contains the length of the TCP header (at least **sizeof**(**struct tcphdr**)).

### Returns

On success, lower 32 bits hold the generated SYN cookie in followed by 16 bits which hold the MSS value for that cookie, and the top 16 bits are unused.

On failure, the returned value is one of the following:

**-EINVAL** if _th_len_ is invalid.

**-EPROTONOSUPPORT** if CONFIG_IPV6 is not builtin.

`#!c static __s64 (*bpf_tcp_raw_gen_syncookie_ipv6)(struct ipv6hdr *iph, struct tcphdr *th, __u32 th_len) = (void *) 205;`
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
