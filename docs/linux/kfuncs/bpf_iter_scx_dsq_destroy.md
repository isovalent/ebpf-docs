---
title: "KFunc 'bpf_iter_scx_dsq_destroy'"
description: "This page documents the 'bpf_iter_scx_dsq_destroy' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_scx_dsq_destroy`

<!-- [FEATURE_TAG](bpf_iter_scx_dsq_destroy) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/650ba21b131ed1f8ee57826b2c6295a3be221132)
<!-- [/FEATURE_TAG] -->

This function destroys a DSQ iterator.

## Definition

Undo [`bpf_iter_scx_dsq_new`](bpf_iter_scx_dsq_new.md).

**Parameters**

`it`: iterator to destroy

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_iter_scx_dsq_destroy(struct bpf_iter_scx_dsq *it)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

