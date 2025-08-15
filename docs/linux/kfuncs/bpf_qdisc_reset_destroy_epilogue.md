---
title: "KFunc 'bpf_qdisc_reset_destroy_epilogue'"
description: "This page documents the 'bpf_qdisc_reset_destroy_epilogue' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_qdisc_reset_destroy_epilogue`

<!-- [FEATURE_TAG](bpf_qdisc_reset_destroy_epilogue) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/7a2dafda950b78611dc441c83d105dfdc7082681)
<!-- [/FEATURE_TAG] -->

A hidden / internal kfunc called by the function epilogue of `reset` and `destroy` qdisc struct ops programs.

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_qdisc_reset_destroy_epilogue(struct Qdisc *sch)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc cannot be called by users directly, it is implicitly called on function exit.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

