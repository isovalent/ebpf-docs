---
title: "Helper Function 'bpf_perf_event_read_value'"
description: "This page documents the 'bpf_perf_event_read_value' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_perf_event_read_value`

<!-- [FEATURE_TAG](bpf_perf_event_read_value) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/908432ca84fc229e906ba164219e9ad0fe56f755)
<!-- [/FEATURE_TAG] -->

This helper function reads the value of a perf event counter, and store it into `buf` of size `buf_size`.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_perf_event_read_value)(void *map, __u64 flags, struct bpf_perf_event_value *buf, __u32 buf_size) = (void *) 55;`

## Usage

This helper relies on a `map` of type [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md). The nature of the perf event counter is selected when `map` is updated with perf event file descriptors. The `map` is an array whose size is the number of available CPUs, and each cell contains a value relative to one CPU. The value to retrieve is indicated by `flags`, that contains the index of the CPU to look up, masked with `BPF_F_INDEX_MASK`. Alternatively, `flags` can be set to `BPF_F_CURRENT_CPU` to indicate that the value for the current CPU should be retrieved.

This helper behaves in a way close to [`bpf_perf_event_read`](bpf_perf_event_read.md) helper, save that instead of just returning the value observed, it fills the `buf` structure. This allows for additional data to be retrieved: in particular, the enabled and running times (in `buf->enabled` and `buf->running`, respectively) are copied. In general, [`bpf_perf_event_read_value`](bpf_perf_event_read_value.md) is recommended over [`bpf_perf_event_read`](bpf_perf_event_read.md), which has some ABI issues and provides fewer functionalities.

These values are interesting, because hardware PMU (Performance Monitoring Unit) counters are limited resources. When there are more PMU based perf events opened than available counters, kernel will multiplex these events so each event gets certain percentage (but not all) of the PMU time. In case that multiplexing happens, the number of samples or counter value will not reflect the case compared to when no multiplexing occurs. This makes comparison between different runs difficult. Typically, the counter value should be normalized before comparing to other experiments. The usual normalization is done as follows.

`#!c normalized_counter = counter * t_enabled / t_running`

Where t_enabled is the time enabled for event and t_running is the time running for event since last normalization. The enabled and running times are accumulated since the perf event open. To achieve scaling factor between two invocations of an eBPF program, users can use CPU id as the key (which is typical for perf array usage model) to remember the previous value and do the calculation inside the eBPF program.


### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
