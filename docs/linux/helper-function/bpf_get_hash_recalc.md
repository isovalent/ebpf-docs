---
title: "Helper Function 'bpf_get_hash_recalc'"
description: "This page documents the 'bpf_get_hash_recalc' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_hash_recalc`

<!-- [FEATURE_TAG](bpf_get_hash_recalc) -->
[:octicons-tag-24: v4.8](https://github.com/torvalds/linux/commit/13c5c240f789bbd2bcacb14a23771491485ae61f)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Retrieve the hash of the packet, _skb_**->hash**. If it is not set, in particular if the hash was cleared due to mangling, recompute this hash. Later accesses to the hash can be done directly with _skb_**->hash**.

Calling **bpf_set_hash_invalid**(), changing a packet prototype with **bpf_skb_change_proto**(), or calling **bpf_skb_store_bytes**() with the **BPF_F_INVALIDATE_HASH** are actions susceptible to clear the hash and to trigger a new computation for the next call to **bpf_get_hash_recalc**().

### Returns

The 32-bit hash.

`#!c static __u32 (* const bpf_get_hash_recalc)(struct __sk_buff *skb) = (void *) 34;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
