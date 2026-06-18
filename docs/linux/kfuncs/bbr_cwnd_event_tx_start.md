---
title: "KFunc 'bbr_cwnd_event_tx_start'"
description: "This page documents the 'bbr_cwnd_event_tx_start' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bbr_cwnd_event_tx_start`

<!-- [FEATURE_TAG](bbr_cwnd_event_tx_start) -->
[:octicons-tag-24: v7.1](https://github.com/torvalds/linux/commit/d1e59a46973719e458bec78d00dd767d7a7ba71f)
<!-- [/FEATURE_TAG] -->

Default BBR implementation of [`tcp_congestion_ops->cwnd_event_tx_start`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/tcp_congestion_ops.md#cwnd_event_tx_start).


## Definition

**Parameters**

`sk`: The socket given to the `cwnd_event_tx_start` callback.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bbr_cwnd_event_tx_start(struct sock *sk)`
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

