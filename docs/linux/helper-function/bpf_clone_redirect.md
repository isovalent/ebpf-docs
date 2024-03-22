---
title: "Helper Function 'bpf_clone_redirect'"
description: "This page documents the 'bpf_clone_redirect' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_clone_redirect`

<!-- [FEATURE_TAG](bpf_clone_redirect) -->
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/3896d655f4d491c67d669a15f275a39f713410f8)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Clone and redirect the packet associated to _skb_ to another net device of index _ifindex_. Both ingress and egress interfaces can be used for redirection. The **BPF_F_INGRESS** value in _flags_ is used to make the distinction (ingress path is selected if the flag is present, egress path otherwise). This is the only flag supported for now.

In comparison with **bpf_redirect**() helper, **bpf_clone_redirect**() has the associated cost of duplicating the packet buffer, but this can be executed out of the eBPF program. Conversely, **bpf_redirect**() is more efficient, but it is handled through an action code where the redirection happens only after the eBPF program has returned.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure. Positive error indicates a potential drop or congestion in the target device. The particular positive error codes are not defined.

`#!c static long (*bpf_clone_redirect)(struct __sk_buff *skb, __u32 ifindex, __u64 flags) = (void *) 13;`
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

```c
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <linux/pkt_cls.h>


SEC("tc/egress")
int bpf_clone_redirect_example(struct __sk_buff *skb) {

    __u32 if_index = 2; // interface index to redirect to

    int ret = bpf_clone_redirect(skb, if_index, 0); // redirect to egress path because BPF_F_INGRESS flag is not set

    if (ret) {
        bpf_printk("bpf_clone_redirect error: %d", ret);
    }

    return TC_ACT_OK;
}

char LICENSE[] SEC("license") = "Dual BSD/GPL";
```
