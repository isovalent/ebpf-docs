---
title: "Helper Function 'bpf_sk_redirect_map'"
description: "This page documents the 'bpf_sk_redirect_map' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_redirect_map`

<!-- [FEATURE_TAG](bpf_sk_redirect_map) -->
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/174a79ff9515f400b9a6115643dafd62a635b7e6)
<!-- [/FEATURE_TAG] -->

The message redirect map helper is used to redirect a message to a socket referenced by a map.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
`SK_PASS` on success, or `SK_DROP` on error.

`#!c static long (* const bpf_sk_redirect_map)(struct __sk_buff *skb, void *map, __u32 key, __u64 flags) = (void *) 52;`

## Usage

Redirect the packet to the socket referenced by `map` (of type [`BPF_MAP_TYPE_SOCKMAP`](../map-type/BPF_MAP_TYPE_SOCKMAP.md)) at index `key`. Both ingress and egress interfaces can be used for redirection. The `BPF_F_INGRESS` value in `flags` is used to make the distinction (ingress path is selected if the flag is present, egress path otherwise). This is the only flag supported for now.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [`BPF_MAP_TYPE_SOCKMAP`](../map-type/BPF_MAP_TYPE_SOCKMAP.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
