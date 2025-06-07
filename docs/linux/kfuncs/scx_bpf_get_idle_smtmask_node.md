---
title: "KFunc 'scx_bpf_get_idle_smtmask_node'"
description: "This page documents the 'scx_bpf_get_idle_smtmask_node' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_get_idle_smtmask_node`

<!-- [FEATURE_TAG](scx_bpf_get_idle_smtmask_node) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/01059219b0cfdb9fc0d5bd60458e614a3135e6e7)
<!-- [/FEATURE_TAG] -->

Get a referenced kptr to the idle-tracking, per-physical-core cpumask of a target NUMA node.

## Definition

Can be used to determine if an entire physical core is free.

**Parameters**

`node`: target NUMA node

**Returns**

An empty cpumask if idle tracking is not enabled, if `node` is not valid, or running on a `UP` kernel. In this case the actual error will be reported to the BPF scheduler via `scx_ops_error`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c const struct cpumask *scx_bpf_get_idle_smtmask_node(int node)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.
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

