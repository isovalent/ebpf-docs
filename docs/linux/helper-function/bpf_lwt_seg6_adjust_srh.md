---
title: "Helper Function 'bpf_lwt_seg6_adjust_srh'"
description: "This page documents the 'bpf_lwt_seg6_adjust_srh' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_lwt_seg6_adjust_srh`

<!-- [FEATURE_TAG](bpf_lwt_seg6_adjust_srh) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/fe94cc290f535709d3c5ebd1e472dfd0aec7ee79)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Adjust the size allocated to TLVs in the outermost IPv6 Segment Routing Header contained in the packet associated to _skb_, at position _offset_ by _delta_ bytes. Only offsets after the segments are accepted. _delta_ can be as well positive (growing) as negative (shrinking).

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_lwt_seg6_adjust_srh)(struct __sk_buff *skb, __u32 offset, __s32 delta) = (void *) 75;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
