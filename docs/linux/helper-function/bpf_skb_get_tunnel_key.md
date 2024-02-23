---
title: "Helper Function 'bpf_skb_get_tunnel_key' - eBPF Docs"
description: "This page documents the 'bpf_skb_get_tunnel_key' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_get_tunnel_key`

<!-- [FEATURE_TAG](bpf_skb_get_tunnel_key) -->
[:octicons-tag-24: v4.3](https://github.com/torvalds/linux/commit/d3aa45ce6b94c65b83971257317867db13e5f492)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get tunnel metadata. This helper takes a pointer _key_ to an empty **struct bpf_tunnel_key** of **size**, that will be filled with tunnel metadata for the packet associated to _skb_. The _flags_ can be set to **BPF_F_TUNINFO_IPV6**, which indicates that the tunnel is based on IPv6 protocol instead of IPv4.

The **struct bpf_tunnel_key** is an object that generalizes the principal parameters used by various tunneling protocols into a single struct. This way, it can be used to easily make a decision based on the contents of the encapsulation header, "summarized" in this struct. In particular, it holds the IP address of the remote end (IPv4 or IPv6, depending on the case) in _key_**->remote_ipv4** or _key_**->remote_ipv6**. Also, this struct exposes the _key_**->tunnel_id**, which is generally mapped to a VNI (Virtual Network Identifier), making it programmable together with the **bpf_skb_set_tunnel_key**\ () helper.

Let's imagine that the following code is part of a program attached to the TC ingress interface, on one end of a GRE tunnel, and is supposed to filter out all messages coming from remote ends with IPv4 address other than 10.0.0.1:

```
int ret; struct bpf_tunnel_key key = {};
```

&nbsp;&nbsp;&nbsp;&nbsp;ret = bpf_skb_get_tunnel_key(skb, &key, sizeof(key), 0); if (ret < 0)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;return TC_ACT_SHOT;// drop packet

&nbsp;&nbsp;&nbsp;&nbsp;if (key.remote_ipv4 != 0x0a000001)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;return TC_ACT_SHOT;// drop packet

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;return TC_ACT_OK;// accept packet

This interface can also be used with all encapsulation devices that can operate in "collect metadata" mode: instead of having one network device per specific configuration, the "collect metadata" mode only requires a single device where the configuration can be extracted from this helper.

This can be used together with various tunnels such as VXLan, Geneve, GRE or IP in IP (IPIP).

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_skb_get_tunnel_key)(struct __sk_buff *skb, struct bpf_tunnel_key *key, __u32 size, __u64 flags) = (void *) 20;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
