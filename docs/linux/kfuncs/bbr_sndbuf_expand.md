---
title: "KFunc 'bbr_sndbuf_expand'"
description: "This page documents the 'bbr_sndbuf_expand' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bbr_sndbuf_expand`

<!-- [FEATURE_TAG](bbr_sndbuf_expand) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Returns the multiplier used in `tcp_sndbuf_expand`

## Definition

Provision 3 * `cwnd` since BBR may slow-start even during recovery.

<!-- [KFUNC_DEF] -->
`#!c u32 bbr_sndbuf_expand(struct sock *sk)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

