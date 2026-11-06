---
title: "KFunc 'scx_bpf_sub_caps'"
description: "This page documents the 'scx_bpf_sub_caps' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_sub_caps`

<!-- [FEATURE_TAG](scx_bpf_sub_caps) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/86094b95efcf747c80930a2a675c5b14505a19b1)
<!-- [/FEATURE_TAG] -->

Read the capabilities held by the calling scheduler or a direct child sub-scheduler.

## Definition

Read the set of cids on which the calling scheduler (`cgroup_id` is `0`), or the direct child sub-scheduler attached to cGroup `cgroup_id`, holds the capabilities `caps`. When multiple capabilities are given, the union of the cid sets of these capabilities is returned. Only the range of cids covered by `out__arena` is read.

This reads the literally granted capabilities, capabilities implied by other capabilities are not included. The read is lock-less, it is a snapshot which may be outdated by concurrent grants and revokes.

**Parameters**

`cgroup_id`: `0` to read the capabilities of the calling scheduler, or the cGroup ID of a direct child sub-scheduler.

`caps`: Bitmask of one or more [`SCX_CAP_*`](scx_bpf_sub_grant.md#capabilities) flags to read.

`out__arena`: Pointer to a [`struct scx_cmask`](scx_bpf_sub_grant.md#struct-scx_cmask) in the [arena](../map-type/BPF_MAP_TYPE_ARENA.md) of the scheduler, which receives the cids on which `caps` are held.

**Returns**

- `0`: Success.
- `-ENODEV`: The calling program is not associated with a scheduler, `cgroup_id` does not name a sub-scheduler, or the capability storage is not initialized yet (for example during `ops.init_cids`).
- `-EOPNOTSUPP`: The calling scheduler is not a cid-form scheduler, or the kernel was built without sub-scheduler support (`CONFIG_EXT_SUB_SCHED`).
- `-EINVAL`: `cgroup_id` names a scheduler that is not a direct child, `caps` contains unknown bits, or `out__arena` is invalid.

The `-EINVAL` errors, the `-ENODEV` error for uninitialized capability storage, and the `-EOPNOTSUPP` error for schedulers that are not cid-form, also abort the calling scheduler with an error.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c s32 scx_bpf_sub_caps(u64 cgroup_id, u64 caps, __arena struct scx_cmask *out__arena)`
<!-- [/KFUNC_DEF] -->

## Usage

See [`scx_bpf_sub_grant`](scx_bpf_sub_grant.md) for an explanation of sub-scheduler capabilities.

A sub-scheduler typically uses this kfunc to find out on which cids it is allowed to run tasks, for example when it is notified of a change via its `ops.sub_caps_updated` callback. Since the notification only summarizes which cids and capabilities changed, but not whether they were granted or revoked, this kfunc is used to query the current state. A parent can also use it to inspect the capabilities it delegated to a direct child.

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

This snippet, from the `scx_qmap` example scheduler, finds out on which cids the scheduler holds `SCX_CAP_ENQ` (exclusive use), and on which it only holds `SCX_CAP_ENQ_IMMED` (shared use).

Example from [`tools/sched_ext/scx_qmap.bpf.c`](https://github.com/torvalds/linux/blob/c03c1a14507dfaf8b856c2608f9ac216021fc222/tools/sched_ext/scx_qmap.bpf.c#L1408-L1411)

!!! example
	```c
	/* find out the cids we hold */
	scx_bpf_sub_caps(0, SCX_CAP_ENQ, &qa.held_excl.mask);
	scx_bpf_sub_caps(0, SCX_CAP_ENQ_IMMED, &qa.held_shared.mask);
	cmask_andnot(&qa.held_shared.mask, &qa.held_excl.mask);	/* held only as ENQ_IMMED */
	```
