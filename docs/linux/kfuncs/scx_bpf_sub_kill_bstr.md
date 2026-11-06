---
title: "KFunc 'scx_bpf_sub_kill_bstr'"
description: "This page documents the 'scx_bpf_sub_kill_bstr' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_sub_kill_bstr`

<!-- [FEATURE_TAG](scx_bpf_sub_kill_bstr) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/b0a2ca6ae0ed03c625cd343d662d83bc94d84217)
<!-- [/FEATURE_TAG] -->

Kill a direct child sub-scheduler.

## Definition

Evict the direct child sub-scheduler attached to cGroup `cgroup_id`, disabling it with the given reason. The child and its sub-tree are torn down asynchronously through the usual disable path, with exit kind `SCX_EXIT_PARENT_KILL`. The formatted reason is passed to the child's struct [`scx_exit_info`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#struct-scx_exit_info).

Unlike [`scx_bpf_exit_bstr`](scx_bpf_exit_bstr.md), no exit code is taken. The child is a separate scheduler with its own exit code semantics, so an exit code chosen by the parent would have no defined meaning. The reason string carries the intent instead.

Intended to be called through the [`scx_bpf_sub_kill`](https://github.com/torvalds/linux/blob/v7.3/tools/sched_ext/include/scx/common.bpf.h#L179) helper macro.

**Parameters**

`cgroup_id`: cGroup ID of the direct child sub-scheduler to kill.

`fmt`: reason message format string

`data`: format string parameters packaged using [`___bpf_fill`](../../ebpf-library/libbpf/ebpf/___bpf_fill.md) macro

`data__sz`: `data` len, [must end in '__sz' for the verifier](../concepts/kfuncs.md#__sz-annotation)

**Returns**

- `0`: Success.
- `-ENODEV`: The calling program is not associated with a scheduler, or `cgroup_id` does not name a sub-scheduler. The latter can race with the child detaching on its own, so it is not treated as a scheduler error.
- `-EOPNOTSUPP`: The calling scheduler is not a cid-form scheduler, or the kernel was built without sub-scheduler support (`CONFIG_EXT_SUB_SCHED`).
- `-EINVAL`: `cgroup_id` names a scheduler that is not a direct child.

The `-EINVAL` error, and the `-EOPNOTSUPP` error for schedulers that are not cid-form, also abort the calling scheduler with an error.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_sub_kill_bstr(u64 cgroup_id, char *fmt, long long unsigned int *data, u32 data__sz)`
<!-- [/KFUNC_DEF] -->

## Usage

A cid-form parent scheduler can delegate cids to and take them back from its child sub-schedulers with [`scx_bpf_sub_grant`](scx_bpf_sub_grant.md) and [`scx_bpf_sub_revoke`](scx_bpf_sub_revoke.md). This kfunc allows the parent to tear down a child altogether, for example when the parent has no cids left to give to it.

The `scx_bpf_sub_kill` macro from `scx/common.bpf.h` takes a `printf`-style format string with variadic arguments and packs them for this kfunc. When the kernel does not support this kfunc, the macro returns `-EOPNOTSUPP`.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

This snippet, from the `scx_qmap` example scheduler, kills all child sub-schedulers when the scheduler holds no exclusive cids to distribute.

Example from [`tools/sched_ext/scx_qmap.bpf.c`](https://github.com/torvalds/linux/blob/c03c1a14507dfaf8b856c2608f9ac216021fc222/tools/sched_ext/scx_qmap.bpf.c#L1590-L1599)

!!! example
	```c
	/* no excl cpu: run own tasks on the held shares, evict children */
	if (!qa.part.nr_excl) {
		cmask_copy(&qa.self_cids.mask, &qa.held_shared.mask);
		refresh_usable();
		bpf_for(i, 0, MAX_SUB_SCHEDS)
			if (qa.sub_sched_ctxs[i].cgroup_id)
				scx_bpf_sub_kill(qa.sub_sched_ctxs[i].cgroup_id,
						 "parent holds no excl cpu to distribute");
		return;
	}
	```
