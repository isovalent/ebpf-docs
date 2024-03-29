---
title: "Helper Function 'bpf_set_hash_invalid'"
description: "This page documents the 'bpf_set_hash_invalid' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_set_hash_invalid`

<!-- [FEATURE_TAG](bpf_set_hash_invalid) -->
[:octicons-tag-24: v4.9](https://github.com/torvalds/linux/commit/7a4b28c6cc9ffac50f791b99cc7e46106436e5d8)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Invalidate the current _skb_**->hash**. It can be used after mangling on headers through direct packet access, in order to indicate that the hash is outdated and to trigger a recalculation the next time the kernel tries to access this hash or when the **bpf_get_hash_recalc**() helper is called.

### Returns

void.

`#!c static void (* const bpf_set_hash_invalid)(struct __sk_buff *skb) = (void *) 41;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
