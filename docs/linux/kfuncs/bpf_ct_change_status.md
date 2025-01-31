---
title: "KFunc 'bpf_ct_change_status'"
description: "This page documents the 'bpf_ct_change_status' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_ct_change_status`

<!-- [FEATURE_TAG](bpf_ct_change_status) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/ef69aa3a986ef94f01ce8b5b619f550db54432fe)
<!-- [/FEATURE_TAG] -->

Change status of inserted `nf_conn`

## Definition

Change the status field of the provided connection tracking entry.
This must be invoked for referenced `PTR_TO_BTF_ID` to `nf_conn`.

**Parameters**

`nfct`: Pointer to referenced nf_conn object, obtained using [`bpf_ct_insert_entry`](bpf_ct_insert_entry.md), [`bpf_xdp_ct_lookup`](bpf_xdp_ct_lookup.md) or [`bpf_skb_ct_lookup`](bpf_skb_ct_lookup.md).

`status`: New status value.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_ct_change_status(struct nf_conn *nfct, u32 status)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [`bpf_ct_set_nat_info`](bpf_ct_set_nat_info.md#example) for an example of how to use this kfunc.
