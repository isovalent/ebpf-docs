# Helper function `bpf_skb_load_bytes_relative`

<!-- [FEATURE_TAG](bpf_skb_load_bytes_relative) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/4e1ec56cdc59746943b2acfab3c171b930187bbe)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
This helper is similar to **bpf_skb_load_bytes**() in that it provides an easy way to load _len_ bytes from _offset_ from the packet associated to _skb_, into the buffer pointed by _to_. The difference to **bpf_skb_load_bytes**() is that a fifth argument _start_header_ exists in order to select a base offset to start from. _start_header_ can be one of:

**BPF_HDR_START_MAC**

&nbsp;&nbsp;&nbsp;&nbsp;Base offset to load data from is _skb_'s mac header.

**BPF_HDR_START_NET**

&nbsp;&nbsp;&nbsp;&nbsp;Base offset to load data from is _skb_'s network header.

In general, "direct packet access" is the preferred method to access packet data, however, this helper is in particular useful in socket filters where _skb_**->data** does not always point to the start of the mac header and where "direct packet access" is not available.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_skb_load_bytes_relative)(const void *skb, __u32 offset, void *to, __u32 len, __u32 start_header) = (void *) 68;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
