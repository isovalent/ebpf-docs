---
title: "KFunc 'tcp_reno_ssthresh'"
description: "This page documents the 'tcp_reno_ssthresh' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `tcp_reno_ssthresh`

<!-- [FEATURE_TAG](tcp_reno_ssthresh) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Returns slow start threshold

## Definition

Slow start threshold is half the congestion window (min 2)

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 tcp_reno_ssthresh(struct sock *sk)`
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

