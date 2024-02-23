---
title: "KFunc 'tcp_cong_avoid_ai' - eBPF Docs"
description: "This page documents the 'tcp_cong_avoid_ai' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `tcp_cong_avoid_ai`

<!-- [FEATURE_TAG](tcp_cong_avoid_ai) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

## Definition

In theory this is tp->snd_cwnd += 1 / tp->snd_cwnd (or alternative w), for every packet that was ACKed.

<!-- [KFUNC_DEF] -->
`#!c void tcp_cong_avoid_ai(struct tcp_sock *tp, u32 w, u32 acked)`
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

