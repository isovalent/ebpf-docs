---
title: "Libbpf userspace function 'perf_buffer__consume_buffer'"
description: "This page documents the 'perf_buffer__consume_buffer' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `perf_buffer__consume_buffer`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Consume data from perf ring buffer corresponding to slot `buf_idx` in [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../../../linux/map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) BPF map without waiting/polling. If there is no data to consume, do nothing and return success.

## Definition

`#!c int perf_buffer__consume_buffer(struct perf_buffer *pb, size_t buf_idx);`

**Parameters**

- `pb`: perf buffer manager to consume data from
- `buf_idx`: index of the perf ring buffer to consume data from

**Returns**

`0` on success; `<0` on failure.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
