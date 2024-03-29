---
title: "Helper Function 'bpf_skb_change_head'"
description: "This page documents the 'bpf_skb_change_head' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_change_head`

<!-- [FEATURE_TAG](bpf_skb_change_head) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/3a0af8fd61f90920f6fa04e4f1e9a6a73c1b4fd2)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Grows headroom of packet associated to _skb_ and adjusts the offset of the MAC header accordingly, adding _len_ bytes of space. It automatically extends and reallocates memory as required.

This helper can be used on a layer 3 _skb_ to push a MAC header for redirection into a layer 2 device.

All values for _flags_ are reserved for future usage, and must be left at zero.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_skb_change_head)(struct __sk_buff *skb, __u32 len, __u64 flags) = (void *) 43;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md) [:octicons-tag-24: v5.8](6f3f65d80dac8f2bafce2213005821fccdce194c)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) [:octicons-tag-24: v5.8](6f3f65d80dac8f2bafce2213005821fccdce194c)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
