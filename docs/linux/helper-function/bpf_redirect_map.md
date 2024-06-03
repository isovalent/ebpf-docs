---
title: "Helper Function 'bpf_redirect_map'"
description: "This page documents the 'bpf_redirect_map' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_redirect_map`

<!-- [FEATURE_TAG](bpf_redirect_map) -->
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/546ac1ffb70d25b56c1126940e5ec639c4dd7413)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Redirect the packet to the endpoint referenced by _map_ at index _key_. Depending on its type, this _map_ can contain references to net devices (for forwarding packets through other ports), or to CPUs (for redirecting XDP frames to another CPU; but this is only implemented for native XDP (with driver support) as of this writing).

The lower two bits of _flags_ are used as the return code if the map lookup fails. This is so that the return value can be one of the XDP program return codes up to **XDP_TX**, as chosen by the caller. The higher bits of _flags_ can be set to BPF_F_BROADCAST or BPF_F_EXCLUDE_INGRESS as defined below.

With BPF_F_BROADCAST the packet will be broadcasted to all the interfaces in the map, with BPF_F_EXCLUDE_INGRESS the ingress interface will be excluded when do broadcasting.

See also **bpf_redirect**(), which only supports redirecting to an ifindex, but doesn't require a map to do so.

### Returns

**XDP_REDIRECT** on success, or the value of the two lower bits of the _flags_ argument on error.

`#!c static long (* const bpf_redirect_map)(void *map, __u64 key, __u64 flags) = (void *) 51;`
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
