# Helper function `bpf_tcp_send_ack`

<!-- [FEATURE_TAG](bpf_tcp_send_ack) -->
[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/206057fe020ac5c037d5e2dd6562a9bd216ec765)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Send out a tcp-ack. _tp_ is the in-kernel struct **tcp_sock**. _rcv_nxt_ is the ack_seq to be sent out.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_tcp_send_ack)(void *tp, __u32 rcv_nxt) = (void *) 116;`
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
