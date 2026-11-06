---
title: "KFunc 'scx_bpf_online_cmask'"
description: "This page documents the 'scx_bpf_online_cmask' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_online_cmask`

<!-- [FEATURE_TAG](scx_bpf_online_cmask) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/a9e3760b0838299649c0d57cca44daaf40ba3c33)
<!-- [/FEATURE_TAG] -->

Return the online cid mask in the arena of the scheduler.

## Definition

Return a kernel-maintained [`struct scx_cmask`](scx_bpf_sub_grant.md#struct-scx_cmask) covering all cids, [0, [`scx_bpf_nr_cids`](scx_bpf_nr_cids.md)), in which the bits of online cids are set. The mask lives in the [arena](../map-type/BPF_MAP_TYPE_ARENA.md) of the scheduler.

The mask follows the sched_ext hot-plug notifications: the bit of a cid is updated before the `ops.cid_online` or `ops.cid_offline` callback runs for it. So it only lags behind the online CPU mask during a hot-plug transition. Reads are live, not atomic snapshots, so outside of `ops.init` of the root scheduler, which runs with hot-plug excluded, concurrent updates can be observed.

**Returns**

A pointer to the online cid mask in the arena of the scheduler. Or `NULL` if the calling program is not associated with a live cid-form scheduler, or the mask is not allocated yet, such as during `ops.init_cids`.

The pointer is valid from [`ops.init`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#init) until [`ops.exit`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#exit), no reference needs to be taken. The memory is writable since it is part of the arena, but should be treated as read-only.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c __arena const void *scx_bpf_online_cmask()`
<!-- [/KFUNC_DEF] -->

## Usage

Schedulers using the default cid mapping can treat [0, [`scx_bpf_nr_online_cids`](scx_bpf_nr_online_cids.md)) as the set of online cids, since the scheduler is restarted on hot-plug. Schedulers that install their own mapping with [`scx_bpf_cid_override`](scx_bpf_cid_override.md) to handle hot-plug without a restart have no such guarantee. They can use this kfunc to get the initial set of online cids, and then track changes via the `ops.cid_online` and `ops.cid_offline` callbacks.

The kfunc returns the arena offset as a `void` pointer, which should be cast to an arena pointer. The [`scx/common.bpf.h`](https://github.com/torvalds/linux/blob/v7.3/tools/sched_ext/include/scx/common.bpf.h#L116) header declares the kfunc with a `const void __arena *` return type for this purpose.

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

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
