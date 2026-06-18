---
title: "KFunc 'scx_bpf_dsq_reenq'"
description: "This page documents the 'scx_bpf_dsq_reenq' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dsq_reenq`

<!-- [FEATURE_TAG](scx_bpf_dsq_reenq) -->
[:octicons-tag-24: 7.1](https://github.com/torvalds/linux/commit/9c34c5074d1bc22072fc7f9c86b0028f7e273b2c)
<!-- [/FEATURE_TAG] -->

Re-enqueue tasks on a DSQ.

## Definition

Iterate over all of the tasks currently enqueued on the DSQ identified by `dsq_id`, and re-enqueue them in the BPF scheduler. 
The following DSQs are supported:

 * Local DSQs (`SCX_DSQ_LOCAL` or `SCX_DSQ_LOCAL_ON | $cpu`)
 * User DSQs

Re-enqueues are performed asynchronously. Can be called from anywhere.

**Parameters**

`dsq_id`: DSQ to re-enqueue.
`reenq_flags`: `SCX_RENQ_*`.
`aux`: implicit BPF argument to access bpf_prog_aux hidden from BPF progs.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dsq_reenq(u64 dsq_id, u64 reenq_flags)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

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

