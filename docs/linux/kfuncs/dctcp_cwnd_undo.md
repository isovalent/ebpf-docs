---
title: "KFunc 'dctcp_cwnd_undo' - eBPF Docs"
description: "This page documents the 'dctcp_cwnd_undo' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `dctcp_cwnd_undo`

<!-- [FEATURE_TAG](dctcp_cwnd_undo) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Return new value of congestion window after loss.

## Definition

<!-- [KFUNC_DEF] -->
`#!c u32 dctcp_cwnd_undo(struct sock *sk)`
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

