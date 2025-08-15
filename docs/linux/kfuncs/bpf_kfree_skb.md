---
title: "KFunc 'bpf_kfree_skb'"
description: "This page documents the 'bpf_kfree_skb' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_kfree_skb`

<!-- [FEATURE_TAG](bpf_kfree_skb) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/870c28588afa20d786e2c26e8fcc14e2b9a55616)
<!-- [/FEATURE_TAG] -->

Release an skb's reference and drop it immediately.

## Definition

**Parameters**

`skb`: The skb whose reference to be released and dropped.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_kfree_skb(struct sk_buff *skb)`
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

