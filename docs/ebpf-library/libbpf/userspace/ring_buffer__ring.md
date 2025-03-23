---
title: "Libbpf userspace function 'ring_buffer__ring'"
description: "This page documents the 'ring_buffer__ring' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring_buffer__ring`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Returns the ring buffer object inside a given ring buffer manager representing a single [`BPF_MAP_TYPE_RINGBUF`](../../../linux/map-type/BPF_MAP_TYPE_RINGBUF.md) map instance.

## Definition

`#!c struct ring *ring_buffer__ring(struct ring_buffer *rb, unsigned int idx);`

**Parameters**

- `rb`: A ringbuffer manager object.
- `idx`: An index into the ring buffers contained within the ring buffer manager object. The index is 0-based and corresponds to the order in which [`ring_buffer__add`](ring_buffer__add.md) was called.

**Return**

A ring buffer object on success; `NULL` and [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) set if the index is invalid.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
