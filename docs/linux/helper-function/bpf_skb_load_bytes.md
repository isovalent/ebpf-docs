---
title: "Helper Function 'bpf_skb_load_bytes'"
description: "This page documents the 'bpf_skb_load_bytes' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_load_bytes`

<!-- [FEATURE_TAG](bpf_skb_load_bytes) -->
[:octicons-tag-24: v4.5](https://github.com/torvalds/linux/commit/05c74e5e53f6cb07502c3e6a820f33e2777b6605)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
This helper was provided as an easy way to load data from a packet. It can be used to load _len_ bytes from _offset_ from the packet associated to _skb_, into the buffer pointed by _to_.

Since Linux 4.7, usage of this helper has mostly been replaced by "direct packet access", enabling packet data to be manipulated with _skb_**->data** and _skb_**->data_end** pointing respectively to the first byte of packet data and to the byte after the last byte of packet data. However, it remains useful if one wishes to read large quantities of data at once from a packet into the eBPF stack.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_skb_load_bytes)(const void *skb, __u32 offset, void *to, __u32 len) = (void *) 26;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
