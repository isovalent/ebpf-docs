---
title: "Helper Function 'bpf_perf_event_output'"
description: "This page documents the 'bpf_perf_event_output' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_perf_event_output`

<!-- [FEATURE_TAG](bpf_perf_event_output) -->
[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/a43eec304259a6c637f4014a6d4767159b6a3aa3)
<!-- [/FEATURE_TAG] -->

This helper writes a raw `data` blob into a special BPF perf event held by `map` of type [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (*bpf_perf_event_output)(void *ctx, void *map, __u64 flags, void *data, __u64 size) = (void *) 25;`

## Usage

The perf event in `map` must have the following attributes: `PERF_SAMPLE_RAW` as `sample_type`, `PERF_TYPE_SOFTWARE` as `type`, and `PERF_COUNT_SW_BPF_OUTPUT` as `config`.

The `flags` are used to indicate the index in `map` for which the value must be put, masked with `BPF_F_INDEX_MASK`. Alternatively, `flags` can be set to `BPF_F_CURRENT_CPU` to indicate that the index of the current CPU core should be used.

!!! warning
    Each perf event is created on a specific CPU. This helper can only write to perf events on the same CPU as the eBPF program is running. Manually picking an index containing a perf event on a different CPU will result in a `-EOPNOTSUPP` error at runtime. So unless there is a good reason to do so, its recommended to use `BPF_F_CURRENT_CPU` and to populate the [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) map in such a way where the CPU indices and map indices are the same.

The value to write, of `size`, is passed through eBPF stack and pointed by `data`.

The context of the program `ctx` needs also be passed to the helper.

On user space, a program willing to read the values needs to call [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) on the perf event (either for one or for all CPUs) and to store the file descriptor into the `map`. This must be done before the eBPF program can send data into it. An example is available in file [`samples/bpf/trace_output_user.c`](https://github.com/torvalds/linux/blob/v6.2/samples/bpf/trace_output_user.c) in the Linux kernel source tree (the eBPF program counterpart is in [`samples/bpf/trace_output_kern.c`](https://github.com/torvalds/linux/blob/v6.2/samples/bpf/trace_output_kern.c)).

`bpf_perf_event_output` achieves better performance than [`bpf_trace_printk`](bpf_trace_printk.md) for sharing data with user space, and is much better suitable for streaming data from eBPF programs.

Note that this helper is not restricted to tracing use cases and can be used with programs attached to TC or XDP as well, where it allows for passing data to user space listeners. Data can be:

* Only custom structs,
* Only the packet payload, or
* A combination of both.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_DEVICE](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [BPF_MAP_TYPE_PERF_EVENT_ARRAY](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
