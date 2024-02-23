---
title: "Helper Function 'bpf_skb_get_xfrm_state' - eBPF Docs"
description: "This page documents the 'bpf_skb_get_xfrm_state' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_get_xfrm_state`

<!-- [FEATURE_TAG](bpf_skb_get_xfrm_state) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/12bed760a78da6e12ac8252fec64d019a9eac523)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Retrieve the XFRM state (IP transform framework, see also **ip-xfrm(8)**) at _index_ in XFRM "security path" for _skb_.

The retrieved value is stored in the **struct bpf_xfrm_state** pointed by _xfrm_state_ and of length _size_.

All values for _flags_ are reserved for future usage, and must be left at zero.

This helper is available only if the kernel was compiled with **CONFIG_XFRM** configuration option.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_skb_get_xfrm_state)(struct __sk_buff *skb, __u32 index, struct bpf_xfrm_state *xfrm_state, __u32 size, __u64 flags) = (void *) 66;`
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
