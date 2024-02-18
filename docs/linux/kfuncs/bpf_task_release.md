# KFunc `bpf_task_release`

<!-- [FEATURE_TAG](bpf_task_release) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/3f00c52393445ed49aadc1a567aa502c6333b1a1)
<!-- [/FEATURE_TAG] -->

Release the reference acquired on a task.

## Definition

`p`: The task on which a reference is being released.

<!-- [KFUNC_DEF] -->
`#!c void bpf_task_release(struct task_struct *p)`

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
- [BPF_PROG_TYPE_LSM](../../program-types/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_SCHED_CLS](../../program-types/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../../program-types/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_TRACING](../../program-types/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../../program-types/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

