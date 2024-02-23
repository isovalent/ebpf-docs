---
title: "KFunc 'bpf_ct_change_timeout'"
description: "This page documents the 'bpf_ct_change_timeout' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_ct_change_timeout`

<!-- [FEATURE_TAG](bpf_ct_change_timeout) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/0b3892364431684e883682b85d008979e08d4ce6)
<!-- [/FEATURE_TAG] -->

Change timeout of inserted nf_conn

## Definition

Change timeout associated of the inserted or looked up nf_conn.
This helper must be invoked for refcounted pointer to nf_conn.

**Parameters**

`nfct`: Pointer to referenced nf_conn object, obtained using [`bpf_ct_insert_entry`](bpf_ct_insert_entry.md), [`bpf_xdp_ct_lookup`](bpf_xdp_ct_lookup.md), or [`bpf_skb_ct_lookup`](bpf_skb_ct_lookup.md).

`timeout`: New timeout in msecs.

<!-- [KFUNC_DEF] -->
`#!c int bpf_ct_change_timeout(struct nf_conn *nfct, u32 timeout)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [bpf_ct_set_nat_info](bpf_ct_set_nat_info.md#example) for an example of how to use this kfunc.
