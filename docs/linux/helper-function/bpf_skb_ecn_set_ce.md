---
title: "Helper Function 'bpf_skb_ecn_set_ce' - eBPF Docs"
description: "This page documents the 'bpf_skb_ecn_set_ce' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_ecn_set_ce`

<!-- [FEATURE_TAG](bpf_skb_ecn_set_ce) -->
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/f7c917ba11a67632a8452ea99fe132f626a7a2cc)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Set ECN (Explicit Congestion Notification) field of IP header to **CE** (Congestion Encountered) if current value is **ECT** (ECN Capable Transport). Otherwise, do nothing. Works with IPv6 and IPv4.

### Returns

1 if the **CE** flag is set (either by the current helper call or because it was already present), 0 if it is not set.

`#!c static long (*bpf_skb_ecn_set_ce)(struct __sk_buff *skb) = (void *) 97;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
