---
title: "Helper Function 'bpf_l4_csum_replace' - eBPF Docs"
description: "This page documents the 'bpf_l4_csum_replace' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_l4_csum_replace`

<!-- [FEATURE_TAG](bpf_l4_csum_replace) -->
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/91bc4822c3d61b9bb7ef66d3b77948a4f9177954)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Recompute the layer 4 (e.g. TCP, UDP or ICMP) checksum for the packet associated to _skb_. Computation is incremental, so the helper must know the former value of the header field that was modified (_from_), the new value of this field (_to_), and the number of bytes (2 or 4) for this field, stored on the lowest four bits of _flags_. Alternatively, it is possible to store the difference between the previous and the new values of the header field in _to_, by setting _from_ and the four lowest bits of _flags_ to 0. For both methods, _offset_ indicates the location of the IP checksum within the packet. In addition to the size of the field, _flags_ can be added (bitwise OR) actual flags. With **BPF_F_MARK_MANGLED_0**, a null checksum is left untouched (unless **BPF_F_MARK_ENFORCE** is added as well), and for updates resulting in a null checksum the value is set to **CSUM_MANGLED_0** instead. Flag **BPF_F_PSEUDO_HDR** indicates the checksum is to be computed against a pseudo-header.

This helper works in combination with **bpf_csum_diff**(), which does not update the checksum in-place, but offers more flexibility and can handle sizes larger than 2 or 4 for the checksum to update.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_l4_csum_replace)(struct __sk_buff *skb, __u32 offset, __u64 from, __u64 to, __u64 flags) = (void *) 11;`
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
