---
title: "KFunc 'scx_bpf_dispatch_from_dsq_set_slice'"
description: "This page documents the 'scx_bpf_dispatch_from_dsq_set_slice' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dispatch_from_dsq_set_slice`

<!-- [FEATURE_TAG](scx_bpf_dispatch_from_dsq_set_slice) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/4c30f5ce4f7af4f639af99e0bdeada8b268b7361)
<!-- [/FEATURE_TAG] -->

This function was renamed to [`scx_bpf_dsq_move_set_slice`](scx_bpf_dsq_move_set_slice.md) in [:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/5cbb302880f50f3edf35f8c6a1d38b6948bf4d11). But will be aliased until v6.17.

!!! warning
    The alias was removed in [v6.17](https://github.com/torvalds/linux/commit/4ecf83741401c70d4420588ee1f3b1ca04ef58d5), and is no longer available.

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dispatch_from_dsq_set_slice(struct bpf_iter_scx_dsq *it__iter, u64 slice)`
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

