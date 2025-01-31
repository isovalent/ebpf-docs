---
title: "KFunc 'scx_bpf_dsq_move_set_slice'"
description: "This page documents the 'scx_bpf_dsq_move_set_slice' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_move_set_slice`

<!-- [FEATURE_TAG](scx_bpf_dsq_move_set_slice) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/5cbb302880f50f3edf35f8c6a1d38b6948bf4d11)
<!-- [/FEATURE_TAG] -->

This function overrides slice when moving between DSQs.

## Definition

Override the slice of the next task that will be moved from `it__iter` using [`scx_bpf_dsq_move`](scx_bpf_dsq_move.md) or [`scx_bpf_dsq_move_vtime`](scx_bpf_dsq_move_vtime.md). If this function is not called, the previous slice duration is kept.

**Parameters**

`it__iter`: DSQ iterator in progress

`slice`: duration the moved task can run for in nanoseconds


**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dsq_move_set_slice(struct bpf_iter_scx_dsq *it__iter, u64 slice)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

