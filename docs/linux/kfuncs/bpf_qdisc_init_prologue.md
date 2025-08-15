---
title: "KFunc 'bpf_qdisc_init_prologue'"
description: "This page documents the 'bpf_qdisc_init_prologue' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_qdisc_init_prologue`

<!-- [FEATURE_TAG](bpf_qdisc_init_prologue) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/7a2dafda950b78611dc441c83d105dfdc7082681)
<!-- [/FEATURE_TAG] -->

A hidden / internal kfunc called by the function prolog of `init` qdisc struct ops programs.

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_qdisc_init_prologue(struct Qdisc *sch, struct netlink_ext_ack *extack)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc cannot be called by users directly, it is implicitly called on function entry.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

