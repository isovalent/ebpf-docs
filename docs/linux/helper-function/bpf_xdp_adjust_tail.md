---
title: "Helper Function 'bpf_xdp_adjust_tail'"
description: "This page documents the 'bpf_xdp_adjust_tail' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_adjust_tail`

<!-- [FEATURE_TAG](bpf_xdp_adjust_tail) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/b32cc5b9a346319c171e3ad905e0cddda032b5eb)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Adjust (move) _xdp_md_**->data_end** by _delta_ bytes. It is possible to both shrink and grow the packet tail. Shrink done via _delta_ being a negative integer.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_xdp_adjust_tail)(struct xdp_md *xdp_md, int delta) = (void *) 65;`
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
