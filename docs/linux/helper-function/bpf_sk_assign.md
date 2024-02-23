---
title: "Helper Function 'bpf_sk_assign' - eBPF Docs"
description: "This page documents the 'bpf_sk_assign' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_assign`

<!-- [FEATURE_TAG](bpf_sk_assign) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/cf7fbe660f2dbd738ab58aea8e9b0ca6ad232449)
<!-- [/FEATURE_TAG] -->

This helper function is used to direct incoming traffic to a specific socket, overruling the normal socket selection logic.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


`#!c static long (*bpf_sk_assign)(void *ctx, void *sk, __u64 flags) = (void *) 124;`

### Return value in `BPF_PROG_TYPE_SCHED_CLS` and `BPF_PROG_TYPE_SCHED_ACT`

0 on success, or a negative error in case of failure:

* `-EINVAL` if specified `flags` are not supported.
* `-ENOENT` if the socket is unavailable for assignment.
* `-ENETUNREACH` if the socket is unreachable (wrong netns).
* `-EOPNOTSUPP` if the operation is not supported, for example a call from outside of TC ingress.
* `-ESOCKTNOSUPPORT` if the socket type is not supported (reuseport).

### Return value in `BPF_PROG_TYPE_SK_LOOKUP` returns

0 on success, or a negative errno in case of failure.

