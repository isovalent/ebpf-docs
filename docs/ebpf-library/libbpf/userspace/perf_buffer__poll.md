---
title: "Libbpf userspace function 'perf_buffer__poll'"
description: "This page documents the 'perf_buffer__poll' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `perf_buffer__poll`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.4](https://github.com/libbpf/libbpf/releases/tag/v0.0.4)
<!-- [/LIBBPF_TAG] -->

Poll for available data and consume records, if any are available.

## Definition

`#!c int perf_buffer__poll(struct perf_buffer *pb, int timeout_ms);`

**Parameters**

- `pb`: perf buffer to poll
- `timeout_ms`: timeout in milliseconds, wait for this long if no data is available

**Returns**

Number of records consumed (or `INT_MAX`, whichever is less), or negative number, if any of the registered callbacks returned error.

## Usage

Call this function to poll for available data on the perf buffer. If any data is available, the registered callback functions will be called. If no data is available, the function will wait for `timeout_ms` milliseconds for data to arrive and will block until then.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
