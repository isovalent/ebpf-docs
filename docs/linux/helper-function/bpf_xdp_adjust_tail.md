---
title: "Helper Function 'bpf_xdp_adjust_tail'"
description: "This page documents the 'bpf_xdp_adjust_tail' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_adjust_tail`

<!-- [FEATURE_TAG](bpf_xdp_adjust_tail) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/b32cc5b9a346319c171e3ad905e0cddda032b5eb)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


Adjust (move) `xdp_md->data_end` by `delta` bytes. It is possible to both shrink and grow the packet tail. Shrink done via `delta` being a negative integer.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

This helper invokes `#!c memset` under the hood, if `delta` is strictly positive, to explicitly clear the newly added memory area. Thus, it impacts the speed of packet processing in this particular case.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_xdp_adjust_tail)(struct xdp_md *xdp_md, int delta) = (void *) 65;`

## Usage

This helper should be called when the same `xdp_md` buffer is reused to build and send back a "reply" packet. The "reply" packet, due to some protocol standards, may be larger or smaller than the initial request.

In another use case, we need to add some specific trailer at the end of the payload. Then this trailer is usually parsed in some [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) program, which is working in cooperation.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
/* Following examples could be loaded and attached with these commands:
 * bpftool --debug prog loadall test.bpf.o /sys/fs/bpf/prog
 * bpftool net attach xdpdrv name xdp_test_adjust_tail dev <dev_name>
 * tc filter add dev <dev_name> ingress bpf object-pinned \
 * 	/sys/fs/bpf/prog/tc_test_adjust_tail
 */

#include "vmlinux.h"

#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>

#include "bpf_tracing_net.h"

#define TARGET_PORT 666

u32 trailer = 0xffffff00;

SEC("tc")
int tc_test_adjust_tail(struct __sk_buff *skb)
{
	void *data;
	void *data_end;

	if (bpf_skb_pull_data(skb, skb->len) < 0) {
		return TC_ACT_SHOT;
	}

	data = (void *)(unsigned long)skb->data;
	data_end = (void *)(unsigned long)skb->data_end;

	/* packet parsing to check if this is a valid UDP packet for the
	 * TARGET_PORT is omitted for brevity
	 */

	/* get trailer from packet destined to TARGET_PORT */
	u16 len = data_end - data;
	if ((data + len) > data_end)
		return TC_ACT_SHOT;

	/*  0x3fff - largest value for a packet size that allows 9K jumbo */
	u16 offset = (len - sizeof(trailer)) & 0x3fff;

	u32 *trailer_start = (u32 *)(data + offset);
	if ((trailer_start + 1) > (u32 *)data_end)
		return TC_ACT_SHOT;

	if (*trailer_start & 0xffffff00)
		bpf_printk("Packet was received on XDP RX queue %d\n",
			   (*trailer_start) & 0x000000ff);

	return TC_ACT_OK;
}

SEC("xdp")
int xdp_test_adjust_tail(struct xdp_md *xdp)
{
	void *data, *data_end;
	u16 init_pkt_size;
	struct ethhdr *eth;
	struct iphdr *ip = NULL;
	struct udphdr *udp;
	u16 eth_proto;

	data = (void *)(unsigned long)xdp->data;
	data_end = (void *)(unsigned long)xdp->data_end;
	eth = (struct ethhdr *)data;
	init_pkt_size = data_end - data;

	if (eth + 1 > data_end)
		return XDP_DROP;

	eth_proto = bpf_ntohs(eth->h_proto);
	if (eth_proto == ETH_P_IP) {
		ip = (struct iphdr *)(eth + 1);
	} else
		return XDP_PASS;

	if (((ip + 1) > data_end) || (ip->ihl < 5) || (ip->ihl > 15)) {
		return XDP_DROP;
	}

	if (ip->protocol == IPPROTO_UDP) {
		udp = (struct udphdr *)(ip + 1);
		if (udp + 1 > data_end)
			return XDP_DROP;

		if (udp->len < sizeof(*udp))
			return XDP_DROP;

		if (bpf_ntohs(udp->dest) == TARGET_PORT) {
			u16 len, offset;
			u32 *trailer_start;

			/* 4-byte pattern to add at the end of a packet, that is
			 * destined to TARGET_PORT
			 */
			trailer |= xdp->rx_queue_index;

			if (bpf_xdp_adjust_tail(xdp, init_pkt_size + sizeof(trailer)) < 0) {
				return XDP_ABORTED;
			}
			data = (void *)(unsigned long)xdp->data;
			data_end = (void *)(unsigned long)xdp->data_end;

			/* write trailer */
			len = data_end - data;
			if ((data + len) > data_end) {
				return XDP_ABORTED;
			}

			/*  0x3fff - largest value for a packet size that allows 9K jumbo */
			offset = (len - sizeof(trailer)) & 0x3fff;

			trailer_start = (u32 *)(data + offset);
			if ((trailer_start + 1) > (u32 *)data_end) {
				return XDP_ABORTED;
			}

			*trailer_start = trailer;

			return XDP_PASS;
		}
	}
	return XDP_PASS;
}

char _license[] SEC("license") = "GPL";

```
