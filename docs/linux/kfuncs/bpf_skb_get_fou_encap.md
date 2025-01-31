---
title: "KFunc 'bpf_skb_get_fou_encap'"
description: "This page documents the 'bpf_skb_get_fou_encap' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_skb_get_fou_encap`

<!-- [FEATURE_TAG](bpf_skb_get_fou_encap) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c50e96099edb134bf107fafc02715fbc4aa2277f)
<!-- [/FEATURE_TAG] -->

Get FOU (Foo over UDP) encapsulation parameters

## Definition

This function allows for reading encapsulation metadata from a packet received on an IPIP device in collect-metadata mode.

**Parameters**

`skb_ctx`:  Pointer to ctx (__sk_buff) in TC program. Cannot be NULL

`encap`:    Pointer to a struct `bpf_fou_encap` storing UDP source and destination port. Cannot be NULL


 **Returns**

`0` on success, a negative error code on failure

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_skb_get_fou_encap(struct __sk_buff *skb_ctx, struct bpf_fou_encap *encap)`
<!-- [/KFUNC_DEF] -->

## Usage

On the ingress path `bpf_skb_get_fou_encap` can be used to read UDP source and destination ports from the receiver's point of view and allows for packet multiplexing across different destination ports within a single BPF program and IPIP device.

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

SEC("tc")
int ipip_gue_set_tunnel(struct __sk_buff *skb)
{
	struct bpf_tunnel_key key = {};
	struct bpf_fou_encap encap = {};
	void *data = (void *)(long)skb->data;
	struct iphdr *iph = data;
	void *data_end = (void *)(long)skb->data_end;
	int ret;

	if (data + sizeof(*iph) > data_end) {
		log_err(1);
		return TC_ACT_SHOT;
	}

	key.tunnel_ttl = 64;
	if (iph->protocol == IPPROTO_ICMP)
		key.remote_ipv4 = 0xac100164; /* 172.16.1.100 */

	ret = bpf_skb_set_tunnel_key(skb, &key, sizeof(key), 0);
	if (ret < 0) {
		log_err(ret);
		return TC_ACT_SHOT;
	}

	encap.sport = 0;
	encap.dport = bpf_htons(5555);

	ret = bpf_skb_set_fou_encap(skb, &encap, FOU_BPF_ENCAP_GUE);
	if (ret < 0) {
		log_err(ret);
		return TC_ACT_SHOT;
	}

	return TC_ACT_OK;
}
```

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2016 VMware
 * Copyright (c) 2016 Facebook
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of version 2 of the GNU General Public
 * License as published by the Free Software Foundation.
 */

SEC("tc")
int ipip_encap_get_tunnel(struct __sk_buff *skb)
{
	int ret;
	struct bpf_tunnel_key key = {};
	struct bpf_fou_encap encap = {};

	ret = bpf_skb_get_tunnel_key(skb, &key, sizeof(key), 0);
	if (ret < 0) {
		log_err(ret);
		return TC_ACT_SHOT;
	}

	ret = bpf_skb_get_fou_encap(skb, &encap);
	if (ret < 0) {
		log_err(ret);
		return TC_ACT_SHOT;
	}

	if (bpf_ntohs(encap.dport) != 5555)
		return TC_ACT_SHOT;

	bpf_printk("%d remote ip 0x%x, sport %d, dport %d\n", ret,
		   key.remote_ipv4, bpf_ntohs(encap.sport),
		   bpf_ntohs(encap.dport));
	return TC_ACT_OK;
}
```
