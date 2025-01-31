---
title: "KFunc 'tcp_slow_start'"
description: "This page documents the 'tcp_slow_start' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `tcp_slow_start`

<!-- [FEATURE_TAG](tcp_slow_start) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Slow start is used when congestion window is no greater than the slow start threshold.

## Definition

Slow start is used when congestion window is no greater than the slow start threshold. We base on RFC2581 and also handle stretch ACKs properly. We do not implement RFC3465 Appropriate Byte Counting (ABC) <nospell>per se</nospell> but something better;) a packet is only considered (s)acked in its entirety to defend the ACK attacks described in the RFC. Slow start processes a stretch ACK of degree N as if N ACKs of degree 1 are received back to back except ABC caps N to 2. Slow start exits when `cwnd` grows over `ssthresh` and returns the leftover ACKs to adjust `cwnd` in congestion avoidance mode.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 tcp_slow_start(struct tcp_sock *tp, u32 acked)`
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

