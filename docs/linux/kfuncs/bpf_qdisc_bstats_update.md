---
title: "KFunc 'bpf_qdisc_bstats_update'"
description: "This page documents the 'bpf_qdisc_bstats_update' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_qdisc_bstats_update`

<!-- [FEATURE_TAG](bpf_qdisc_bstats_update) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/544e0a1f1e56de5a28251c188aa8f78fe50b31c9)
<!-- [/FEATURE_TAG] -->

Update Qdisc basic statistics.

## Definition

**Parameters**

`sch`: The qdisc from which an skb is dequeued.

`skb`: The skb to be dequeued.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_qdisc_bstats_update(struct Qdisc *sch, const struct sk_buff *skb)`
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

