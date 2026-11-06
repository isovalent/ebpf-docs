---
title: "KFunc 'bpf_icmp_send'"
description: "This page documents the 'bpf_icmp_send' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_icmp_send`

<!-- [FEATURE_TAG](bpf_icmp_send) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/f3603df9aebb2a2fe2f745bd71ca38aeca60e6e7)
<!-- [/FEATURE_TAG] -->

Send an ICMP control message in response to a packet.

## Definition

Sends an ICMP control message in response to the packet `skb_ctx`. The original packet is cloned before sending the ICMP message, so the packet itself is left untouched and the BPF program can still decide to drop or pass it.

Currently only the "destination unreachable" message type is supported, `ICMP_DEST_UNREACH` (`3`) for IPv4 packets and `ICMPV6_DEST_UNREACH` (`1`) for IPv6 packets. The `type` parameter exists to allow for future extension to other ICMP message types.

`skb_ctx`: The packet that triggered the control message, which is the context of the program.

`type`: The ICMP type. Must be `ICMP_DEST_UNREACH` for IPv4 packets or `ICMPV6_DEST_UNREACH` for IPv6 packets.

`code`: The ICMP code. For IPv4 any code from `0` to `15` (`NR_ICMP_UNREACH`) except `ICMP_FRAG_NEEDED` (`4`), since that code requires a valid next-hop MTU. For IPv6 any code from `0` to `6` (`ICMPV6_REJECT_ROUTE`).

**Returns**

Returns `0` if the ICMP message was submitted for sending, or a negative error code on failure:

- `-EBUSY`: Recursion detected, the packet was sent by the kernel ICMP socket itself.
- `-EPROTONOSUPPORT`: The packet is not an IPv4 or IPv6 packet.
- `-EOPNOTSUPP`: Unsupported ICMP `type`.
- `-EINVAL`: Invalid ICMP `code`.
- `-ENETUNREACH`: The packet has no usable route to send the ICMP reply on.
- `-ENOMEM`: Failed to allocate memory for the clone of the packet.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_icmp_send(struct __sk_buff *skb_ctx, int type, int code)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc allows a cGroup socket buffer program to actively reject traffic, similar to the `REJECT` target of netfilter. Instead of silently dropping a packet, which causes the sending application to retry until it times out, the program can send an ICMP "destination unreachable" message back. This gives the application early feedback, for example the client TCP stack will abort the connection attempt instead of retrying. A typical use case is enforcing network policies on east-west traffic.

Unlike [`bpf_sock_destroy`](bpf_sock_destroy.md), which aborts a socket and sends a TCP reset, this kfunc does not affect the socket, it only sends an ICMP message.

The program is expected to drop the packet after calling this kfunc, by returning `SK_DROP`. Note that nothing prevents a program from calling this kfunc multiple times, which sends multiple ICMP messages, or from passing the packet after sending an ICMP message, which can lead to a confusing situation where a connection is established while the client also received an ICMP "destination unreachable" message.

Since the ICMP message is sent by the kernel ICMP socket, it may trigger the same program again, for example when the program is attached to the root cGroup. To protect against recursion, this kfunc returns `-EBUSY` when it is called for a packet originating from the kernel ICMP socket.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

This example rejects outgoing TCP connections to `127.0.0.1:server_port` and `[::1]:server_port` by sending an ICMP "destination unreachable" message and dropping the packet.

Example based on [`tools/testing/selftests/bpf/progs/icmp_send.c`](https://github.com/torvalds/linux/blob/c03c1a14507dfaf8b856c2608f9ac216021fc222/tools/testing/selftests/bpf/progs/icmp_send.c)

!!! example
	```c
	// SPDX-License-Identifier: GPL-2.0
	#include "vmlinux.h"
	#include <bpf/bpf_helpers.h>
	#include <bpf/bpf_endian.h>

	/* 127.0.0.1 in host byte order */
	#define SERVER_IP 0x7F000001
	/* ::1 in host byte order (last 32-bit word) */
	#define SERVER_IP6_LO 0x00000001

	extern int bpf_icmp_send(struct __sk_buff *skb_ctx, int type, int code) __ksym;

	__u16 server_port = 0;
	int unreach_type = 0;
	int unreach_code = 0;
	int kfunc_ret = -1;

	SEC("cgroup_skb/egress")
	int egress(struct __sk_buff *skb)
	{
		void *data = (void *)(long)skb->data;
		void *data_end = (void *)(long)skb->data_end;
		struct iphdr *iph;
		struct ipv6hdr *ip6h;
		struct tcphdr *tcph;
		__u8 version;

		if (data + 1 > data_end)
			return SK_PASS;

		version = (*((__u8 *)data)) >> 4;

		if (version == 4) {
			iph = data;
			if ((void *)(iph + 1) > data_end ||
			    iph->protocol != IPPROTO_TCP ||
			    iph->daddr != bpf_htonl(SERVER_IP))
				return SK_PASS;

			tcph = (void *)iph + iph->ihl * 4;
			if ((void *)(tcph + 1) > data_end ||
			    tcph->dest != bpf_htons(server_port))
				return SK_PASS;

		} else if (version == 6) {
			ip6h = data;
			if ((void *)(ip6h + 1) > data_end ||
			    ip6h->nexthdr != IPPROTO_TCP)
				return SK_PASS;

			if (ip6h->daddr.in6_u.u6_addr32[0] != 0 ||
			    ip6h->daddr.in6_u.u6_addr32[1] != 0 ||
			    ip6h->daddr.in6_u.u6_addr32[2] != 0 ||
			    ip6h->daddr.in6_u.u6_addr32[3] != bpf_htonl(SERVER_IP6_LO))
				return SK_PASS;

			tcph = (void *)(ip6h + 1);
			if ((void *)(tcph + 1) > data_end ||
			    tcph->dest != bpf_htons(server_port))
				return SK_PASS;
		} else {
			return SK_PASS;
		}

		kfunc_ret = bpf_icmp_send(skb, unreach_type, unreach_code);

		return SK_DROP;
	}

	char LICENSE[] SEC("license") = "Dual BSD/GPL";
	```
