---
title: "KFunc 'scx_bpf_cid_topo'"
description: "This page documents the 'scx_bpf_cid_topo' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cid_topo`

<!-- [FEATURE_TAG](scx_bpf_cid_topo) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/e9b55af47edf6e60c9a47b5604c3e15c5162ec86)
<!-- [/FEATURE_TAG] -->

Copy out per-cid topology info.

## Definition

Fill `out__uninit` with the topology info for `cid`. Trigger [`scx_error`](https://elixir.bootlin.com/linux/v7.2.2/source/kernel/sched/ext/internal.h#L1545) if `cid` is out of range. If `cid` is valid but in the no-topo section, all fields are set to `-1`.

**Parameters**

`cid`: cid to look up

`out_uninit`: where to copy the topology info; fully written by this call

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_cid_topo(s32 cid, struct scx_cid_topo *out__uninit)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

