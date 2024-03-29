---
title: "Helper Function 'bpf_skb_change_proto'"
description: "This page documents the 'bpf_skb_change_proto' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_change_proto`

<!-- [FEATURE_TAG](bpf_skb_change_proto) -->
[:octicons-tag-24: v4.8](https://github.com/torvalds/linux/commit/6578171a7ff0c31dc73258f93da7407510abf085)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Change the protocol of the _skb_ to _proto_. Currently supported are transition from IPv4 to IPv6, and from IPv6 to IPv4. The helper takes care of the groundwork for the transition, including resizing the socket buffer. The eBPF program is expected to fill the new headers, if any, via **skb_store_bytes**() and to recompute the checksums with **bpf_l3_csum_replace**() and **bpf_l4_csum_replace**\ (). The main case for this helper is to perform NAT64 operations out of an eBPF program.

Internally, the GSO type is marked as dodgy so that headers are checked and segments are recalculated by the GSO/GRO engine. The size for GSO target is adapted as well.

All values for _flags_ are reserved for future usage, and must be left at zero.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_skb_change_proto)(struct __sk_buff *skb, __be16 proto, __u64 flags) = (void *) 31;`
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
