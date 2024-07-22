---
title: "Helper Function 'bpf_skb_change_type'"
description: "This page documents the 'bpf_skb_change_type' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_change_type`

<!-- [FEATURE_TAG](bpf_skb_change_type) -->
[:octicons-tag-24: v4.8](https://github.com/torvalds/linux/commit/d2485c4242a826fdf493fd3a27b8b792965b9b9e)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Change the packet type for the packet associated to _skb_. This comes down to setting _skb_**->pkt_type** to _type_, except the eBPF program does not have a write access to _skb_\ **->pkt_type** beside this helper. Using a helper here allows for graceful handling of errors.

The major use case is to change incoming _skb_s to **PACKET_HOST** in a programmatic way instead of having to recirculate via **redirect**(..., **BPF_F_INGRESS**), for example.

Note that _type_ only allows certain values. At this time, they are:

**PACKET_HOST**

&nbsp;&nbsp;&nbsp;&nbsp;Packet is for us.

**PACKET_BROADCAST**

&nbsp;&nbsp;&nbsp;&nbsp;Send packet to all.

**PACKET_MULTICAST**

&nbsp;&nbsp;&nbsp;&nbsp;Send packet to group.

**PACKET_OTHERHOST**

&nbsp;&nbsp;&nbsp;&nbsp;Send packet to someone else.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_skb_change_type)(struct __sk_buff *skb, __u32 type) = (void *) 32;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
