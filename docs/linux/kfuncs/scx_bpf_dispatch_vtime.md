---
title: "KFunc 'scx_bpf_dispatch_vtime'"
description: "This page documents the 'scx_bpf_dispatch_vtime' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dispatch_vtime`

<!-- [FEATURE_TAG](scx_bpf_dispatch_vtime) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/06e51be3d5e7a07aea5c9012773df8d5de01db6c)
<!-- [/FEATURE_TAG] -->

This function was renamed to [`scx_bpf_dsq_insert_vtime`](scx_bpf_dsq_insert_vtime.md) in [:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/cc26abb1a19adbb91b79d25a2e74976633ece429). But will be aliased until v6.17.

!!! warning
    The alias was removed in [v6.17](https://github.com/torvalds/linux/commit/4ecf83741401c70d4420588ee1f3b1ca04ef58d5), and is no longer available.

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dispatch_vtime(struct task_struct *p, u64 dsq_id, u64 slice, u64 vtime, u64 enq_flags)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

