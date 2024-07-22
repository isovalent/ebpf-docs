---
title: "Helper Function 'bpf_redirect'"
description: "This page documents the 'bpf_redirect' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_redirect`

<!-- [FEATURE_TAG](bpf_redirect) -->
[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/27b29f63058d26c6c1742f1993338280d5a41dc6)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Redirect the packet to another net device of index _ifindex_. This helper is somewhat similar to **bpf_clone_redirect**\ (), except that the packet is not cloned, which provides increased performance.

Except for XDP, both ingress and egress interfaces can be used for redirection. The **BPF_F_INGRESS** value in _flags_ is used to make the distinction (ingress path is selected if the flag is present, egress path otherwise). Currently, XDP only supports redirection to the egress interface, and accepts no flag at all.

The same effect can also be attained with the more generic **bpf_redirect_map**(), which uses a BPF map to store the redirect target instead of providing it directly to the helper.

### Returns

For XDP, the helper returns **XDP_REDIRECT** on success or **XDP_ABORTED** on error. For other program types, the values are **TC_ACT_REDIRECT** on success or **TC_ACT_SHOT** on error.

`#!c static long (* const bpf_redirect)(__u32 ifindex, __u64 flags) = (void *) 23;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>


SEC("tc/egress")
int bpf_clone_redirect_example(struct __sk_buff *skb) {

    __u32 if_index = 2; // interface index to redirect to

    return bpf_redirect(if_index, 0);
}

char LICENSE[] SEC("license") = "Dual BSD/GPL";
```
