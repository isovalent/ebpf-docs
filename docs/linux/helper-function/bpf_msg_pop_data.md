---
title: "Helper Function 'bpf_msg_pop_data'"
description: "This page documents the 'bpf_msg_pop_data' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_msg_pop_data`

<!-- [FEATURE_TAG](bpf_msg_pop_data) -->
[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/7246d8ed4dcce23f7509949a77be15fa9f0e3d28)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Will remove _len_ bytes from a _msg_ starting at byte _start_. This may result in **ENOMEM** errors under certain situations if an allocation and copy are required due to a full ring buffer. However, the helper will try to avoid doing the allocation if possible. Other errors can occur if input parameters are invalid either due to _start_ byte not being valid part of _msg_ payload and/or _pop_ value being to large.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_msg_pop_data)(struct sk_msg_md *msg, __u32 start, __u32 len, __u64 flags) = (void *) 91;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
