---
title: "Helper Function 'bpf_skb_pull_data' - eBPF Docs"
description: "This page documents the 'bpf_skb_pull_data' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_pull_data`

<!-- [FEATURE_TAG](bpf_skb_pull_data) -->
[:octicons-tag-24: v4.9](https://github.com/torvalds/linux/commit/36bbef52c7eb646ed6247055a2acd3851e317857)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Pull in non-linear data in case the _skb_ is non-linear and not all of _len_ are part of the linear section. Make _len_ bytes from _skb_ readable and writable. If a zero value is passed for _len_, then all bytes in the linear part of _skb_ will be made readable and writable.

This helper is only needed for reading and writing with direct packet access.

For direct packet access, testing that offsets to access are within packet boundaries (test on _skb_**->data_end**) is susceptible to fail if offsets are invalid, or if the requested data is in non-linear parts of the _skb_. On failure the program can just bail out, or in the case of a non-linear buffer, use a helper to make the data available. The **bpf_skb_load_bytes**() helper is a first solution to access the data. Another one consists in using **bpf_skb_pull_data** to pull in once the non-linear parts, then retesting and eventually access the data.

At the same time, this also makes sure the _skb_ is uncloned, which is a necessary condition for direct write. As this needs to be an invariant for the write part only, the verifier detects writes and adds a prologue that is calling **bpf_skb_pull_data()** to effectively unclone the _skb_ from the very beginning in case it is indeed cloned.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_skb_pull_data)(struct __sk_buff *skb, __u32 len) = (void *) 39;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
