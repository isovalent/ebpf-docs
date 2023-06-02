# Helper function `bpf_fib_lookup`

<!-- [FEATURE_TAG](bpf_fib_lookup) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/87f5fc7e48dd3175b30dd03b41564e1a8e136323)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Do FIB lookup in kernel tables using parameters in _params_. If lookup is successful and result shows packet is to be forwarded, the neighbor tables are searched for the nexthop. If successful (ie., FIB lookup shows forwarding and nexthop is resolved), the nexthop address is returned in ipv4_dst or ipv6_dst based on family, smac is set to mac address of egress device, dmac is set to nexthop mac address, rt_metric is set to metric from route (IPv4/IPv6 only), and ifindex is set to the device index of the nexthop from the FIB lookup.

_plen_ argument is the size of the passed in struct. _flags_ argument can be a combination of one or more of the following values:

**BPF_FIB_LOOKUP_DIRECT**

&nbsp;&nbsp;&nbsp;&nbsp;Do a direct table lookup vs full lookup using FIB rules.

**BPF_FIB_LOOKUP_OUTPUT**

&nbsp;&nbsp;&nbsp;&nbsp;Perform lookup from an egress perspective (default is ingress).

**BPF_FIB_LOOKUP_SKIP_NEIGH**

&nbsp;&nbsp;&nbsp;&nbsp;Skip the neighbour table lookup. _params_->dmac and _params_->smac will not be set as output. A common use case is to call **bpf_redirect_neigh**() after doing **bpf_fib_lookup**().

_ctx_ is either **struct xdp_md** for XDP programs or **struct sk_buff** tc cls_act programs.

### Returns

* < 0 if any input argument is invalid
*   0 on success (packet is forwarded, nexthop neighbor exists)
* > 0 one of **BPF_FIB_LKUP_RET_** codes explaining why the
  packet is not forwarded or needs assist from full stack

If lookup fails with BPF_FIB_LKUP_RET_FRAG_NEEDED, then the MTU was exceeded and output params->mtu_result contains the MTU.

`#!c static long (*bpf_fib_lookup)(void *ctx, struct bpf_fib_lookup *params, int plen, __u32 flags) = (void *) 69;`
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
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
