---
title: "Libbpf userspace function 'perf_buffer__buffer_cnt'"
description: "This page documents the 'perf_buffer__buffer_cnt' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `perf_buffer__buffer_cnt`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Return number of [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../../../linux/map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) map slots set up by this perf buffer.

## Definition

`#!c size_t perf_buffer__buffer_cnt(const struct perf_buffer *pb);`

**Parameters**

- `pb`: perf buffer manager to get number of slots from

**Returns**

Number of slots in the perf buffer.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
