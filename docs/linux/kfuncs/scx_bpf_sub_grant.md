---
title: "KFunc 'scx_bpf_sub_grant'"
description: "This page documents the 'scx_bpf_sub_grant' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_sub_grant`

<!-- [FEATURE_TAG](scx_bpf_sub_grant) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/86094b95efcf747c80930a2a675c5b14505a19b1)
<!-- [/FEATURE_TAG] -->

Grant capabilities on a set of cids to a direct child sub-scheduler.

## Definition

Grant the capabilities `caps` on the cids in `cmask__arena` to the direct child sub-scheduler attached to cGroup `cgroup_id`. A cid is only granted to the child if the calling (parent) scheduler holds every requested capability on it, it is all-or-nothing per cid. Cids that are refused are added to `denied_out__arena__nullable` when provided. Refusals for cids outside of the range of the denied mask are not recorded.

**Parameters**

`cgroup_id`: cGroup ID of the direct child sub-scheduler.

`caps`: Bitmask of [`SCX_CAP_*`](#capabilities) flags to grant.

`cmask__arena`: Pointer to a `struct scx_cmask` in the [arena](../map-type/BPF_MAP_TYPE_ARENA.md) of the scheduler, containing the cids to grant `caps` on.

`denied_out__arena__nullable`: Optional pointer to a `struct scx_cmask` in the arena of the scheduler, in which refused cids are accumulated. May be `NULL`.

**Returns**

- `0`: All cids were granted.
- `-EPERM`: One or more cids were refused because the caller does not hold all requested capabilities on them.
- `-ENODEV`: The calling program is not associated with a scheduler, or `cgroup_id` does not name a sub-scheduler.
- `-EOPNOTSUPP`: The calling scheduler is not a cid-form scheduler, or the kernel was built without sub-scheduler support (`CONFIG_EXT_SUB_SCHED`).
- `-EINVAL`: `cgroup_id` names a scheduler that is not a direct child, `caps` contains unknown bits, or a `struct scx_cmask` is invalid.

The `-EINVAL` errors, and the `-EOPNOTSUPP` error for schedulers that are not cid-form, also abort the calling scheduler with an error.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_sub_grant(u64 cgroup_id, u64 caps, __arena const struct scx_cmask *cmask__arena, __arena struct scx_cmask *denied_out__arena__nullable)`
<!-- [/KFUNC_DEF] -->

## Usage

A sched_ext scheduler can be attached as a sub-scheduler on a cGroup (see [`sub_cgroup_id`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#sub_cgroup_id)), creating a hierarchy of schedulers. Capabilities (caps) are per-cid permissions that a parent scheduler delegates to its direct children with this kfunc, and takes away again with [`scx_bpf_sub_revoke`](scx_bpf_sub_revoke.md). A child's caps are always a subset of its parent's caps. Since caps are checked locally by each sub-scheduler, communication between schedulers is only needed when the delegation itself changes.

The sub-scheduler capability kfuncs can only be used by cid-form schedulers. These are schedulers implemented with the `sched_ext_ops_cid` struct ops type, which address CPUs by cid (see [`scx_bpf_this_cid`](scx_bpf_this_cid.md)) and use `struct scx_cmask` cid masks stored in the arena of the scheduler instead of `struct cpumask` CPU masks.

When the caps of a sub-scheduler change, its `ops.sub_caps_updated` callback is invoked with a summary of the cids and caps that changed. Updates are coalesced, the summary does not indicate whether caps were gained or lost, so the callback should use [`scx_bpf_sub_caps`](scx_bpf_sub_caps.md) to query the current state.

### Capabilities

| Capability | Description |
| ---------- | ----------- |
| `SCX_CAP_ENQ_IMMED` | Insert an `SCX_ENQ_IMMED` task onto the local DSQ of the cid and kick the CPU of the cid (except with `SCX_KICK_PREEMPT`). This is the minimal capability needed to make any use of a CPU, also available as `SCX_CAP_BASE`. |
| `SCX_CAP_ENQ` | Insert any task onto the local DSQ of the cid. Implies `SCX_CAP_ENQ_IMMED`. |
| `SCX_CAP_PREEMPT` | Preempt any task running on the cid regardless of the scheduler that owns it, via `SCX_ENQ_PREEMPT` inserts and `SCX_KICK_PREEMPT` kicks. Implies `SCX_CAP_ENQ`. Preempting a task in the own sub-tree of the scheduler doesn't require any capability. |
| `SCX_CAP_PERF` | Control the power/performance state of the CPU of the cid, such as the `cpufreq` target set via [`scx_bpf_cidperf_set`](scx_bpf_cidperf_set.md). Neither implies nor is implied by the other capabilities. |

Implied capabilities only apply to the holder's own use of a cid, not to delegation. This kfunc delegates only literally held capabilities, so a capability held only through implication is usable but can't be granted to a child. When granting a capability, it usually makes sense to also explicitly grant the capabilities it implies.

### `struct scx_cmask`

```c
struct scx_cmask {
	u32 base;
	u32 nr_cids;
	u32 alloc_words;
	u64 bits[];
};
```

A `struct scx_cmask` is a bitmask covering the range of cids `[base, base + nr_cids)`, stored in the arena of the scheduler.

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

This snippet, from the `scx_qmap` example scheduler, moves a child's exclusive cids from the previous to the new set. It revokes all caps on cids the child no longer gets, and grants all caps on newly assigned cids.

Example from [`tools/sched_ext/scx_qmap.bpf.c`](https://github.com/torvalds/linux/blob/c03c1a14507dfaf8b856c2608f9ac216021fc222/tools/sched_ext/scx_qmap.bpf.c#L1646-L1658)

!!! example
	```c
	cmask_copy(&qa.to_revoke_cids.mask, &ssc->prev_granted.mask);
	cmask_andnot(&qa.to_revoke_cids.mask, &ssc->granted_cids.mask);
	cmask_copy(&qa.to_grant_cids.mask, &ssc->granted_cids.mask);
	cmask_andnot(&qa.to_grant_cids.mask, &ssc->prev_granted.mask);

	scx_bpf_sub_revoke(cgid, SCX_CAP_ENQ | SCX_CAP_PREEMPT |
			   SCX_CAP_ENQ_IMMED | SCX_CAP_PERF,
			   &qa.to_revoke_cids.mask);
	scx_bpf_sub_grant(cgid, SCX_CAP_ENQ | SCX_CAP_PREEMPT |
			  SCX_CAP_ENQ_IMMED | SCX_CAP_PERF,
			  &qa.to_grant_cids.mask, NULL);
	```
