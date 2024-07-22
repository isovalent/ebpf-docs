---
title: "Helper Function 'bpf_skb_set_tunnel_opt'"
description: "This page documents the 'bpf_skb_set_tunnel_opt' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_set_tunnel_opt`

<!-- [FEATURE_TAG](bpf_skb_set_tunnel_opt) -->
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/14ca0751c96f8d3d0f52e8ed3b3236f8b34d3460)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Set tunnel options metadata for the packet associated to _skb_ to the option data contained in the raw buffer _opt_ of _size_.

See also the description of the **bpf_skb_get_tunnel_opt**() helper for additional information.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_skb_set_tunnel_opt)(struct __sk_buff *skb, void *opt, __u32 size) = (void *) 30;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
