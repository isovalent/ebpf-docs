---
title: "Helper Function 'bpf_sk_ancestor_cgroup_id'"
description: "This page documents the 'bpf_sk_ancestor_cgroup_id' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_ancestor_cgroup_id`

<!-- [FEATURE_TAG](bpf_sk_ancestor_cgroup_id) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f307fa2cb4c935f7f1ff0aeb880c7b44fb9a642b)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return id of cgroup v2 that is ancestor of cgroup associated with the _sk_ at the _ancestor_level_.  The root cgroup is at _ancestor_level_ zero and each step down the hierarchy increments the level. If _ancestor_level_ == level of cgroup associated with _sk_, then return value will be same as that of **bpf_sk_cgroup_id**().

The helper is useful to implement policies based on cgroups that are upper in hierarchy than immediate cgroup associated with _sk_.

The format of returned id and helper limitations are same as in **bpf_sk_cgroup_id**().

### Returns

The id is returned or 0 in case the id could not be retrieved.

`#!c static __u64 (* const bpf_sk_ancestor_cgroup_id)(void *sk, int ancestor_level) = (void *) 129;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
