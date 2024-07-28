---
title: "Helper Function 'bpf_xdp_output'"
description: "This page documents the 'bpf_xdp_output' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_output`

<!-- [FEATURE_TAG](bpf_xdp_output) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/d831ee84bfc9173eecf30dbbc2553ae81b996c60)
<!-- [/FEATURE_TAG] -->

This helper writes a raw `data` blob into a special BPF perf event held by `map` of type [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_xdp_output)(void *ctx, void *map, __u64 flags, void *data, __u64 size) = (void *) 121;`

## Usage

The perf event must have the following attributes: `PERF_SAMPLE_RAW` as `sample_type`, `PERF_TYPE_SOFTWARE` as `type`, and `PERF_COUNT_SW_BPF_OUTPUT` as `config`.

The `flags` are used to indicate the index in `map` for which the value must be put, masked with `BPF_F_INDEX_MASK`. Alternatively, `flags` can be set to `BPF_F_CURRENT_CPU` to indicate that the index of the current CPU core should be used.

The value to write, of `size`, is passed through eBPF stack and pointed by `data`.

`ctx` is a pointer to in-kernel struct xdp_buff.

This helper is similar to [`bpf_perf_event_output`](bpf_perf_event_output.md) but restricted to raw_tracepoint bpf programs.

### Program types

This helper call can be used in the following program types:

 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
