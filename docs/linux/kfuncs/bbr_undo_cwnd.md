---
title: "KFunc 'bbr_undo_cwnd'"
description: "This page documents the 'bbr_undo_cwnd' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bbr_undo_cwnd`

<!-- [FEATURE_TAG](bbr_undo_cwnd) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Return new value of congestion window after loss.

## Definition

In theory BBR does not need to undo the cwnd since it does not always reduce cwnd on losses (see bbr_main()).

<!-- [KFUNC_DEF] -->
`#!c u32 bbr_undo_cwnd(struct sock *sk)`
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

