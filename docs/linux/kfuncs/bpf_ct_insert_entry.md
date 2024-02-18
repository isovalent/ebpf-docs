# KFunc `bpf_ct_insert_entry`

<!-- [FEATURE_TAG](bpf_ct_insert_entry) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/d7e79c97c00ca82dace0e3b645d4b3b02fa273c2)
<!-- [/FEATURE_TAG] -->

Add the provided entry into a CT map

## Definition

This must be invoked for referenced PTR_TO_BTF_ID.

**Parameters**

`nfct__ref`: Pointer to referenced nf_conn___init object, obtained using bpf_xdp_ct_alloc or bpf_skb_ct_alloc.

<!-- [KFUNC_DEF] -->
`#!c struct nf_conn *bpf_ct_insert_entry(struct nf_conn___init *nfct_i)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).

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
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [bpf_ct_set_nat_info](bpf_ct_set_nat_info.md#example) for an example of how to use this kfunc.
