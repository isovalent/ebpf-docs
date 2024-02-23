---
title: "Helper Function 'bpf_sk_fullsock'"
description: "This page documents the 'bpf_sk_fullsock' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_fullsock`

<!-- [FEATURE_TAG](bpf_sk_fullsock) -->
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/46f8bc92758c6259bcf945e9216098661c1587cd)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
This helper gets a **struct bpf_sock** pointer such that all the fields in this **bpf_sock** can be accessed.

### Returns

A **struct bpf_sock** pointer on success, or **NULL** in case of failure.

`#!c static struct bpf_sock *(*bpf_sk_fullsock)(struct bpf_sock *sk) = (void *) 95;`
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
