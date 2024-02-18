# KFunc `bpf_cgroup_release`

<!-- [FEATURE_TAG](bpf_cgroup_release) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/fda01efc61605af7c6fa03c4109f14d59c9228b7)
<!-- [/FEATURE_TAG] -->

Release the reference acquired on a cgroup.

## Definition

If this kfunc is invoked in an RCU read region, the cgroup is guaranteed to not be freed until the current grace period has ended, even if its refcount drops to `0`.

<!-- [KFUNC_DEF] -->
`#!c void bpf_cgroup_release(struct cgroup *cgrp)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

