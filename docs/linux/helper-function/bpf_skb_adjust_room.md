---
title: "Helper Function 'bpf_skb_adjust_room'"
description: "This page documents the 'bpf_skb_adjust_room' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_adjust_room`

<!-- [FEATURE_TAG](bpf_skb_adjust_room) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/2be7e212d5419a400d051c84ca9fdd083e5aacac)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Grow or shrink the room for data in the packet associated to _skb_ by _len_diff_, and according to the selected _mode_.

By default, the helper will reset any offloaded checksum indicator of the skb to CHECKSUM_NONE. This can be avoided by the following flag:

* **BPF_F_ADJ_ROOM_NO_CSUM_RESET**: Do not reset offloaded
  checksum data of the skb to CHECKSUM_NONE.

There are two supported modes at this time:

* **BPF_ADJ_ROOM_MAC**: Adjust room at the mac layer
  (room space is added or removed between the layer 2 and   layer 3 headers).

* **BPF_ADJ_ROOM_NET**: Adjust room at the network layer
  (room space is added or removed between the layer 3 and   layer 4 headers).

The following flags are supported at this time:

* **BPF_F_ADJ_ROOM_FIXED_GSO**: Do not adjust gso_size.
  Adjusting mss in this way is not allowed for datagrams.

* **BPF_F_ADJ_ROOM_ENCAP_L3_IPV4**,
  **BPF_F_ADJ_ROOM_ENCAP_L3_IPV6**:   Any new space is reserved to hold a tunnel header.   Configure skb offsets and other fields accordingly.

* **BPF_F_ADJ_ROOM_ENCAP_L4_GRE**,
  **BPF_F_ADJ_ROOM_ENCAP_L4_UDP**:   Use with ENCAP_L3 flags to further specify the tunnel type.

* **BPF_F_ADJ_ROOM_ENCAP_L2**(_len_):
  Use with ENCAP_L3/L4 flags to further specify the tunnel   type; _len_ is the length of the inner MAC header.

* **BPF_F_ADJ_ROOM_ENCAP_L2_ETH**:
  Use with BPF_F_ADJ_ROOM_ENCAP_L2 flag to further specify the   L2 type as Ethernet.

* **BPF_F_ADJ_ROOM_DECAP_L3_IPV4**,
  **BPF_F_ADJ_ROOM_DECAP_L3_IPV6**:   Indicate the new IP header version after decapsulating the outer   IP header. Used when the inner and outer IP versions are different.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_skb_adjust_room)(struct __sk_buff *skb, __s32 len_diff, __u32 mode, __u64 flags) = (void *) 50;`
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
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
