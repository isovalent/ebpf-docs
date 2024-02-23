---
title: "KFunc 'bpf_skb_get_xfrm_info'"
description: "This page documents the 'bpf_skb_get_xfrm_info' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_skb_get_xfrm_info`

<!-- [FEATURE_TAG](bpf_skb_get_xfrm_info) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/94151f5aa9667c562281abeaaa5e89b9d5c17729)
<!-- [/FEATURE_TAG] -->

Get XFRM metadata

## Definition

**Parameters**

`skb_ctx`: Pointer to ctx (__sk_buff) in TC program. Cannot be NULL

`to`: Pointer to memory to which the metadata will be copied. Cannot be NULL

**Members**

`to.if_id`: XFRM if_id:

- Transmit: if_id to be used in policy and state lookups
- Receive: if_id of the state matched for the incoming packet

`to.link`: Underlying device ifindex:

- Transmit: used as the underlying device in VRF routing
- Receive: the device on which the packet had been received

<!-- [KFUNC_DEF] -->
`#!c int bpf_skb_get_xfrm_info(struct __sk_buff *skb_ctx, struct bpf_xfrm_info *to)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
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
