---
title: "Libbpf userspace function 'ring_buffer__consume_n'"
description: "This page documents the 'ring_buffer__consume_n' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring_buffer__consume_n`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Consume available ring buffer(s) data without event polling, up to `n` records.

## Definition

`#!c int ring_buffer__consume_n(struct ring_buffer *rb, size_t n);`

**Parameters**

- `rb`: ring buffer manager to consume data from
- `n`: number of records to consume

**Returns**

Number of records consumed across all registered ring buffers (or `n`, whichever is less), or negative number if any of the callbacks return error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
