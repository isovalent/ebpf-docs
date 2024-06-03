---
title: "Helper Function 'bpf_skc_to_mptcp_sock'"
description: "This page documents the 'bpf_skc_to_mptcp_sock' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_skc_to_mptcp_sock`

<!-- [FEATURE_TAG](bpf_skc_to_mptcp_sock) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/3bc253c2e652cf5f12cd8c00d80d8ec55d67d1a7)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Dynamically cast a _sk_ pointer to a _mptcp_sock_ pointer.

### Returns

_sk_ if casting is valid, or **NULL** otherwise.

`#!c static struct mptcp_sock *(* const bpf_skc_to_mptcp_sock)(void *sk) = (void *) 196;`
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
