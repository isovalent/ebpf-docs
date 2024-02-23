---
title: "Helper Function 'bpf_lwt_seg6_action'"
description: "This page documents the 'bpf_lwt_seg6_action' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_lwt_seg6_action`

<!-- [FEATURE_TAG](bpf_lwt_seg6_action) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/fe94cc290f535709d3c5ebd1e472dfd0aec7ee79)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Apply an IPv6 Segment Routing action of type _action_ to the packet associated to _skb_. Each action takes a parameter contained at address _param_, and of length _param_len_ bytes. _action_ can be one of:

**SEG6_LOCAL_ACTION_END_X**

&nbsp;&nbsp;&nbsp;&nbsp;End.X action: Endpoint with Layer-3 cross-connect. Type of _param_: **struct in6_addr**.

**SEG6_LOCAL_ACTION_END_T**

&nbsp;&nbsp;&nbsp;&nbsp;End.T action: Endpoint with specific IPv6 table lookup. Type of _param_: **int**.

**SEG6_LOCAL_ACTION_END_B6**

&nbsp;&nbsp;&nbsp;&nbsp;End.B6 action: Endpoint bound to an SRv6 policy. Type of _param_: **struct ipv6_sr_hdr**.

**SEG6_LOCAL_ACTION_END_B6_ENCAP**

&nbsp;&nbsp;&nbsp;&nbsp;End.B6.Encap action: Endpoint bound to an SRv6 encapsulation policy. Type of _param_: **struct ipv6_sr_hdr**.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_lwt_seg6_action)(struct __sk_buff *skb, __u32 action, void *param, __u32 param_len) = (void *) 76;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
