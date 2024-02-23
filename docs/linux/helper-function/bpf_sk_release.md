---
title: "Helper Function 'bpf_sk_release' - eBPF Docs"
description: "This page documents the 'bpf_sk_release' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_release`

<!-- [FEATURE_TAG](bpf_sk_release) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/6acc9b432e6714d72d7d77ec7c27f6f8358d0c71)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Release the reference held by _sock_. _sock_ must be a non-**NULL** pointer that was returned from **bpf_sk_lookup_xxx**().

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_sk_release)(void *sock) = (void *) 86;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
