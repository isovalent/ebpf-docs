---
title: "KFunc 'cgroup_rstat_updated'"
description: "This page documents the 'cgroup_rstat_updated' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `cgroup_rstat_updated`

<!-- [FEATURE_TAG](cgroup_rstat_updated) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/a319185be9f5ad13c2a296d448ac52ffe45d194c)
<!-- [/FEATURE_TAG] -->

Keep track of updated `rstat_cpu`

## Definition

`cgrp`'s rstat_cpu on `cpu` was updated. Put it on the parent's matching `rstat_cpu->updated_children` list. See the comment on top of `cgroup_rstat_cpu` definition for details.


**Parameters**

`cgrp`: target cgroup

`cpu`: cpu on which rstat_cpu was updated

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void cgroup_rstat_updated(struct cgroup *cgrp, int cpu)`
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

