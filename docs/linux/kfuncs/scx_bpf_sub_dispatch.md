---
title: "KFunc 'scx_bpf_sub_dispatch'"
description: "This page documents the 'scx_bpf_sub_dispatch' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_sub_dispatch`

<!-- [FEATURE_TAG](scx_bpf_sub_dispatch) -->
[:octicons-tag-24: 7.1](https://github.com/torvalds/linux/commit/4f8b122848dbc353a193de0fa707bc40b5f067ff)
<!-- [/FEATURE_TAG] -->

Trigger dispatching on a child scheduler.

## Definition

Allows a parent scheduler to trigger dispatching on one of its direct child schedulers. The child scheduler runs its dispatch operation to move tasks from dispatch queues to the local run-queue.

**Parameters**

`cgroup_id`: cgroup ID of the child scheduler to dispatch.
`aux`: implicit BPF argument to access bpf_prog_aux hidden from BPF progs.

**Returns**

`true` on success, `false` if `cgroup_id` is invalid, not a direct child, or caller lacks dispatch permission.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c bool scx_bpf_sub_dispatch(u64 cgroup_id)`
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

