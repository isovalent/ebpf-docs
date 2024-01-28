# Helper function `bpf_skb_set_tstamp`

<!-- [FEATURE_TAG](bpf_skb_set_tstamp) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/9bb984f28d5bcb917d35d930fcfb89f90f9449fd)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Change the __sk_buff->tstamp_type to _tstamp_type_ and set _tstamp_ to the __sk_buff->tstamp together.

If there is no need to change the __sk_buff->tstamp_type, the tstamp value can be directly written to __sk_buff->tstamp instead.

BPF_SKB_TSTAMP_DELIVERY_MONO is the only tstamp that will be kept during bpf_redirect__().  A non zero _tstamp_ must be used with the BPF_SKB_TSTAMP_DELIVERY_MONO _tstamp_type_.

A BPF_SKB_TSTAMP_UNSPEC _tstamp_type_ can only be used with a zero _tstamp_.

Only IPv4 and IPv6 skb->protocol are supported.

This function is most useful when it needs to set a mono delivery time to __sk_buff->tstamp and then bpf_redirect__() to the egress of an iface.  For example, changing the (rcv) timestamp in __sk_buff->tstamp at ingress to a mono delivery time and then bpf_redirect__() to sch_fq@phy-dev.

### Returns

0 on success. **-EINVAL** for invalid input **-EOPNOTSUPP** for unsupported protocol

`#!c static long (*bpf_skb_set_tstamp)(struct __sk_buff *skb, __u64 tstamp, __u32 tstamp_type) = (void *) 192;`
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
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
