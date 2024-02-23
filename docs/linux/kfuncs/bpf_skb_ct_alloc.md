---
title: "KFunc 'bpf_skb_ct_alloc' - eBPF Docs"
description: "This page documents the 'bpf_skb_ct_alloc' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_skb_ct_alloc`

<!-- [FEATURE_TAG](bpf_skb_ct_alloc) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/d7e79c97c00ca82dace0e3b645d4b3b02fa273c2)
<!-- [/FEATURE_TAG] -->

Allocate a new CT entry

## Definition

**Parameters**

`skb_ctx`: Pointer to ctx (__sk_buff) in TC program. Cannot be NULL

`bpf_tuple`: Pointer to memory representing the tuple to look up. Cannot be NULL

`tuple__sz`: Length of the tuple structure. Must be one of sizeof(bpf_tuple->ipv4) or sizeof(bpf_tuple->ipv6)

`opts`: Additional options for lookup `struct bpf_ct_opts`. Cannot be NULL.

**Members**

`opts.netns_id`: Specify the network namespace for allocation, Values:

- `BPF_F_CURRENT_NETNS` - (-1) Use namespace associated with ctx (xdp_md, __sk_buff)
- `[0, S32_MAX]` - Network Namespace ID
  
`opts.error`: Out parameter, set for any errors encountered, Values:

- `-EINVAL` - Passed NULL for bpf_tuple pointer
- `-EINVAL` - opts->reserved is not 0
- `-EINVAL` - netns_id is less than -1
- `-EINVAL` - opts__sz isn't `NF_BPF_CT_OPTS_SZ` (12)
- `-EPROTO` - l4proto isn't one of `IPPROTO_TCP` or `IPPROTO_UDP`
- `-ENONET` - No network namespace found for netns_id
- `-ENOENT` - Conntrack lookup could not find entry for tuple
- `-EAFNOSUPPORT` - tuple__sz isn't one of sizeof(tuple->ipv4) or sizeof(tuple->ipv6)

`opts.l4proto`: Layer 4 protocol, Values: `IPPROTO_TCP`, `IPPROTO_UDP`

`opts.reserved`: Reserved member, will be reused for more options in future, Values: `0`

`opts__sz`: Length of the bpf_ct_opts structure. Must be `NF_BPF_CT_OPTS_SZ` (12)

<!-- [KFUNC_DEF] -->
`#!c struct nf_conn___init *bpf_skb_ct_alloc(struct __sk_buff *skb_ctx, struct bpf_sock_tuple *bpf_tuple, u32 tuple__sz, struct bpf_ct_opts *opts, u32 opts__sz)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
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

