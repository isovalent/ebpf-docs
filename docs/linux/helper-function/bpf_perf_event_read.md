# Helper function `bpf_perf_event_read`

<!-- [FEATURE_TAG](bpf_perf_event_read) -->
[:octicons-tag-24: v4.3](https://github.com/torvalds/linux/commit/35578d7984003097af2b1e34502bc943d40c1804)
<!-- [/FEATURE_TAG] -->

This helper reads the value of a perf event counter.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
The value of the perf event counter read from the map, or a
negative error code in case of failure.

`#!c static __u64 (*bpf_perf_event_read)(void *map, __u64 flags) = (void *) 22;`

## Usage

This helper relies on a `map` of type [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md). The nature of the perf event counter is selected when `map` is updated with perf event file descriptors. The `map` is an array whose size is the number of available CPUs, and each cell contains a value relative to one CPU. The value to retrieve is indicated by `flags`, that contains the index of the CPU to look up, masked with `BPF_F_INDEX_MASK`. Alternatively, `flags` can be set to `BPF_F_CURRENT_CPU` to indicate that the value for the current CPU should be retrieved.

!!! note
    before Linux 4.13, only hardware perf event can be retrieved.

Also, be aware that the newer helper [`bpf_perf_event_read_value`](bpf_perf_event_read_value.md) is recommended over [`bpf_perf_event_read`](bpf_perf_event_read.md) in general. The latter has some ABI quirks where error and counter value are used as a return code (which is wrong to do since ranges may overlap). This issue is fixed with [`bpf_perf_event_read_value`](bpf_perf_event_read_value.md), which at the same time provides more features over the [`bpf_perf_event_read`](bpf_perf_event_read.md) interface. Please refer to the page of [`bpf_perf_event_read_value`](bpf_perf_event_read_value.md) for details.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->

<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
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
