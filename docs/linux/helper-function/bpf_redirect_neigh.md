# Helper function `bpf_redirect_neigh`

<!-- [FEATURE_TAG](bpf_redirect_neigh) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/b4ab31414970a7a03a5d55d75083f2c101a30592)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Redirect the packet to another net device of index _ifindex_ and fill in L2 addresses from neighboring subsystem. This helper is somewhat similar to **bpf_redirect**(), except that it populates L2 addresses as well, meaning, internally, the helper relies on the neighbor lookup for the L2 address of the nexthop.

The helper will perform a FIB lookup based on the skb's networking header to get the address of the next hop, unless this is supplied by the caller in the _params_ argument. The _plen_ argument indicates the len of _params_ and should be set to 0 if _params_ is NULL.

The _flags_ argument is reserved and must be 0. The helper is currently only supported for tc BPF program types, and enabled for IPv4 and IPv6 protocols.

### Returns

The helper returns **TC_ACT_REDIRECT** on success or **TC_ACT_SHOT** on error.

`#!c static long (*bpf_redirect_neigh)(__u32 ifindex, struct bpf_redir_neigh *params, int plen, __u64 flags) = (void *) 152;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
