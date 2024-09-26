---
title: "KFunc 'bpf_dynptr_from_xdp'"
description: "This page documents the 'bpf_dynptr_from_xdp' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_dynptr_from_xdp`

<!-- [FEATURE_TAG](bpf_dynptr_from_xdp) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/05421aecd4ed65da0dc17b0c3c13779ef334e9e5)
<!-- [/FEATURE_TAG] -->

Get dynptrs whose underlying pointer points to a xdp_buff.

## Definition

For reads and writes on the dynptr, this includes reading/writing from/to and across fragments. Data slices through the [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md) API are not supported; instead [`bpf_dynptr_slice()`](bpf_dynptr_slice.md) and [`bpf_dynptr_slice_rdwr()`](bpf_dynptr_slice_rdwr.md) should be used.

<!-- [KFUNC_DEF] -->
`#!c int bpf_dynptr_from_xdp(struct xdp_md *x, u64 flags, struct bpf_dynptr *ptr__uninit)`
<!-- [/KFUNC_DEF] -->

## Usage

The dynptr acts on xdp data. xdp dynptrs have two main benefits. One is that they allow operations on sizes that are not statically known at compile-time (for example variable-sized accesses). Another is that parsing the packet data through dynptrs (instead of through direct access of xdp->data and xdp->data_end) can be more ergonomic and less brittle (for example does not need manual if checking for being within bounds of data_end).

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

??? example "Parse TCP header option"
    ```c
    // SPDX-License-Identifier: GPL-2.0

    /* This logic is lifted from a real-world use case of packet parsing, used in
     * the open source library katran, a layer 4 load balancer.
     *
     * This test demonstrates how to parse packet contents using dynptrs. The
     * original code (parsing without dynptrs) can be found in test_parse_tcp_hdr_opt.c
     */

    #include <linux/bpf.h>
    #include <bpf/bpf_helpers.h>
    #include <linux/tcp.h>
    #include <stdbool.h>
    #include <linux/ipv6.h>
    #include <linux/if_ether.h>
    #include "test_tcp_hdr_options.h"
    #include "bpf_kfuncs.h"

    char _license[] SEC("license") = "GPL";

    /* Kind number used for experiments */
    const __u32 tcp_hdr_opt_kind_tpr = 0xFD;
    /* Length of the tcp header option */
    const __u32 tcp_hdr_opt_len_tpr = 6;
    /* maximum number of header options to check to lookup server_id */
    const __u32 tcp_hdr_opt_max_opt_checks = 15;

    __u32 server_id;

    static int parse_hdr_opt(struct bpf_dynptr *ptr, __u32 *off, __u8 *hdr_bytes_remaining,
    			 __u32 *server_id)
    {
    	__u8 *tcp_opt, kind, hdr_len;
    	__u8 buffer[sizeof(kind) + sizeof(hdr_len) + sizeof(*server_id)];
    	__u8 *data;

    	__builtin_memset(buffer, 0, sizeof(buffer));

    	data = bpf_dynptr_slice(ptr, *off, buffer, sizeof(buffer));
    	if (!data)
    		return -1;

    	kind = data[0];

    	if (kind == TCPOPT_EOL)
    		return -1;

    	if (kind == TCPOPT_NOP) {
    		*off += 1;
    		*hdr_bytes_remaining -= 1;
    		return 0;
    	}

    	if (*hdr_bytes_remaining < 2)
    		return -1;

    	hdr_len = data[1];
    	if (hdr_len > *hdr_bytes_remaining)
    		return -1;

    	if (kind == tcp_hdr_opt_kind_tpr) {
    		if (hdr_len != tcp_hdr_opt_len_tpr)
    			return -1;

    		__builtin_memcpy(server_id, (__u32 *)(data + 2), sizeof(*server_id));
    		return 1;
    	}

    	*off += hdr_len;
    	*hdr_bytes_remaining -= hdr_len;
    	return 0;
    }

    SEC("xdp")
    int xdp_ingress_v6(struct xdp_md *xdp)
    {
    	__u8 buffer[sizeof(struct tcphdr)] = {};
    	__u8 hdr_bytes_remaining;
    	struct tcphdr *tcp_hdr;
    	__u8 tcp_hdr_opt_len;
    	int err = 0;
    	__u32 off;

    	struct bpf_dynptr ptr;

    	bpf_dynptr_from_xdp(xdp, 0, &ptr);

    	off = sizeof(struct ethhdr) + sizeof(struct ipv6hdr);

    	tcp_hdr = bpf_dynptr_slice(&ptr, off, buffer, sizeof(buffer));
    	if (!tcp_hdr)
    		return XDP_DROP;

    	tcp_hdr_opt_len = (tcp_hdr->doff * 4) - sizeof(struct tcphdr);
    	if (tcp_hdr_opt_len < tcp_hdr_opt_len_tpr)
    		return XDP_DROP;

    	hdr_bytes_remaining = tcp_hdr_opt_len;

    	off += sizeof(struct tcphdr);

    	/* max number of bytes of options in tcp header is 40 bytes */
    	for (int i = 0; i < tcp_hdr_opt_max_opt_checks; i++) {
    		err = parse_hdr_opt(&ptr, &off, &hdr_bytes_remaining, &server_id);

    		if (err || !hdr_bytes_remaining)
    			break;
    	}

    	if (!server_id)
    		return XDP_DROP;

    	return XDP_PASS;
    }
    ```

