---
title: "KFunc 'bbr_cwnd_event'"
description: "This page documents the 'bbr_cwnd_event' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bbr_cwnd_event`

<!-- [FEATURE_TAG](bbr_cwnd_event) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Call when congestion window event occurs.

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bbr_cwnd_event(struct sock *sk, tcp_ca_event event)`
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

