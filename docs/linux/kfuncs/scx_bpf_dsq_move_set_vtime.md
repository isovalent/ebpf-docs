---
title: "KFunc 'scx_bpf_dsq_move_set_vtime'"
description: "This page documents the 'scx_bpf_dsq_move_set_vtime' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_move_set_vtime`

<!-- [FEATURE_TAG](scx_bpf_dsq_move_set_vtime) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/5cbb302880f50f3edf35f8c6a1d38b6948bf4d11)
<!-- [/FEATURE_TAG] -->

This function overrides `vtime` when moving between DSQs.

## Definition

Override the `vtime` of the next task that will be moved from `it__iter` using [`scx_bpf_dsq_move_vtime`](scx_bpf_dsq_move_vtime.md). If this function is not called, the previous slice vtime is kept. If [`scx_bpf_dsq_move`](scx_bpf_dsq_move.md) is used to dispatch the next task, the override is ignored and cleared.

**Parameters**

`it__iter`: DSQ iterator in progress

`vtime`: task's ordering inside the vtime-sorted queue of the target DSQ


**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dsq_move_set_vtime(struct bpf_iter_scx_dsq *it__iter, u64 vtime)`
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

