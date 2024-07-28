---
title: "KFunc 'bpf_xdp_metadata_rx_hash'"
description: "This page documents the 'bpf_xdp_metadata_rx_hash' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_xdp_metadata_rx_hash`

<!-- [FEATURE_TAG](bpf_xdp_metadata_rx_hash) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/3d76a4d3d4e591af3e789698affaad88a5a8e8ab)
<!-- [/FEATURE_TAG] -->

Read XDP frame RX hash.

## Definition

If `bpf_xdp_metadata_rx_hash` is not supported by the target device, the default implementation is called instead. The verifier, at load time, replaces a call to the generic kfunc with a call to the per-device one.

<!-- [KFUNC_DEF] -->
`#!c int bpf_xdp_metadata_rx_hash(const struct xdp_md *ctx, u32 *hash, xdp_rss_hash_type *rss_type)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