??? example "Tunnel encapsulate with dynamic point"
    ```c
    // SPDX-License-Identifier: GPL-2.0
    /* Copyright (c) 2022 Meta */
    #include <stddef.h>
    #include <string.h>
    #include <linux/bpf.h>
    #include <linux/if_ether.h>
    #include <linux/if_packet.h>
    #include <linux/ip.h>
    #include <linux/ipv6.h>
    #include <linux/in.h>
    #include <linux/udp.h>
    #include <linux/tcp.h>
    #include <linux/pkt_cls.h>
    #include <sys/socket.h>
    #include <bpf/bpf_helpers.h>
    #include <bpf/bpf_endian.h>
    #include "test_iptunnel_common.h"
    #include "bpf_kfuncs.h"

    const size_t tcphdr_sz = sizeof(struct tcphdr);
    const size_t udphdr_sz = sizeof(struct udphdr);
    const size_t ethhdr_sz = sizeof(struct ethhdr);
    const size_t iphdr_sz = sizeof(struct iphdr);
    const size_t ipv6hdr_sz = sizeof(struct ipv6hdr);

    struct {
        __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
        __uint(max_entries, 256);
        __type(key, __u32);
        __type(value, __u64);
    } rxcnt SEC(".maps");

    struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __uint(max_entries, MAX_IPTNL_ENTRIES);
        __type(key, struct vip);
        __type(value, struct iptnl_info);
    } vip2tnl SEC(".maps");

    static __always_inline void count_tx(__u32 protocol)
    {
        __u64 *rxcnt_count;

        rxcnt_count = bpf_map_lookup_elem(&rxcnt, &protocol);
        if (rxcnt_count)
            *rxcnt_count += 1;
    }

    static __always_inline int get_dport(void *trans_data, __u8 protocol)
    {
        struct tcphdr *th;
        struct udphdr *uh;

        switch (protocol) {
        case IPPROTO_TCP:
            th = (struct tcphdr *)trans_data;
            return th->dest;
        case IPPROTO_UDP:
            uh = (struct udphdr *)trans_data;
            return uh->dest;
        default:
            return 0;
        }
    }

    static __always_inline void set_ethhdr(struct ethhdr *new_eth,
                        const struct ethhdr *old_eth,
                        const struct iptnl_info *tnl,
                        __be16 h_proto)
    {
        memcpy(new_eth->h_source, old_eth->h_dest, sizeof(new_eth->h_source));
        memcpy(new_eth->h_dest, tnl->dmac, sizeof(new_eth->h_dest));
        new_eth->h_proto = h_proto;
    }

    static __always_inline int handle_ipv4(struct xdp_md *xdp, struct bpf_dynptr *xdp_ptr)
    {
        __u8 eth_buffer[ethhdr_sz + iphdr_sz + ethhdr_sz];
        __u8 iph_buffer_tcp[iphdr_sz + tcphdr_sz];
        __u8 iph_buffer_udp[iphdr_sz + udphdr_sz];
        struct bpf_dynptr new_xdp_ptr;
        struct iptnl_info *tnl;
        struct ethhdr *new_eth;
        struct ethhdr *old_eth;
        __u32 transport_hdr_sz;
        struct iphdr *iph;
        __u16 *next_iph;
        __u16 payload_len;
        struct vip vip = {};
        int dport;
        __u32 csum = 0;
        int i;

        __builtin_memset(eth_buffer, 0, sizeof(eth_buffer));
        __builtin_memset(iph_buffer_tcp, 0, sizeof(iph_buffer_tcp));
        __builtin_memset(iph_buffer_udp, 0, sizeof(iph_buffer_udp));

        if (ethhdr_sz + iphdr_sz + tcphdr_sz > xdp->data_end - xdp->data)
            iph = bpf_dynptr_slice(xdp_ptr, ethhdr_sz, iph_buffer_udp, sizeof(iph_buffer_udp));
        else
            iph = bpf_dynptr_slice(xdp_ptr, ethhdr_sz, iph_buffer_tcp, sizeof(iph_buffer_tcp));

        if (!iph)
            return XDP_DROP;

        dport = get_dport(iph + 1, iph->protocol);
        if (dport == -1)
            return XDP_DROP;

        vip.protocol = iph->protocol;
        vip.family = AF_INET;
        vip.daddr.v4 = iph->daddr;
        vip.dport = dport;
        payload_len = bpf_ntohs(iph->tot_len);

        tnl = bpf_map_lookup_elem(&vip2tnl, &vip);
        /* It only does v4-in-v4 */
        if (!tnl || tnl->family != AF_INET)
            return XDP_PASS;

        if (bpf_xdp_adjust_head(xdp, 0 - (int)iphdr_sz))
            return XDP_DROP;

        bpf_dynptr_from_xdp(xdp, 0, &new_xdp_ptr);
        new_eth = bpf_dynptr_slice_rdwr(&new_xdp_ptr, 0, eth_buffer, sizeof(eth_buffer));
        if (!new_eth)
            return XDP_DROP;

        iph = (struct iphdr *)(new_eth + 1);
        old_eth = (struct ethhdr *)(iph + 1);

        set_ethhdr(new_eth, old_eth, tnl, bpf_htons(ETH_P_IP));

        if (new_eth == eth_buffer)
            bpf_dynptr_write(&new_xdp_ptr, 0, eth_buffer, sizeof(eth_buffer), 0);

        iph->version = 4;
        iph->ihl = iphdr_sz >> 2;
        iph->frag_off =	0;
        iph->protocol = IPPROTO_IPIP;
        iph->check = 0;
        iph->tos = 0;
        iph->tot_len = bpf_htons(payload_len + iphdr_sz);
        iph->daddr = tnl->daddr.v4;
        iph->saddr = tnl->saddr.v4;
        iph->ttl = 8;

        next_iph = (__u16 *)iph;
        for (i = 0; i < iphdr_sz >> 1; i++)
            csum += *next_iph++;

        iph->check = ~((csum & 0xffff) + (csum >> 16));

        count_tx(vip.protocol);

        return XDP_TX;
    }

    static __always_inline int handle_ipv6(struct xdp_md *xdp, struct bpf_dynptr *xdp_ptr)
    {
        __u8 eth_buffer[ethhdr_sz + ipv6hdr_sz + ethhdr_sz];
        __u8 ip6h_buffer_tcp[ipv6hdr_sz + tcphdr_sz];
        __u8 ip6h_buffer_udp[ipv6hdr_sz + udphdr_sz];
        struct bpf_dynptr new_xdp_ptr;
        struct iptnl_info *tnl;
        struct ethhdr *new_eth;
        struct ethhdr *old_eth;
        __u32 transport_hdr_sz;
        struct ipv6hdr *ip6h;
        __u16 payload_len;
        struct vip vip = {};
        int dport;

        __builtin_memset(eth_buffer, 0, sizeof(eth_buffer));
        __builtin_memset(ip6h_buffer_tcp, 0, sizeof(ip6h_buffer_tcp));
        __builtin_memset(ip6h_buffer_udp, 0, sizeof(ip6h_buffer_udp));

        if (ethhdr_sz + iphdr_sz + tcphdr_sz > xdp->data_end - xdp->data)
            ip6h = bpf_dynptr_slice(xdp_ptr, ethhdr_sz, ip6h_buffer_udp, sizeof(ip6h_buffer_udp));
        else
            ip6h = bpf_dynptr_slice(xdp_ptr, ethhdr_sz, ip6h_buffer_tcp, sizeof(ip6h_buffer_tcp));

        if (!ip6h)
            return XDP_DROP;

        dport = get_dport(ip6h + 1, ip6h->nexthdr);
        if (dport == -1)
            return XDP_DROP;

        vip.protocol = ip6h->nexthdr;
        vip.family = AF_INET6;
        memcpy(vip.daddr.v6, ip6h->daddr.s6_addr32, sizeof(vip.daddr));
        vip.dport = dport;
        payload_len = ip6h->payload_len;

        tnl = bpf_map_lookup_elem(&vip2tnl, &vip);
        /* It only does v6-in-v6 */
        if (!tnl || tnl->family != AF_INET6)
            return XDP_PASS;

        if (bpf_xdp_adjust_head(xdp, 0 - (int)ipv6hdr_sz))
            return XDP_DROP;

        bpf_dynptr_from_xdp(xdp, 0, &new_xdp_ptr);
        new_eth = bpf_dynptr_slice_rdwr(&new_xdp_ptr, 0, eth_buffer, sizeof(eth_buffer));
        if (!new_eth)
            return XDP_DROP;

        ip6h = (struct ipv6hdr *)(new_eth + 1);
        old_eth = (struct ethhdr *)(ip6h + 1);

        set_ethhdr(new_eth, old_eth, tnl, bpf_htons(ETH_P_IPV6));

        if (new_eth == eth_buffer)
            bpf_dynptr_write(&new_xdp_ptr, 0, eth_buffer, sizeof(eth_buffer), 0);

        ip6h->version = 6;
        ip6h->priority = 0;
        memset(ip6h->flow_lbl, 0, sizeof(ip6h->flow_lbl));
        ip6h->payload_len = bpf_htons(bpf_ntohs(payload_len) + ipv6hdr_sz);
        ip6h->nexthdr = IPPROTO_IPV6;
        ip6h->hop_limit = 8;
        memcpy(ip6h->saddr.s6_addr32, tnl->saddr.v6, sizeof(tnl->saddr.v6));
        memcpy(ip6h->daddr.s6_addr32, tnl->daddr.v6, sizeof(tnl->daddr.v6));

        count_tx(vip.protocol);

        return XDP_TX;
    }

    SEC("xdp")
    int _xdp_tx_iptunnel(struct xdp_md *xdp)
    {
        __u8 buffer[ethhdr_sz];
        struct bpf_dynptr ptr;
        struct ethhdr *eth;
        __u16 h_proto;

        __builtin_memset(buffer, 0, sizeof(buffer));

        bpf_dynptr_from_xdp(xdp, 0, &ptr);
        eth = bpf_dynptr_slice(&ptr, 0, buffer, sizeof(buffer));
        if (!eth)
            return XDP_DROP;

        h_proto = eth->h_proto;

        if (h_proto == bpf_htons(ETH_P_IP))
            return handle_ipv4(xdp, &ptr);
        else if (h_proto == bpf_htons(ETH_P_IPV6))

            return handle_ipv6(xdp, &ptr);
        else
            return XDP_DROP;
    }

    char _license[] SEC("license") = "GPL";
    ```
