# KFunc `cgroup_rstat_flush`

<!-- [FEATURE_TAG](cgroup_rstat_flush) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/a319185be9f5ad13c2a296d448ac52ffe45d194c)
<!-- [/FEATURE_TAG] -->

Flush stats in `cgrp`'s subtree

## Definition

Collect all per-cpu stats in `cgrp`'s subtree into the global countersand propagate them upwards. After this function returns, all cgroups in the subtree have up-to-date ->stat.

This also gets all cgroups in the subtree including `cgrp` off the ->updated_children lists.

This function may block.

**Parameters**

`cgrp`: target cgroup

<!-- [KFUNC_DEF] -->
`#!c void cgroup_rstat_flush(struct cgroup *cgrp)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../../syscall/BPF_PROG_LOAD/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

