---
title: "Helper Function 'bpf_redirect_peer'"
description: "This page documents the 'bpf_redirect_peer' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_redirect_peer`

<!-- [FEATURE_TAG](bpf_redirect_peer) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/9aa1206e8f48222f35a0c809f33b2f4aaa1e2661)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Redirect the packet to another net device of index _ifindex_. This helper is somewhat similar to **bpf_redirect**(), except that the redirection happens to the _ifindex_' peer device and the netns switch takes place from ingress to ingress without going through the CPU's backlog queue.

The _flags_ argument is reserved and must be 0. The helper is currently only supported for tc BPF program types at the ingress hook and for veth and netkit target device types. The peer device must reside in a different network namespace.

### Returns

The helper returns **TC_ACT_REDIRECT** on success or **TC_ACT_SHOT** on error.

`#!c static long (*bpf_redirect_peer)(__u32 ifindex, __u64 flags) = (void *) 155;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
