---
title: "KFunc 'bpf_xdp_flow_lookup'"
description: "This page documents the 'bpf_xdp_flow_lookup' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_xdp_flow_lookup`

<!-- [FEATURE_TAG](bpf_xdp_flow_lookup) -->
[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/391bb6594fd3a567efb1cd3efc8136c78c4c9e31)
<!-- [/FEATURE_TAG] -->

Perform the lookup of a given flow-table entry based on a fib tuple of incoming traffic.

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct flow_offload_tuple_rhash *bpf_xdp_flow_lookup(struct xdp_md *ctx, struct bpf_fib_lookup *fib_tuple, struct bpf_flowtable_opts *opts, u32 opts_len)`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
<!-- [/KFUNC_DEF] -->

## Usage

`bpf_xdp_flow_lookup` can be used as building block to offload in xdp the processing of sw flowtable when hw flowtable is not available.

See also: [patch set](https://lore.kernel.org/bpf/cover.1719698275.git.lorenzo@kernel.org/)

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

[Source](https://github.com/torvalds/linux/blob/c77e572d3a8c0e21c5dca4cc2883c7cd8cbe981f/tools/testing/selftests/bpf/progs/xdp_flowtable.c)

```c
// SPDX-License-Identifier: GPL-2.0
#define BPF_NO_KFUNC_PROTOTYPES
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>

#define ETH_P_IP	0x0800
#define ETH_P_IPV6	0x86dd
#define IP_MF		0x2000	/* "More Fragments" */
#define IP_OFFSET	0x1fff	/* "Fragment Offset" */
#define AF_INET		2
#define AF_INET6	10

struct bpf_flowtable_opts___local {
	s32 error;
};

struct flow_offload_tuple_rhash *
bpf_xdp_flow_lookup(struct xdp_md *, struct bpf_fib_lookup *,
		    struct bpf_flowtable_opts___local *, u32) __ksym;

struct {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__type(key, __u32);
	__type(value, __u32);
	__uint(max_entries, 1);
} stats SEC(".maps");

static bool xdp_flowtable_offload_check_iphdr(struct iphdr *iph)
{
	/* ip fragmented traffic */
	if (iph->frag_off & bpf_htons(IP_MF | IP_OFFSET))
		return false;

	/* ip options */
	if (iph->ihl * 4 != sizeof(*iph))
		return false;

	if (iph->ttl <= 1)
		return false;

	return true;
}

static bool xdp_flowtable_offload_check_tcp_state(void *ports, void *data_end,
						  u8 proto)
{
	if (proto == IPPROTO_TCP) {
		struct tcphdr *tcph = ports;

		if (tcph + 1 > data_end)
			return false;

		if (tcph->fin || tcph->rst)
			return false;
	}

	return true;
}

struct flow_ports___local {
	__be16 source, dest;
} __attribute__((preserve_access_index));

SEC("xdp.frags")
int xdp_flowtable_do_lookup(struct xdp_md *ctx)
{
	void *data_end = (void *)(long)ctx->data_end;
	struct bpf_flowtable_opts___local opts = {};
	struct flow_offload_tuple_rhash *tuplehash;
	struct bpf_fib_lookup tuple = {
		.ifindex = ctx->ingress_ifindex,
	};
	void *data = (void *)(long)ctx->data;
	struct ethhdr *eth = data;
	struct flow_ports___local *ports;
	__u32 *val, key = 0;

	if (eth + 1 > data_end)
		return XDP_DROP;

	switch (eth->h_proto) {
	case bpf_htons(ETH_P_IP): {
		struct iphdr *iph = data + sizeof(*eth);

		ports = (struct flow_ports___local *)(iph + 1);
		if (ports + 1 > data_end)
			return XDP_PASS;

		/* sanity check on ip header */
		if (!xdp_flowtable_offload_check_iphdr(iph))
			return XDP_PASS;

		if (!xdp_flowtable_offload_check_tcp_state(ports, data_end,
							   iph->protocol))
			return XDP_PASS;

		tuple.family		= AF_INET;
		tuple.tos		= iph->tos;
		tuple.l4_protocol	= iph->protocol;
		tuple.tot_len		= bpf_ntohs(iph->tot_len);
		tuple.ipv4_src		= iph->saddr;
		tuple.ipv4_dst		= iph->daddr;
		tuple.sport		= ports->source;
		tuple.dport		= ports->dest;
		break;
	}
	case bpf_htons(ETH_P_IPV6): {
		struct in6_addr *src = (struct in6_addr *)tuple.ipv6_src;
		struct in6_addr *dst = (struct in6_addr *)tuple.ipv6_dst;
		struct ipv6hdr *ip6h = data + sizeof(*eth);

		ports = (struct flow_ports___local *)(ip6h + 1);
		if (ports + 1 > data_end)
			return XDP_PASS;

		if (ip6h->hop_limit <= 1)
			return XDP_PASS;

		if (!xdp_flowtable_offload_check_tcp_state(ports, data_end,
							   ip6h->nexthdr))
			return XDP_PASS;

		tuple.family		= AF_INET6;
		tuple.l4_protocol	= ip6h->nexthdr;
		tuple.tot_len		= bpf_ntohs(ip6h->payload_len);
		*src			= ip6h->saddr;
		*dst			= ip6h->daddr;
		tuple.sport		= ports->source;
		tuple.dport		= ports->dest;
		break;
	}
	default:
		return XDP_PASS;
	}

	tuplehash = bpf_xdp_flow_lookup(ctx, &tuple, &opts, sizeof(opts));
	if (!tuplehash)
		return XDP_PASS;

	val = bpf_map_lookup_elem(&stats, &key);
	if (val)
		__sync_add_and_fetch(val, 1);

	return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
```
