# SCX

The kernel concludes a number of [example sched_ext BPF schedulers](https://github.com/torvalds/linux/tree/master/tools/sched_ext). These examples make use of some common header files that aid in writing schedulers. These are located in [`tools/sched_ext/include/scx`](https://github.com/torvalds/linux/tree/master/tools/sched_ext/include/scx). We document macros inside of these to make following examples easier and because kfuncs and the [sched_ext_ops](../../linux/program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md) reference them.

!!! Note
    These in-tree headers are based on files developed in an out of tree [sched-ext/scx](https://github.com/sched-ext/scx) project, which moves faster than the kernel. If you are looking for a library to use in your own projects, you may want to look at the out of tree project.

## [`common.bpf.h`](https://github.com/torvalds/linux/tree/master/tools/sched_ext/include/scx/common.bpf.h)

This header file provides a number of macros that might be useful when writing schedulers. It also contains forward declarations for most kfuncs that might be of use.

Definitions in this file are:

* [`BPF_FOR_EACH_ITER`](BPF_FOR_EACH_ITER.md)
* [`scx_bpf_bstr_preamble`](scx_bpf_bstr_preamble.md)
* [`scx_bpf_exit`](scx_bpf_exit.md)
* [`scx_bpf_error`](scx_bpf_error.md)
* [`scx_bpf_dump`](scx_bpf_dump.md)
* [`BPF_STRUCT_OPS`](BPF_STRUCT_OPS.md)
* [`BPF_STRUCT_OPS_SLEEPABLE`](BPF_STRUCT_OPS_SLEEPABLE.md)
* [`RESIZABLE_ARRAY`](RESIZABLE_ARRAY.md)
* [`ARRAY_ELEM_PTR`](ARRAY_ELEM_PTR.md)
* [`MEMBER_VPTR`](MEMBER_VPTR.md)
* [`__contains`](__contains.md)
* [`private`](private.md)
* [`bpf_obj_new`](bpf_obj_new.md)
* [`bpf_obj_drop`](bpf_obj_drop.md)
* [`bpf_rbtree_add`](bpf_rbtree_add.md)
* [`bpf_refcount_acquire`](bpf_refcount_acquire.md)
* [`cast_mask`](cast_mask.md)
* [`likely`](likely.md)
* [`unlikely`](unlikely.md)
* [`READ_ONCE`](READ_ONCE.md)
* [`WRITE_ONCE`](WRITE_ONCE.md)
* [`log2_u32`](log2_u32.md)
* [`log2_u64`](log2_u64.md)

## [`compat.bpf.h`](https://github.com/torvalds/linux/tree/master/tools/sched_ext/include/scx/compat.bpf.h)

This header file provides compatibility wrappers that pick the right implementation based on the kernel version, dealing with the fact that the sched_ext implementation is still in flux.

Definitions in this file are:

* [`__COMPAT_ENUM_OR_ZERO`](__COMPAT_ENUM_OR_ZERO.md)
* [`__COMPAT_scx_bpf_task_cgroup`](__COMPAT_scx_bpf_task_cgroup.md)
* [`scx_bpf_dsq_insert`](scx_bpf_dsq_insert.md)
* [`scx_bpf_dsq_insert_vtime`](scx_bpf_dsq_insert_vtime.md)
* [`scx_bpf_dsq_move_to_local`](scx_bpf_dsq_move_to_local.md)
* [`__COMPAT_scx_bpf_dsq_move_set_slice`](__COMPAT_scx_bpf_dsq_move_set_slice.md)
* [`__COMPAT_scx_bpf_dsq_move_set_vtime`](__COMPAT_scx_bpf_dsq_move_set_vtime.md)
* [`__COMPAT_scx_bpf_dsq_move`](__COMPAT_scx_bpf_dsq_move.md)
* [`__COMPAT_scx_bpf_dsq_move_vtime`](__COMPAT_scx_bpf_dsq_move_vtime.md)
* [`SCX_OPS_DEFINE`](SCX_OPS_DEFINE.md)
