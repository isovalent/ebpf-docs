---
title: "Helper Function 'bpf_ktime_get_coarse_ns'"
description: "This page documents the 'bpf_ktime_get_coarse_ns' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_ktime_get_coarse_ns`

<!-- [FEATURE_TAG](bpf_ktime_get_coarse_ns) -->
[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/d055126180564a57fe533728a4e93d0cb53d49b3)
<!-- [/FEATURE_TAG] -->
## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return a coarse-grained version of the time elapsed since system boot, in nanoseconds. Does not include time the system was suspended.

See: **clock_gettime**(**CLOCK_MONOTONIC_COARSE**)

### Returns

Current _ktime_.

`#!c static __u64 (* const bpf_ktime_get_coarse_ns)(void) = (void *) 160;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

This function returns a coarse-grained 64-bit timestamp in nanoseconds since system boot, excluding suspended time. It is similar to `bpf_ktime_get_ns()`, but offers lower precision in exchange for better performance. It is suitable for low-overhead time measurements.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->

<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```
__u64 start_time = bpf_ktime_get_coarse_ns();
/* some tasks */
__u64 end_time = bpf_ktime_get_coarse_ns();
__u64 duration = end_time - start_time;
```