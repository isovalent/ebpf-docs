---
title: "KFunc 'scx_bpf_sub_revoke'"
description: "This page documents the 'scx_bpf_sub_revoke' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_sub_revoke`

<!-- [FEATURE_TAG](scx_bpf_sub_revoke) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/86094b95efcf747c80930a2a675c5b14505a19b1)
<!-- [/FEATURE_TAG] -->

Revoke capabilities on a set of cids from a direct child sub-scheduler.

## Definition

Clear the capabilities `caps` on the cids in `cmask__arena` from the direct child sub-scheduler attached to cGroup `cgroup_id`, and from all of its descendants. Since a child's capabilities must always be a subset of its parent's, revoking a capability from a child also revokes it from the entire sub-tree below that child.

**Parameters**

`cgroup_id`: cGroup ID of the direct child sub-scheduler.

`caps`: Bitmask of [`SCX_CAP_*`](scx_bpf_sub_grant.md#capabilities) flags to revoke.

`cmask__arena`: Pointer to a [`struct scx_cmask`](scx_bpf_sub_grant.md#struct-scx_cmask) in the [arena](../map-type/BPF_MAP_TYPE_ARENA.md) of the scheduler, containing the cids to revoke `caps` on.

**Returns**

This kfunc does not return a value. If the calling scheduler is not a cid-form scheduler, `cgroup_id` names a scheduler that is not a direct child, `caps` contains unknown bits or `cmask__arena` is invalid, the calling scheduler is aborted with an error. If `cgroup_id` does not name a sub-scheduler, or the kernel was built without sub-scheduler support, nothing happens.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_sub_revoke(u64 cgroup_id, u64 caps, __arena const struct scx_cmask *cmask__arena)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is the counterpart of [`scx_bpf_sub_grant`](scx_bpf_sub_grant.md), see its page for an explanation of sub-scheduler capabilities. A parent scheduler uses it to take back cids it previously delegated to a child, for example to give them to another child.

Every sub-scheduler in the sub-tree that lost capabilities is notified via its `ops.sub_caps_updated` callback. Tasks that were queued on, or running on, a cid that the scheduler can no longer use because it lost `SCX_CAP_ENQ` or `SCX_CAP_ENQ_IMMED`, are re-enqueued.

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

See [`scx_bpf_sub_grant`](scx_bpf_sub_grant.md#example) for an example.
