---
title: "KFunc 'bpf_qdisc_skb_drop'"
description: "This page documents the 'bpf_qdisc_skb_drop' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_qdisc_skb_drop`

<!-- [FEATURE_TAG](bpf_qdisc_skb_drop) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/870c28588afa20d786e2c26e8fcc14e2b9a55616)
<!-- [/FEATURE_TAG] -->

Drop an skb by adding it to a deferred free list.

## Definition

**Parameters**

`skb`: The skb whose reference to be released and dropped.

`to_free_list`: The list of skbs to be dropped.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_qdisc_skb_drop(struct sk_buff *skb, struct bpf_sk_buff_ptr *to_free_list)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is used to add `skb` to the `to_free_list`. Packets added to this `to_free_list` are marked as dropped and their memory freed.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

