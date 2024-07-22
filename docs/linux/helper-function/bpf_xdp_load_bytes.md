---
title: "Helper Function 'bpf_xdp_load_bytes'"
description: "This page documents the 'bpf_xdp_load_bytes' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_load_bytes`

<!-- [FEATURE_TAG](bpf_xdp_load_bytes) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/3f364222d032eea6b245780e845ad213dab28cdd)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
This helper is provided as an easy way to load data from a xdp buffer. It can be used to load _len_ bytes from _offset_ from the frame associated to _xdp_md_, into the buffer pointed by _buf_.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_xdp_load_bytes)(struct xdp_md *xdp_md, __u32 offset, void *buf, __u32 len) = (void *) 189;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
