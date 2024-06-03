---
title: "KFunc 'bpf_ct_set_timeout'"
description: "This page documents the 'bpf_ct_set_timeout' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_ct_set_timeout`

<!-- [FEATURE_TAG](bpf_ct_set_timeout) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/0b3892364431684e883682b85d008979e08d4ce6)
<!-- [/FEATURE_TAG] -->

Set timeout of allocated nf_conn

## Definition

Sets the default timeout of newly allocated nf_conn before insertion.
This helper must be invoked for refcounted pointer to nf_conn___init.

**Parameters**

`nfct`: Pointer to referenced nf_conn object, obtained using [`bpf_xdp_ct_alloc`](bpf_xdp_ct_alloc.md) or [`bpf_skb_ct_alloc`](bpf_skb_ct_alloc.md).

`timeout`: Timeout in msecs.

<!-- [KFUNC_DEF] -->
`#!c void bpf_ct_set_timeout(struct nf_conn___init *nfct, u32 timeout)`
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
