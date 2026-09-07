---
title: "KFunc 'bpf_mem_cgroup_page_state'"
description: "This page documents the 'bpf_mem_cgroup_page_state' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_mem_cgroup_page_state`

<!-- [FEATURE_TAG](bpf_mem_cgroup_page_state) -->
[:octicons-tag-24: v7.0](https://github.com/torvalds/linux/commit/99430ab8b804c26b8a0dec93fcbfe75469f3edc7)
<!-- [/FEATURE_TAG] -->

Flush memory cGroup's statistics.

## Definition

Propagate memory cGroup's statistics up the cGroup tree.

**Parameters**

`memcg`: memory cGroup

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_mem_cgroup_flush_stats(struct mem_cgroup *memcg)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

