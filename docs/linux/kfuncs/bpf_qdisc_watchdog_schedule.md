---
title: "KFunc 'bpf_qdisc_watchdog_schedule'"
description: "This page documents the 'bpf_qdisc_watchdog_schedule' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_qdisc_watchdog_schedule`

<!-- [FEATURE_TAG](bpf_qdisc_watchdog_schedule) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/7a2dafda950b78611dc441c83d105dfdc7082681)
<!-- [/FEATURE_TAG] -->

Schedule a qdisc to a later time using a timer.

## Definition

**Parameters**

`sch`: The qdisc to be scheduled.

`expire`: The expiry time of the timer.

`delta_ns`: The slack range of the timer.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_qdisc_watchdog_schedule(struct Qdisc *sch, u64 expire, u64 delta_ns)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is used to signal to a netdev when the first packet in the queue is ready to be sent. This allows the kernel to more efficiently handle qdisc implementations that shape / throttle traffic by delaying when they are sent. Calling this function will hint to the kernel when it should poll for the next packet.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

