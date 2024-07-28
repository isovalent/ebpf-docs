---
title: "KFunc 'bpf_ct_set_status'"
description: "This page documents the 'bpf_ct_set_status' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_ct_set_status`

<!-- [FEATURE_TAG](bpf_ct_set_status) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/ef69aa3a986ef94f01ce8b5b619f550db54432fe)
<!-- [/FEATURE_TAG] -->

Set status field of allocated `nf_conn`

## Definition

Set the status field of the newly allocated `nf_conn` before insertion.
This must be invoked for referenced `PTR_TO_BTF_ID` to `nf_conn___init`.

**Parameters**

`nfct`: Pointer to referenced `nf_conn` object, obtained using [`bpf_xdp_ct_alloc`](bpf_xdp_ct_alloc.md) or [`bpf_skb_ct_alloc`](bpf_skb_ct_alloc.md).

`status`: New status value.

<!-- [KFUNC_DEF] -->
`#!c int bpf_ct_set_status(const struct nf_conn___init *nfct, u32 status)`
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
