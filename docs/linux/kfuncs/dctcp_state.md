---
title: "KFunc 'dctcp_state'"
description: "This page documents the 'dctcp_state' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `dctcp_state`

<!-- [FEATURE_TAG](dctcp_state) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e78aea8b2170be1b88c96a4d138422986a737336)
<!-- [/FEATURE_TAG] -->

Call before changing ca_state

## Definition

<!-- [KFUNC_DEF] -->
`#!c void dctcp_state(struct sock *sk, u8 new_state)`
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

