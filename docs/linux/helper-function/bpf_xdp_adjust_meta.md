---
title: "Helper Function 'bpf_xdp_adjust_meta'"
description: "This page documents the 'bpf_xdp_adjust_meta' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_adjust_meta`

<!-- [FEATURE_TAG](bpf_xdp_adjust_meta) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/de8f3a83b0a0fddb2cf56e7a718127e9619ea3da)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Adjust the address pointed by _xdp_md_**->data_meta** by _delta_ (which can be positive or negative). Note that this operation modifies the address stored in _xdp_md_**->data**, so the latter must be loaded only after the helper has been called.

The use of _xdp_md_**->data_meta** is optional and programs are not required to use it. The rationale is that when the packet is processed with XDP (e.g. as DoS filter), it is possible to push further meta data along with it before passing to the stack, and to give the guarantee that an ingress eBPF program attached as a TC classifier on the same device can pick this up for further post-processing. Since TC works with socket buffers, it remains possible to set from XDP the **mark** or **priority** pointers, or other pointers for the socket buffer. Having this scratch space generic and programmable allows for more flexibility as the user is free to store whatever meta data they need.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_xdp_adjust_meta)(struct xdp_md *xdp_md, int delta) = (void *) 54;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/pkt_cls.h>
#include <bpf/bpf_helpers.h>


/*
 * Metadata written by XDP and later consumed by TC.
 * Size and alignment must satisfy kernel metadata constraints
 * (typically 4-byte aligned).
 */
struct test_meta {
	__u32 mark;
} __attribute__((aligned(4)));

const __u32 marker = 0x1234;

SEC("xdp")
int xdp_store_meta(struct xdp_md *ctx) {
	int ret = 0;

	/*
     * 1) Reserve metadata space by moving data_meta backwards.
     *    Negative delta expands the metadata area.
     */
	ret = bpf_xdp_adjust_meta(ctx, -(int)sizeof(struct test_meta));
	if (ret < 0)
		return XDP_ABORTED;

	/*
     * 2) Reload pointers after calling helper.
     *    The helper may update ctx fields, so old pointers are invalid.
     */
	void *data		= (void *)(long)ctx->data;
	void *data_meta	= (void *)(long)ctx->data_meta;

	
    // 3) Bounds check. Metadata region must stay below data.
	if (data_meta + sizeof(struct test_meta) > data)
		return XDP_ABORTED;

    // 4) Write metadata.
	struct test_meta *m = data_meta;
	m->mark = marker;

	return XDP_PASS;
}

SEC("tc")
int tc_load_meta(struct __sk_buff *skb) {
	void *data		= (void *)(long)skb->data;
	void *data_meta	= (void *)(long)skb->data_meta;

	if (data_meta >= data) {
		return TC_ACT_SHOT;
	}
	

	// Bounds check: metadata must not overlap packet data.
	if (data_meta + sizeof(struct test_meta) > data)
		return TC_ACT_SHOT;

	struct test_meta *m = data_meta;
	if (m->mark != marker) {
		return TC_ACT_SHOT;
	}

	skb->mark = m->mark;

	return TC_ACT_OK;
}

char _license[] SEC("license") = "GPL";
```
