---
title: "Helper Function 'bpf_get_listener_sock' - eBPF Docs"
description: "This page documents the 'bpf_get_listener_sock' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_get_listener_sock`

<!-- [FEATURE_TAG](bpf_get_listener_sock) -->
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/dbafd7ddd62369b2f3926ab847cbf8fc40e800b7)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return a **struct bpf_sock** pointer in **TCP_LISTEN** state. **bpf_sk_release**() is unnecessary and not allowed.

### Returns

A **struct bpf_sock** pointer on success, or **NULL** in case of failure.

`#!c static struct bpf_sock *(*bpf_get_listener_sock)(struct bpf_sock *sk) = (void *) 98;`
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
