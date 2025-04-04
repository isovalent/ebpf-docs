---
title: "KFunc 'bpf_skb_set_xfrm_info'"
description: "This page documents the 'bpf_skb_set_xfrm_info' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_skb_set_xfrm_info`

<!-- [FEATURE_TAG](bpf_skb_set_xfrm_info) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/94151f5aa9667c562281abeaaa5e89b9d5c17729)
<!-- [/FEATURE_TAG] -->

Set XFRM metadata

## Definition

**Parameters**

`skb_ctx`: Pointer to ctx (__sk_buff) in TC program. Cannot be NULL

`from`: Pointer to memory from which the metadata will be copied. Cannot be NULL

**Members**

`from.if_id`: XFRM if_id:

- Transmit: if_id to be used in policy and state lookups
- Receive: if_id of the state matched for the incoming packet

`from.link`: Underlying device ifindex:

- Transmit: used as the underlying device in VRF routing
- Receive: the device on which the packet had been received

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_skb_set_xfrm_info(struct __sk_buff *skb_ctx, const struct bpf_xfrm_info *from)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc allows steering traffic towards different IPsec connections based on logic implemented in bpf programs.

This object is built based on the availability of BTF debug info.

When setting the xfrm metadata, per-CPU metadata destinations are used in order to avoid allocating a metadata destination per packet.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include "bpf_tracing_net.h"
#include <bpf/bpf_helpers.h>

__u32 req_if_id;
__u32 resp_if_id;

int bpf_skb_set_xfrm_info(struct __sk_buff *skb_ctx,
			  const struct bpf_xfrm_info *from) __ksym;
int bpf_skb_get_xfrm_info(struct __sk_buff *skb_ctx,
			  struct bpf_xfrm_info *to) __ksym;

SEC("tc")
int set_xfrm_info(struct __sk_buff *skb)
{
	struct bpf_xfrm_info info = { .if_id = req_if_id };

	return bpf_skb_set_xfrm_info(skb, &info) ? TC_ACT_SHOT : TC_ACT_UNSPEC;
}

SEC("tc")
int get_xfrm_info(struct __sk_buff *skb)
{
	struct bpf_xfrm_info info = {};

	if (bpf_skb_get_xfrm_info(skb, &info) < 0)
		return TC_ACT_SHOT;

	resp_if_id = info.if_id;

	return TC_ACT_UNSPEC;
}

char _license[] SEC("license") = "GPL";
```
