# KFunc `bpf_ct_release`

<!-- [FEATURE_TAG](bpf_ct_release) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/b4c2b9593a1c4c3a718370e34af28e817fd5e5c6)
<!-- [/FEATURE_TAG] -->

Release acquired nf_conn object

## Definition

This must be invoked for referenced PTR_TO_BTF_ID, and the verifier rejects the program if any references remain in the program in all of the explored states.

**Parameters**

`nf_conn`: Pointer to referenced nf_conn object, obtained using [`bpf_xdp_ct_lookup`](bpf_xdp_ct_lookup.md) or [`bpf_skb_ct_lookup`](bpf_skb_ct_alloc.md).

<!-- [KFUNC_DEF] -->
`#!c void bpf_ct_release(struct nf_conn *nfct)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_SCHED_CLS](../../program-types/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_XDP](../../program-types/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [bpf_ct_set_nat_info](bpf_ct_set_nat_info.md#example) for an example of how to use this kfunc.
