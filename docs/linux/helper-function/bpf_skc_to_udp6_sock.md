---
title: "Helper Function 'bpf_skc_to_udp6_sock'"
description: "This page documents the 'bpf_skc_to_udp6_sock' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_skc_to_udp6_sock`

<!-- [FEATURE_TAG](bpf_skc_to_udp6_sock) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/0d4fad3e57df2bf61e8ffc8d12a34b1caf9b8835)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Dynamically cast a _sk_ pointer to a _udp6_sock_ pointer.

### Returns

_sk_ if casting is valid, or **NULL** otherwise.

`#!c static struct udp6_sock *(* const bpf_skc_to_udp6_sock)(void *sk) = (void *) 140;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
