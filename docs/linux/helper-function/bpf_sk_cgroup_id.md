---
title: "Helper Function 'bpf_sk_cgroup_id'"
description: "This page documents the 'bpf_sk_cgroup_id' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_cgroup_id`

<!-- [FEATURE_TAG](bpf_sk_cgroup_id) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f307fa2cb4c935f7f1ff0aeb880c7b44fb9a642b)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return the cgroup v2 id of the socket _sk_.

_sk_ must be a non-**NULL** pointer to a socket, e.g. one returned from **bpf_sk_lookup_xxx**(), **bpf_sk_fullsock**(), etc. The format of returned id is same as in **bpf_skb_cgroup_id**().

This helper is available only if the kernel was compiled with the **CONFIG_SOCK_CGROUP_DATA** configuration option.

### Returns

The id is returned or 0 in case the id could not be retrieved.

`#!c static __u64 (* const bpf_sk_cgroup_id)(void *sk) = (void *) 128;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
