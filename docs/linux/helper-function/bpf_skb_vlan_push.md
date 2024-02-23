---
title: "Helper Function 'bpf_skb_vlan_push' - eBPF Docs"
description: "This page documents the 'bpf_skb_vlan_push' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_vlan_push`

<!-- [FEATURE_TAG](bpf_skb_vlan_push) -->
[:octicons-tag-24: v4.3](https://github.com/torvalds/linux/commit/4e10df9a60d96ced321dd2af71da558c6b750078)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Push a _vlan_tci_ (VLAN tag control information) of protocol _vlan_proto_ to the packet associated to _skb_, then update the checksum. Note that if _vlan_proto_ is different from **ETH_P_8021Q** and **ETH_P_8021AD**, it is considered to be **ETH_P_8021Q**.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_skb_vlan_push)(struct __sk_buff *skb, __be16 vlan_proto, __u16 vlan_tci) = (void *) 18;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
