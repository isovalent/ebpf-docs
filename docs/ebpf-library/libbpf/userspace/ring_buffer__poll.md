---
title: "Libbpf userspace function 'ring_buffer__poll'"
description: "This page documents the 'ring_buffer__poll' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring_buffer__poll`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Poll for available data and consume records, if any are available.

## Definition

`#!c int ring_buffer__poll(struct ring_buffer *rb, int timeout_ms);`

**Parameters**

- `rb`: ring buffer manager to poll
- `timeout_ms`: timeout in milliseconds, wait for this long if no data is available

**Returns**

Number of records consumed (or `INT_MAX`, whichever is less), or negative number, if any of the registered callbacks returned error.

## Usage

Call this function to poll for available data on any of the ring buffers that are part of the ring buffer manager. If any data is available, the registered callback functions will be called. If no data is available, the function will wait for `timeout_ms` milliseconds for data to arrive and will block until then.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
