---
title: "Helper Function 'bpf_get_route_realm'"
description: "This page documents the 'bpf_get_route_realm' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_route_realm`

<!-- [FEATURE_TAG](bpf_get_route_realm) -->
[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/c46646d0484f5d08e2bede9b45034ba5b8b489cc)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Retrieve the realm or the route, that is to say the **tclassid** field of the destination for the _skb_. The identifier retrieved is a user-provided tag, similar to the one used with the net_cls cgroup (see description for **bpf_get_cgroup_classid**() helper), but here this tag is held by a route (a destination entry), not by a task.

Retrieving this identifier works with the clsact TC egress hook (see also **tc-bpf(8)**), or alternatively on conventional classful egress qdiscs, but not on TC ingress path. In case of clsact TC egress hook, this has the advantage that, internally, the destination entry has not been dropped yet in the transmit path. Therefore, the destination entry does not need to be artificially held via **netif_keep_dst**() for a classful qdisc until the _skb_ is freed.

This helper is available only if the kernel was compiled with **CONFIG_IP_ROUTE_CLASSID** configuration option.

### Returns

The realm of the route for the packet associated to _skb_, or 0 if none was found.

`#!c static __u32 (* const bpf_get_route_realm)(struct __sk_buff *skb) = (void *) 24;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
