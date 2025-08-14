---
title: "KFunc 'css_rstat_flush'"
description: "This page documents the 'css_rstat_flush' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `css_rstat_flush`

<!-- [FEATURE_TAG](css_rstat_flush) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/a319185be9f5ad13c2a296d448ac52ffe45d194c)
<!-- [/FEATURE_TAG] -->

Flush stats in `css`'s subtree

## Definition

Collect all per-CPU stats in `css->cgroup`'s sub-tree into the global countersand propagate them upwards. After this function returns, all cGroups in the sub-tree have up-to-date `->stat`.

This also gets all cGroups in the sub-tree including `css->cgroup` off the `->updated_children` lists.

This function may block.

**Parameters**

`css`: target cgroup subsystem state

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void css_rstat_flush(struct cgroup_subsys_state *css)`

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
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

