# KFunc `bpf_skb_set_fou_encap`

<!-- [FEATURE_TAG](bpf_skb_set_fou_encap) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/c50e96099edb134bf107fafc02715fbc4aa2277f)
<!-- [/FEATURE_TAG] -->

Set FOU (Foo Over UDP) encap parameters

## Definition

This function allows for using GUE or FOU encapsulation together with an ipip device in collect-metadata mode.

It is meant to be used in BPF tc-hooks and after a call to the [`bpf_skb_set_tunnel_key`](../helper-function/bpf_skb_set_tunnel_key.md) helper, responsible for setting IP addresses.

**Parameters**
`skb_ctx`:  Pointer to ctx (__sk_buff) in TC program. Cannot be NULL

`encap`:    Pointer to a `struct bpf_fou_encap` storing UDP src and dst ports. If sport is set to 0 the kernel will auto-assign a port. This is similar to using `encap-sport auto`. Cannot be NULL

`type`: Encapsulation type for the packet. Their definitions are specified in `enum bpf_fou_encap_type`, possible values:

- `FOU_BPF_ENCAP_FOU`
- `FOU_BPF_ENCAP_GUE`


**Returns**

`0` on success, a negative error code on failure

<!-- [KFUNC_DEF] -->
`#!c int bpf_skb_set_fou_encap(struct __sk_buff *skb_ctx, struct bpf_fou_encap *encap, int type)`
<!-- [/KFUNC_DEF] -->

## Usage

The bpf_skb_set_fou_encap kfunc is supposed to be used in tandem and after a successful call to the [`bpf_skb_set_tunnel_key`](../helper-function/bpf_skb_set_tunnel_key.md) bpf-helper. UDP source and destination ports can be controlled by passing a `struct bpf_fou_encap`. A source port of zero will auto-assign a source port. `enum bpf_fou_encap_type` is used to specify if the egress path should FOU or GUE encap the packet.

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

SEC("tc")
int ipip_fou_set_tunnel(struct __sk_buff *skb)
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

	ret = bpf_skb_set_fou_encap(skb, &encap, FOU_BPF_ENCAP_FOU);
	if (ret < 0) {
		log_err(ret);
		return TC_ACT_SHOT;
	}

	return TC_ACT_OK;
}
```
