---
title: "KFunc 'bpf_xdp_get_xfrm_state'"
description: "This page documents the 'bpf_xdp_get_xfrm_state' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_xdp_get_xfrm_state`

<!-- [FEATURE_TAG](bpf_xdp_get_xfrm_state) -->
[:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/8f0ec8c681755f523cf842bfe350ea40609b83a9)
<!-- [/FEATURE_TAG] -->

Get XFRM state

## Definition

A `struct xfrm_state *`, if found, must be released with a corresponding bpf_xdp_xfrm_state_release.

**Parameters**

`ctx` - Pointer to ctx (xdp_md) in XDP program. Cannot be NULL

`opts` - Options for lookup. Cannot be NULL

**Members**

`opts.error` - Out parameter, set for any errors encountered, Values:

- `-EINVAL` - netns_id is less than -1
- `-EINVAL` - opts__sz isn't `BPF_XFRM_STATE_OPTS_SZ`
- `-ENONET` - No network namespace found for netns_id
- `-ENOENT` - No xfrm_state found

`opts.netns_id` - Specify the network namespace for lookup, Values:

- `BPF_F_CURRENT_NETNS` - (-1) Use namespace associated with ctx
- `[0, S32_MAX]` - Network Namespace ID

`opts.mark` - XFRM mark to match on

`opts.daddr` - Destination address to match on

`opts.spi` - Security parameter index to match on

`opts.proto` - IP protocol to match on (eg. `IPPROTO_ESP`)

`opts.family` - Protocol family to match on (`AF_INET`/`AF_INET6`)

`opts__sz` - Length of the bpf_xfrm_state_opts structure. Must be `BPF_XFRM_STATE_OPTS_SZ`

<!-- [KFUNC_DEF] -->
`#!c struct xfrm_state *bpf_xdp_get_xfrm_state(struct xdp_md *ctx, struct bpf_xfrm_state_opts *opts, u32 opts__sz)`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc helper accesses internal `xfrm_state` associated with an SA(Security Association). This is intended to be used for the assigning of special per-CPU <nospell>SAs</nospell> to a particular CPU. In other words: for custom software RSS. [^1]

[^1]: [https://datatracker.ietf.org/doc/draft-ietf-ipsecme-multi-sa-performance/03/](https://datatracker.ietf.org/doc/draft-ietf-ipsecme-multi-sa-performance/03/)

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
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
