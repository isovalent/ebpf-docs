# Helper function `bpf_tcp_check_syncookie`

<!-- [FEATURE_TAG](bpf_tcp_check_syncookie) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/399040847084a69f345e0a52fd62f04654e0fce3)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Check whether _iph_ and _th_ contain a valid SYN cookie ACK for the listening socket in _sk_.

_iph_ points to the start of the IPv4 or IPv6 header, while _iph_len_ contains **sizeof**(**struct iphdr**) or **sizeof**(**struct ipv6hdr**).

_th_ points to the start of the TCP header, while _th_len_ contains the length of the TCP header (at least **sizeof**(**struct tcphdr**)).

### Returns

0 if _iph_ and _th_ are a valid SYN cookie ACK, or a negative error otherwise.

`#!c static long (*bpf_tcp_check_syncookie)(void *sk, void *iph, __u32 iph_len, struct tcphdr *th, __u32 th_len) = (void *) 100;`
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
