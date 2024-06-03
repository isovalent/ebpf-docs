---
title: "Helper Function 'bpf_fib_lookup'"
description: "This page documents the 'bpf_fib_lookup' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_fib_lookup`

<!-- [FEATURE_TAG](bpf_fib_lookup) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/87f5fc7e48dd3175b30dd03b41564e1a8e136323)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Do FIB lookup in kernel tables using parameters in _params_. If lookup is successful and result shows packet is to be forwarded, the neighbor tables are searched for the nexthop. If successful (ie., FIB lookup shows forwarding and nexthop is resolved), the nexthop address is returned in ipv4_dst or ipv6_dst based on family, smac is set to mac address of egress device, dmac is set to nexthop mac address, rt_metric is set to metric from route (IPv4/IPv6 only), and ifindex is set to the device index of the nexthop from the FIB lookup.

_plen_ argument is the size of the passed in struct. _flags_ argument can be a combination of one or more of the following values:

**BPF_FIB_LOOKUP_DIRECT**

&nbsp;&nbsp;&nbsp;&nbsp;Do a direct table lookup vs full lookup using FIB rules.

**BPF_FIB_LOOKUP_TBID**

&nbsp;&nbsp;&nbsp;&nbsp;Used with BPF_FIB_LOOKUP_DIRECT. Use the routing table ID present in _params_->tbid for the fib lookup.

**BPF_FIB_LOOKUP_OUTPUT**

&nbsp;&nbsp;&nbsp;&nbsp;Perform lookup from an egress perspective (default is ingress).

**BPF_FIB_LOOKUP_SKIP_NEIGH**

&nbsp;&nbsp;&nbsp;&nbsp;Skip the neighbour table lookup. _params_->dmac and _params_->smac will not be set as output. A common use case is to call **bpf_redirect_neigh**() after doing **bpf_fib_lookup**().

**BPF_FIB_LOOKUP_SRC**

&nbsp;&nbsp;&nbsp;&nbsp;Derive and set source IP addr in _params_->ipv{4,6}_src for the nexthop. If the src addr cannot be derived, **BPF_FIB_LKUP_RET_NO_SRC_ADDR** is returned. In this case, _params_->dmac and _params_->smac are not set either.

**BPF_FIB_LOOKUP_MARK**

&nbsp;&nbsp;&nbsp;&nbsp;Use the mark present in _params_->mark for the fib lookup. This option should not be used with BPF_FIB_LOOKUP_DIRECT, as it only has meaning for full lookups.

_ctx_ is either **struct xdp_md** for XDP programs or **struct sk_buff** tc cls_act programs.

### Returns

* < 0 if any input argument is invalid
*   0 on success (packet is forwarded, nexthop neighbor exists)
* > 0 one of **BPF_FIB_LKUP_RET_** codes explaining why the
  packet is not forwarded or needs assist from full stack

If lookup fails with BPF_FIB_LKUP_RET_FRAG_NEEDED, then the MTU was exceeded and output params->mtu_result contains the MTU.

`#!c static long (* const bpf_fib_lookup)(void *ctx, struct bpf_fib_lookup *params, int plen, __u32 flags) = (void *) 69;`
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
