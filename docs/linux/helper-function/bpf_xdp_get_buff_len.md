---
title: "Helper Function 'bpf_xdp_get_buff_len' - eBPF Docs"
description: "This page documents the 'bpf_xdp_get_buff_len' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_get_buff_len`

<!-- [FEATURE_TAG](bpf_xdp_get_buff_len) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0165cc817075cf701e4289838f1d925ff1911b3e)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get the total size of a given xdp buff (linear and paged area)

### Returns

The total size of a given xdp buffer.

`#!c static __u64 (*bpf_xdp_get_buff_len)(struct xdp_md *xdp_md) = (void *) 188;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
