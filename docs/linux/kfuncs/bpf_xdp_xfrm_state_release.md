---
title: "KFunc 'bpf_xdp_xfrm_state_release'"
description: "This page documents the 'bpf_xdp_xfrm_state_release' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_xdp_xfrm_state_release`

<!-- [FEATURE_TAG](bpf_xdp_xfrm_state_release) -->
[:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/8f0ec8c681755f523cf842bfe350ea40609b83a9)
<!-- [/FEATURE_TAG] -->

## Definition

/* bpf_xdp_xfrm_state_release - Release acquired xfrm_state object
 *
 * This must be invoked for referenced PTR_TO_BTF_ID, and the verifier rejects
 * the program if any references remain in the program in all of the explored
 * states.
 *
 * Parameters:
 * @x		- Pointer to referenced xfrm_state object, obtained using
 *		  bpf_xdp_get_xfrm_state.
 */

<!-- [KFUNC_DEF] -->
`#!c void bpf_xdp_xfrm_state_release(struct xfrm_state *x)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
<!-- [/KFUNC_DEF] -->

## Usage

The intent for this kfunc is to support software RSS (via XDP) for the ongoing/upcoming ipsec pcpu work. [^1]

[^1]: [https://datatracker.ietf.org/doc/draft-ietf-ipsecme-multi-sa-performance/03/](https://datatracker.ietf.org/doc/draft-ietf-ipsecme-multi-sa-performance/03/)

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2016 VMware
 * Copyright (c) 2016 Facebook
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of version 2 of the GNU General Public
 * License as published by the Free Software Foundation.
 */

volatile int xfrm_replay_window = 0;

SEC("xdp")
int xfrm_get_state_xdp(struct xdp_md *xdp)
{
	struct bpf_xfrm_state_opts opts = {};
	struct xfrm_state *x = NULL;
	struct ip_esp_hdr *esph;
	struct bpf_dynptr ptr;
	u8 esph_buf[8] = {};
	u8 iph_buf[20] = {};
	struct iphdr *iph;
	u32 off;

	if (bpf_dynptr_from_xdp(xdp, 0, &ptr))
		goto out;

	off = sizeof(struct ethhdr);
	iph = bpf_dynptr_slice(&ptr, off, iph_buf, sizeof(iph_buf));
	if (!iph || iph->protocol != IPPROTO_ESP)
		goto out;

	off += sizeof(struct iphdr);
	esph = bpf_dynptr_slice(&ptr, off, esph_buf, sizeof(esph_buf));
	if (!esph)
		goto out;

	opts.netns_id = BPF_F_CURRENT_NETNS;
	opts.daddr.a4 = iph->daddr;
	opts.spi = esph->spi;
	opts.proto = IPPROTO_ESP;
	opts.family = AF_INET;

	x = bpf_xdp_get_xfrm_state(xdp, &opts, sizeof(opts));
	if (!x)
		goto out;

	if (!x->replay_esn)
		goto out;

	xfrm_replay_window = x->replay_esn->replay_window;
out:
	if (x)
		bpf_xdp_xfrm_state_release(x);
	return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
```
