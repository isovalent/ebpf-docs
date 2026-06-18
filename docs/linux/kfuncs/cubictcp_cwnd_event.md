---
title: "KFunc 'cubictcp_cwnd_event'"
description: "This page documents the 'cubictcp_cwnd_event' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `cubictcp_cwnd_event`

<!-- [FEATURE_TAG](cubictcp_cwnd_event) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Default cubic TCP implementation of [`tcp_congestion_ops->cwnd_event`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/tcp_congestion_ops.md#cwnd_event).

!!! warning
    This kfunc got removed and replaced by [`cubictcp_cwnd_event_tx_start`](cubictcp_cwnd_event_tx_start.md) in [:octicons-tag-24: v7.1](https://github.com/torvalds/linux/commit/d1e59a46973719e458bec78d00dd767d7a7ba71f)


## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void cubictcp_cwnd_event(struct sock *sk, tcp_ca_event event)`
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

