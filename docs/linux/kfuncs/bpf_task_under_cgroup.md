---
title: "KFunc 'bpf_task_under_cgroup'"
description: "This page documents the 'bpf_task_under_cgroup' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_task_under_cgroup`

<!-- [FEATURE_TAG](bpf_task_under_cgroup) -->
[:octicons-tag-24: v6.5](https://github.com/torvalds/linux/commit/b5ad4cdc46c7d6e7f8d2c9e24b6c9a1edec95154)
<!-- [/FEATURE_TAG] -->

Wrap `task_under_cgroup_hierarchy()` as a kfunc, test task's membership of cGroup ancestry.

## Definition

Tests whether `task`'s default cgroup hierarchy is a descendant of `ancestor`. It follows all the same rules as cgroup_is_descendant, and only applies to the default hierarchy.

<!-- [KFUNC_DEF] -->
`#!c long int bpf_task_under_cgroup(struct task_struct *task, struct cgroup *ancestor)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

