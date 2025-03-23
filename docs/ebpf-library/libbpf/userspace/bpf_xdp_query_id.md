---
title: "Libbpf userspace function 'bpf_xdp_query_id'"
description: "This page documents the 'bpf_xdp_query_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_xdp_query_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Query the program ID of the XDP program attached to a network interface.

## Definition

`#!c int bpf_xdp_query_id(int ifindex, int flags, __u32 *prog_id);`

**Parameters**

- `ifindex`: The index of the network interface to query.
- `flags`: Flags to control the query.
- `prog_id`: A pointer to a `__u32` that will be filled with the program ID of the XDP program attached to the network interface.

**Flags**

* `XDP_FLAGS_SKB_MODE` = `(1U << 1)` - If set, query for programs in SKB (generic) mode.
* `XDP_FLAGS_DRV_MODE` = `(1U << 2)` - If set, query for programs in DRV (driver / native) mode.
* `XDP_FLAGS_HW_MODE` = `(1U << 3)` - If set, query for programs in hardware offload mode.

**Return**

`0` on success. A negative error code on failure.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
