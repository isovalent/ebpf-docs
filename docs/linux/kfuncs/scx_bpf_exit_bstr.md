---
title: "KFunc 'scx_bpf_exit_bstr'"
description: "This page documents the 'scx_bpf_exit_bstr' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_exit_bstr`

<!-- [FEATURE_TAG](scx_bpf_exit_bstr) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
<!-- [/FEATURE_TAG] -->

This function gracefully exits the BPF scheduler.

## Definition

Indicate that the BPF scheduler wants to exit gracefully, and initiate ops disabling. Intended to be called through the [`scx_bpf_exit`](https://github.com/torvalds/linux/blob/2a52ca7c98960aafb0eca9ef96b2d0c932171357/tools/sched_ext/include/scx/common.bpf.h#L80) helper macro.

`exit_code`: Exit value to pass to user space via struct `scx_exit_info`.

`fmt`: error message format string

`data`: format string parameters packaged using [`___bpf_fill`](../../ebpf-library/libbpf/ebpf/___bpf_fill.md) macro

`data__sz`: `data` len, [must end in '__sz' for the verifier](../concepts/kfuncs.md#__sz-annotation)

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_exit_bstr(s64 exit_code, char *fmt, long long unsigned int *data, u32 data__sz)`
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

