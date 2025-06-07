---
title: "KFunc 'bpf_sock_ops_enable_tx_tstamp'"
description: "This page documents the 'bpf_sock_ops_enable_tx_tstamp' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_sock_ops_enable_tx_tstamp`

<!-- [FEATURE_TAG](bpf_sock_ops_enable_tx_tstamp) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00)
<!-- [/FEATURE_TAG] -->

Enable sock ops TX time-stamping callbacks for a given socket.

## Definition

This kfunc enabled the [`BPF_SOCK_OPS_TSTAMP_SCHED_CB`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md#bpf_sock_ops_tstamp_sched_cb), [`BPF_SOCK_OPS_TSTAMP_SND_SW_CB`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md#bpf_sock_ops_tstamp_snd_sw_cb), [`BPF_SOCK_OPS_TSTAMP_SND_HW_CB`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md#bpf_sock_ops_tstamp_snd_hw_cb) and [`BPF_SOCK_OPS_TSTAMP_ACK_CB`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md#bpf_sock_ops_tstamp_ack_cb) 

Can only be called from [`BPF_SOCK_OPS_TSTAMP_SENDMSG_CB`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md#bpf_sock_ops_tstamp_sendmsg_cb)

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_sock_ops_enable_tx_tstamp(struct bpf_sock_ops_kern *skops, u64 flags)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
- [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
- [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
- [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

