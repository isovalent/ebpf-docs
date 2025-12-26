---
title: "KFunc 'bpf_stream_vprintk_impl'"
description: "This page documents the 'bpf_stream_vprintk_impl' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_stream_vprintk_impl`

<!-- [FEATURE_TAG](bpf_stream_vprintk_impl) -->
[:octicons-tag-24: v6.17](https://github.com/torvalds/linux/commit/5ab154f1463a111e1dc8fd5d31eaa7a2a71fe2e6)
<!-- [/FEATURE_TAG] -->

Write a message to a <nospell>stdout/stderr</nospell> like, per-program stream.

!!! note
    In [:octicons-tag-24: v6.17](https://github.com/torvalds/linux/commit/5ab154f1463a111e1dc8fd5d31eaa7a2a71fe2e6) this kfunc was introduced as `bpf_stream_vprintk` and was renamed to `bpf_stream_vprintk_impl` in [:octicons-tag-24: v6.18](https://github.com/torvalds/linux/commit/137cc92ffe2e71705fce112656a460d924934ebe) 

## Definition

**Parameters**

`stream_id`: The ID of the stream to write to (`BPF_STDOUT`(1), or `BPF_STDERR`(2))

`fmt__str`: Format string following the same formatting as [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)

`args`: Pointer to an array of u64 argument values.

`len_sz`: Number of elements in `args`.

`aux__prog`: [Pseudo argument](../concepts/kfuncs.md#__prog-annotation), any value passed in is ignored.

**Returns**

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_stream_vprintk_impl(int stream_id, const char *fmt__str, const void *args, u32 len__sz, void *aux__prog)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc allows programs to write formatted strings to streams, the API is made to resemble <nospell>stdout/stderr</nospell> streams normal userspace programs have access to. 

These streams are intended for debugging, an alternative to [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md) and [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md) which write to a system wide trace log. These streams are per-program, thus making it easier to consume logs from a specific program instead of having to figure out which log message in the system wide log is produces my the program you are interested in.

bpftool can be used to inspect these streams with the following command: `bpftool prog tracelog { stdout | stderr } *PROG*`

The libbpf eBPF side library defines a helper macro [`bpf_stream_printk`](../../ebpf-library/libbpf/ebpf/bpf_stream_printk.md) which makes using this kfunc easier.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
- [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

