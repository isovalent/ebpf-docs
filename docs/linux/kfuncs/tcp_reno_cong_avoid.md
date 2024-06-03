---
title: "KFunc 'tcp_reno_cong_avoid'"
description: "This page documents the 'tcp_reno_cong_avoid' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `tcp_reno_cong_avoid`

<!-- [FEATURE_TAG](tcp_reno_cong_avoid) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Do new congestion window calculation

## Definition

This is Jacobson's slow start and congestion avoidance. SIGCOMM '88, p. 328.

<!-- [KFUNC_DEF] -->
`#!c void tcp_reno_cong_avoid(struct sock *sk, u32 ack, u32 acked)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

