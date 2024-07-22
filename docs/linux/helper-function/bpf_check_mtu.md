---
title: "Helper Function 'bpf_check_mtu'"
description: "This page documents the 'bpf_check_mtu' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_check_mtu`

<!-- [FEATURE_TAG](bpf_check_mtu) -->
[:octicons-tag-24: v5.12](https://github.com/torvalds/linux/commit/34b2021cc61642d61c3cf943d9e71925b827941b)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Check packet size against exceeding MTU of net device (based on _ifindex_).  This helper will likely be used in combination with helpers that adjust/change the packet size.

The argument _len_diff_ can be used for querying with a planned size change. This allows to check MTU prior to changing packet ctx. Providing a _len_diff_ adjustment that is larger than the actual packet size (resulting in negative packet size) will in principle not exceed the MTU, which is why it is not considered a failure.  Other BPF helpers are needed for performing the planned size change; therefore the responsibility for catching a negative packet size belongs in those helpers.

Specifying _ifindex_ zero means the MTU check is performed against the current net device.  This is practical if this isn't used prior to redirect.

On input _mtu_len_ must be a valid pointer, else verifier will reject BPF program.  If the value _mtu_len_ is initialized to zero then the ctx packet size is use.  When value _mtu_len_ is provided as input this specify the L3 length that the MTU check is done against. Remember XDP and TC length operate at L2, but this value is L3 as this correlate to MTU and IP-header tot_len values which are L3 (similar behavior as bpf_fib_lookup).

The Linux kernel route table can configure MTUs on a more specific per route level, which is not provided by this helper. For route level MTU checks use the **bpf_fib_lookup**() helper.

_ctx_ is either **struct xdp_md** for XDP programs or **struct sk_buff** for tc cls_act programs.

The _flags_ argument can be a combination of one or more of the following values:

**BPF_MTU_CHK_SEGS**

&nbsp;&nbsp;&nbsp;&nbsp;This flag will only works for _ctx_ **struct sk_buff**. If packet context contains extra packet segment buffers (often knows as GSO skb), then MTU check is harder to check at this point, because in transmit path it is possible for the skb packet to get re-segmented (depending on net device features).  This could still be a MTU violation, so this flag enables performing MTU check against segments, with a different violation return code to tell it apart. Check cannot use len_diff.

On return _mtu_len_ pointer contains the MTU value of the net device.  Remember the net device configured MTU is the L3 size, which is returned here and XDP and TC length operate at L2. Helper take this into account for you, but remember when using MTU value in your BPF-code.



### Returns

* 0 on success, and populate MTU value in _mtu_len_ pointer.


* < 0 if any input argument is invalid (_mtu_len_ not updated)


MTU violations return positive values, but also populate MTU value in _mtu_len_ pointer, as this can be needed for implementing PMTU handing:

* **BPF_MTU_CHK_RET_FRAG_NEEDED**
* **BPF_MTU_CHK_RET_SEGS_TOOBIG**


`#!c static long (* const bpf_check_mtu)(void *ctx, __u32 ifindex, __u32 *mtu_len, __s32 len_diff, __u64 flags) = (void *) 163;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
