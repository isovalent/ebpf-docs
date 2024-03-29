---
title: "Helper Function 'bpf_l3_csum_replace'"
description: "This page documents the 'bpf_l3_csum_replace' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_l3_csum_replace`

<!-- [FEATURE_TAG](bpf_l3_csum_replace) -->
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/91bc4822c3d61b9bb7ef66d3b77948a4f9177954)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Recompute the layer 3 (e.g. IP) checksum for the packet associated to _skb_. Computation is incremental, so the helper must know the former value of the header field that was modified (_from_), the new value of this field (_to_), and the number of bytes (2 or 4) for this field, stored in _size_. Alternatively, it is possible to store the difference between the previous and the new values of the header field in _to_, by setting _from_ and _size_ to 0. For both methods, _offset_ indicates the location of the IP checksum within the packet.

This helper works in combination with **bpf_csum_diff**(), which does not update the checksum in-place, but offers more flexibility and can handle sizes larger than 2 or 4 for the checksum to update.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_l3_csum_replace)(struct __sk_buff *skb, __u32 offset, __u64 from, __u64 to, __u64 size) = (void *) 10;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