* `-EAFNOSUPPORT` if socket family (`sk->family`) is not compatible with packet family ([`ctx->family`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md#family)).
* `-EEXIST` if socket has been already selected, potentially by another program, and `BPF_SK_LOOKUP_F_REPLACE` flag was not specified.
* `-EINVAL` if unsupported flags were specified.
* `-EPROTOTYPE` if socket L4 protocol (`sk->protocol`) doesn't match packet protocol ([`ctx->protocol`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md#protocol)).
* `-ESOCKTNOSUPPORT` if socket is not in allowed state (TCP listening or UDP unconnected).

## Usage

This helper is overloaded and has slightly different semantics in different program types.

### In `BPF_PROG_TYPE_SCHED_CLS` and `BPF_PROG_TYPE_SCHED_ACT` programs

This helper assign the `sk` to the `skb`. When combined with appropriate routing configuration to receive the packet towards the socket, will cause `skb` to be delivered to the specified socket. Subsequent redirection of `skb` via [`bpf_redirect`](bpf_redirect.md), [`bpf_clone_redirect`](bpf_clone_redirect.md) or other methods outside of BPF may interfere with successful delivery to the socket.

This operation is only valid from TC ingress path.

The `flags` argument must be zero.

### In `BPF_PROG_TYPE_SK_LOOKUP` programs

Select the `sk` as a result of a socket lookup.

For the operation to succeed passed socket must be compatible with the packet description provided by the [`ctx`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md#context) object.
	
L4 protocol (`IPPROTO_TCP` or `IPPROTO_UDP`) must be an exact match. While IP family (`AF_INET` or `AF_INET6`) must be compatible, that is IPv6 sockets that are not v6-only can be selected for IPv4 packets.
	
Only TCP listeners and UDP unconnected sockets can be selected. `sk` can also be NULL to reset any previous selection.
	
`flags` argument can combination of following values:
	
* `BPF_SK_LOOKUP_F_REPLACE` - to override the previous socket selection, potentially done by a BPF program that ran before us.	
* `BPF_SK_LOOKUP_F_NO_REUSEPORT` - to skip load-balancing within reuseport group for the socket being selected.
	
On success [`ctx->sk`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md#sk) will point to the selected socket.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Examples

Socket lookup program example usage
```c
// Copyright (c) 2020 Cloudflare

struct {
	__uint(type, BPF_MAP_TYPE_SOCKMAP);
	__uint(max_entries, 32);
	__type(key, __u32);
	__type(value, __u64);
} redir_map SEC(".maps");

static const __u16 DST_PORT = 7007; /* Host byte order */
static const __u32 DST_IP4 = IP4(127, 0, 0, 1);
static const __u32 KEY_SERVER_A = 0;

/* Redirect packets destined for DST_IP4 address to socket at redir_map[0]. */
SEC("sk_lookup")
int redir_ip4(struct bpf_sk_lookup *ctx)
{
	struct bpf_sock *sk;
	int err;

	if (ctx->family != AF_INET)
		return SK_PASS;
	if (ctx->local_port != DST_PORT)
		return SK_PASS;
	if (ctx->local_ip4 != DST_IP4)
		return SK_PASS;

	sk = bpf_map_lookup_elem(&redir_map, &KEY_SERVER_A);
	if (!sk)
		return SK_PASS;

	err = bpf_sk_assign(ctx, sk, 0);
	bpf_sk_release(sk);
	return err ? SK_DROP : SK_PASS;
}
```

Traffic control classifier example usage
```c
// SPDX-License-Identifier: GPL-2.0
// Copyright (c) 2019 Cloudflare Ltd.
// Copyright (c) 2020 Isovalent, Inc.

#include <stddef.h>
#include <stdbool.h>
#include <string.h>
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/in.h>
#include <linux/ip.h>
#include <linux/ipv6.h>
#include <linux/pkt_cls.h>
#include <linux/tcp.h>
#include <sys/socket.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>

/* Pin map under /sys/fs/bpf/tc/globals/<map name> */
#define PIN_GLOBAL_NS 2

/* Must match struct bpf_elf_map layout from iproute2 */
struct {
	__u32 type;
	__u32 size_key;
	__u32 size_value;
	__u32 max_elem;
	__u32 flags;
	__u32 id;
	__u32 pinning;
} server_map SEC("maps") = {
	.type = BPF_MAP_TYPE_SOCKMAP,
	.size_key = sizeof(int),
	.size_value  = sizeof(__u64),
	.max_elem = 1,
	.pinning = PIN_GLOBAL_NS,
};

char _license[] SEC("license") = "GPL";

/* Fill 'tuple' with L3 info, and attempt to find L4. On fail, return NULL. */
static inline struct bpf_sock_tuple *
get_tuple(struct __sk_buff *skb, bool *ipv4, bool *tcp)
{
	void *data_end = (void *)(long)skb->data_end;
	void *data = (void *)(long)skb->data;
	struct bpf_sock_tuple *result;
	struct ethhdr *eth;
	__u64 tuple_len;
	__u8 proto = 0;
	__u64 ihl_len;

	eth = (struct ethhdr *)(data);
	if (eth + 1 > data_end)
		return NULL;

	if (eth->h_proto == bpf_htons(ETH_P_IP)) {
		struct iphdr *iph = (struct iphdr *)(data + sizeof(*eth));

		if (iph + 1 > data_end)
			return NULL;
		if (iph->ihl != 5)
			/* Options are not supported */
			return NULL;
		ihl_len = iph->ihl * 4;
		proto = iph->protocol;
		*ipv4 = true;
		result = (struct bpf_sock_tuple *)&iph->saddr;
	} else if (eth->h_proto == bpf_htons(ETH_P_IPV6)) {
		struct ipv6hdr *ip6h = (struct ipv6hdr *)(data + sizeof(*eth));

		if (ip6h + 1 > data_end)
			return NULL;
		ihl_len = sizeof(*ip6h);
		proto = ip6h->nexthdr;
		*ipv4 = false;
		result = (struct bpf_sock_tuple *)&ip6h->saddr;
	} else {
		return (struct bpf_sock_tuple *)data;
	}

	if (proto != IPPROTO_TCP && proto != IPPROTO_UDP)
		return NULL;

	*tcp = (proto == IPPROTO_TCP);
	return result;
}

static inline int
handle_udp(struct __sk_buff *skb, struct bpf_sock_tuple *tuple, bool ipv4)
{
	struct bpf_sock *sk;
	const int zero = 0;
	size_t tuple_len;
	__be16 dport;
	int ret;

	tuple_len = ipv4 ? sizeof(tuple->ipv4) : sizeof(tuple->ipv6);
	if ((void *)tuple + tuple_len > (void *)(long)skb->data_end)
		return TC_ACT_SHOT;

	sk = bpf_sk_lookup_udp(skb, tuple, tuple_len, BPF_F_CURRENT_NETNS, 0);
	if (sk)
		goto assign;

	dport = ipv4 ? tuple->ipv4.dport : tuple->ipv6.dport;
	if (dport != bpf_htons(4321))
		return TC_ACT_OK;

	sk = bpf_map_lookup_elem(&server_map, &zero);
	if (!sk)
		return TC_ACT_SHOT;

assign:
	ret = bpf_sk_assign(skb, sk, 0);
	bpf_sk_release(sk);
	return ret;
}

static inline int
handle_tcp(struct __sk_buff *skb, struct bpf_sock_tuple *tuple, bool ipv4)
{
	struct bpf_sock *sk;
	const int zero = 0;
	size_t tuple_len;
	__be16 dport;
	int ret;

	tuple_len = ipv4 ? sizeof(tuple->ipv4) : sizeof(tuple->ipv6);
	if ((void *)tuple + tuple_len > (void *)(long)skb->data_end)
		return TC_ACT_SHOT;

	sk = bpf_skc_lookup_tcp(skb, tuple, tuple_len, BPF_F_CURRENT_NETNS, 0);
	if (sk) {
		if (sk->state != BPF_TCP_LISTEN)
			goto assign;
		bpf_sk_release(sk);
	}

	dport = ipv4 ? tuple->ipv4.dport : tuple->ipv6.dport;
	if (dport != bpf_htons(4321))
		return TC_ACT_OK;

	sk = bpf_map_lookup_elem(&server_map, &zero);
	if (!sk)
		return TC_ACT_SHOT;

	if (sk->state != BPF_TCP_LISTEN) {
		bpf_sk_release(sk);
		return TC_ACT_SHOT;
	}

assign:
	ret = bpf_sk_assign(skb, sk, 0);
	bpf_sk_release(sk);
	return ret;
}

SEC("tc")
int bpf_sk_assign_test(struct __sk_buff *skb)
{
	struct bpf_sock_tuple *tuple;
	bool ipv4 = false;
	bool tcp = false;
	int tuple_len;
	int ret = 0;

	tuple = get_tuple(skb, &ipv4, &tcp);
	if (!tuple)
		return TC_ACT_SHOT;

	/* Note that the verifier socket return type for bpf_skc_lookup_tcp()
	 * differs from bpf_sk_lookup_udp(), so even though the C-level type is
	 * the same here, if we try to share the implementations they will
	 * fail to verify because we're crossing pointer types.
	 */
	if (tcp)
		ret = handle_tcp(skb, tuple, ipv4);
	else
		ret = handle_udp(skb, tuple, ipv4);

	return ret == 0 ? TC_ACT_OK : TC_ACT_SHOT;
}
```
