---
title: "Helper Function 'bpf_skb_ancestor_cgroup_id' - eBPF Docs"
description: "This page documents the 'bpf_skb_ancestor_cgroup_id' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_ancestor_cgroup_id`

<!-- [FEATURE_TAG](bpf_skb_ancestor_cgroup_id) -->
[:octicons-tag-24: v4.19](https://github.com/torvalds/linux/commit/7723628101aaeb1d723786747529b4ea65c5b5c5)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return id of cgroup v2 that is ancestor of cgroup associated with the _skb_ at the _ancestor_level_.  The root cgroup is at _ancestor_level_ zero and each step down the hierarchy increments the level. If _ancestor_level_ == level of cgroup associated with _skb_, then return value will be same as that of **bpf_skb_cgroup_id**().

The helper is useful to implement policies based on cgroups that are upper in hierarchy than immediate cgroup associated with _skb_.

The format of returned id and helper limitations are same as in **bpf_skb_cgroup_id**().

### Returns

The id is returned or 0 in case the id could not be retrieved.

`#!c static __u64 (*bpf_skb_ancestor_cgroup_id)(struct __sk_buff *skb, int ancestor_level) = (void *) 83;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
