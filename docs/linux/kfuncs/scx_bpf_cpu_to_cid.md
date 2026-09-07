---
title: "KFunc 'scx_bpf_cpu_to_cid'"
description: "This page documents the 'scx_bpf_cpu_to_cid' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_cpu_to_cid`

<!-- [FEATURE_TAG](scx_bpf_cpu_to_cid) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/e9b55af47edf6e60c9a47b5604c3e15c5162ec86)
<!-- [/FEATURE_TAG] -->

Return the cid for `cpu`.

## Definition

**Parameters**

`cpu`: cpu to look up

**Returns**

Return the cid for `cpu`. Trigger [`scx_error`](https://elixir.bootlin.com/linux/v7.2.2/source/kernel/sched/ext/internal.h#L1545) and return `-EINVAL` if `cpu` is invalid. The cid<->cpu mapping is static for the lifetime of the loaded scheduler, so the BPF side can cache the result to avoid repeated kfunc invocations.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_cpu_to_cid(s32 cpu)`
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

