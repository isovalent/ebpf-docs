---
title: "Libbpf userspace function 'ring_buffer__consume'"
description: "This page documents the 'ring_buffer__consume' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring_buffer__consume`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Consume available ring buffer(s) data without event polling.

## Definition

`#!c int ring_buffer__consume(struct ring_buffer *rb);`

**Parameters**

- `rb`: ring buffer manager to consume data from

**Returns**

Number of records consumed across all registered ring buffers (or `INT_MAX`, whichever is less), or negative number if any of the callbacks return error.

## Usage

Call this function to consume available data on any of the ring buffers that are part of the ring buffer manager. This function will not block, it will consume all available data and return immediately.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
