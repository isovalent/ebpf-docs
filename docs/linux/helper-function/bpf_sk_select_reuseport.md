---
title: "Helper Function 'bpf_sk_select_reuseport'"
description: "This page documents the 'bpf_sk_select_reuseport' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_select_reuseport`

<!-- [FEATURE_TAG](bpf_sk_select_reuseport) -->
[:octicons-tag-24: v4.19](https://github.com/torvalds/linux/commit/2dbb9b9e6df67d444fbe425c7f6014858d337adf)
<!-- [/FEATURE_TAG] -->

The socket select reuse port helper select which socket to send an incoming request to when multiple sockets are bound to the same port.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_sk_select_reuseport)(struct sk_reuseport_md *reuse, void *map, void *key, __u64 flags) = (void *) 82;`

## Usage

In [:octicons-tag-24: v3.9](https://github.com/torvalds/linux/commit/c617f398edd4db2b8567a28e899a88f8f574798d) the [`SO_REUSEPORT`](https://lwn.net/Articles/542629/) socket option was added which allows multiple sockets to listen to the same port on the same host. The original purpose of the feature being that this allows for high-efficient distribution of traffic across threads which would normally have to be done in userspace causing unnecessary delay.

By default, incoming connections and datagrams are distributed to the server sockets using a hash based on the 4-tuple of the connection—that is, the peer IP address and port plus the local IP address and port.

With the introduction of [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md) programs and this helper we can replace the default distribution behavior with a BPF program. This helper does the actual assigning of an incoming request to a socket in a [`BPF_MAP_TYPE_REUSEPORT_SOCKARRAY`](../map-type/BPF_MAP_TYPE_REUSEPORT_SOCKARRAY.md) `map`.

Since [:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/64d85290d79c0677edb5a8ee2295b36c022fa5df)  [BPF_MAP_TYPE_SOCKHASH](../map-type/BPF_MAP_TYPE_SOCKHASH.md) and [BPF_MAP_TYPE_SOCKMAP](../map-type/BPF_MAP_TYPE_SOCKMAP.md) maps can also be used with this helper.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [BPF_MAP_TYPE_REUSEPORT_SOCKARRAY](../map-type/BPF_MAP_TYPE_REUSEPORT_SOCKARRAY.md)
 * [BPF_MAP_TYPE_SOCKHASH](../map-type/BPF_MAP_TYPE_SOCKHASH.md)
 * [BPF_MAP_TYPE_SOCKMAP](../map-type/BPF_MAP_TYPE_SOCKMAP.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
