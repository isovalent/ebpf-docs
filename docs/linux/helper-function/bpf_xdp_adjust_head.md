---
title: "Helper Function 'bpf_xdp_adjust_head'"
description: "This page documents the 'bpf_xdp_adjust_head' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_adjust_head`

<!-- [FEATURE_TAG](bpf_xdp_adjust_head) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/17bedab2723145d17b14084430743549e6943d03)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Adjust (move) _xdp_md_**->data** by _delta_ bytes. Note that it is possible to use a negative value for _delta_. This helper can be used to prepare the packet for pushing or popping headers.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_xdp_adjust_head)(struct xdp_md *xdp_md, int delta) = (void *) 44;`
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
