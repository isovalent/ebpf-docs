---
title: "KFunc 'scx_bpf_now'"
description: "This page documents the 'scx_bpf_now' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `scx_bpf_now`

<!-- [FEATURE_TAG](scx_bpf_now) -->
[:octicons-tag-24: v6.14](https://github.com/torvalds/linux/commit/3a9910b5904d29c566e3ff9290990b519827ba75)
<!-- [/FEATURE_TAG] -->

This function returns a high-performance monotonically non-decreasing clock for the current CPU. The clock returned is in nanoseconds.

## Definition

It provides the following properties:

1) High performance: Many BPF schedulers call [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md) frequently to account for execution time and track tasks' runtime properties. Unfortunately, in some hardware platforms, [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md) -- which eventually reads a hardware timestamp counter -- is neither performant nor scalable. `scx_bpf_now` aims to provide a high-performance clock by using the rq clock in the scheduler core whenever possible.

2) High enough resolution for the BPF scheduler use cases: In most BPF scheduler use cases, the required clock resolution is lower than the most accurate hardware clock (e.g., rdtsc in x86). `scx_bpf_now` basically uses the rq clock in the scheduler core whenever it is valid. It considers that the rq clock is valid from the time the rq clock is updated (`update_rq_clock`) until the rq is unlocked (`rq_unpin_lock`).

3) Monotonically non-decreasing clock for the same CPU: `scx_bpf_now` guarantees the clock never goes backward when comparing them in the same CPU. On the other hand, when comparing clocks in different CPUs, there is no such guarantee -- the clock can go backward. It provides a monotonically *non-decreasing* clock so that it would provide the same clock values in two different `scx_bpf_now` calls in the same CPU during the same period of when the rq clock is valid.

**Returns**

A high-performance monotonically non-decreasing clock for the current CPU. The clock returned is in nanoseconds.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c u64 scx_bpf_now()`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

