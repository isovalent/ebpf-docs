---
title: "KFunc 'scx_bpf_dump_bstr'"
description: "This page documents the 'scx_bpf_dump_bstr' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_dump_bstr`

<!-- [FEATURE_TAG](scx_bpf_dump_bstr) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/07814a9439a3b03d79a1001614b5bc1cab69bcec)
<!-- [/FEATURE_TAG] -->

This function generates extra debug dump specific to the BPF scheduler.

## Definition

To be called through [`scx_bpf_dump`](../../ebpf-library/scx/scx_bpf_dump.md) helper from [`sched_ext_ops.dump`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dump), [`sched_ext_ops.dump_cpu`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dump_cpu) and [`sched_ext_ops.dump_task`](../program-type/BPF_PROG_TYPE_STRUCT_OPS/sched_ext_ops.md#dump_task) to generate extra debug dump specific to the BPF scheduler.

The extra dump may be multiple lines. A single line may be split over multiple calls. The last line is automatically terminated.

**Parameters**

`fmt`: format string

`data`: format string parameters packaged using [`___bpf_fill`](../../ebpf-library/libbpf/ebpf/___bpf_fill.md) macro

`data__sz`: `data` len, [must end in '__sz' for the verifier](../concepts/kfuncs.md#__sz-annotation)



**Signature**

<!-- [KFUNC_DEF] -->
`#!c void scx_bpf_dump_bstr(char *fmt, long long unsigned int *data, u32 data__sz)`
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

