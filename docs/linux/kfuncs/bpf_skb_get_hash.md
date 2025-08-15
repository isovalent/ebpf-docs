---
title: "KFunc 'bpf_skb_get_hash'"
description: "This page documents the 'bpf_skb_get_hash' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_skb_get_hash`

<!-- [FEATURE_TAG](bpf_skb_get_hash) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/870c28588afa20d786e2c26e8fcc14e2b9a55616)
<!-- [/FEATURE_TAG] -->

Get the flow hash of an skb.

## Definition

**Parameters**

`skb`: The skb to get the flow hash from.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u32 bpf_skb_get_hash(struct sk_buff *skb)`
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

