---
title: "Libbpf userspace function 'perf_buffer__buffer'"
description: "This page documents the 'perf_buffer__buffer' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `perf_buffer__buffer`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)
<!-- [/LIBBPF_TAG] -->

Returns the per-CPU raw [`mmap`](https://man7.org/linux/man-pages/man2/mmap.2.html)'ed underlying memory region of the ring buffer.

## Definition

`#!c int perf_buffer__buffer(struct perf_buffer *pb, int buf_idx, void **buf, size_t *buf_size);`

**Parameters**

- `pb`: the perf buffer structure
- `buf_idx`: the buffer index to retrieve
- `buf`: (out) gets the base pointer of the [`mmap`](https://man7.org/linux/man-pages/man2/mmap.2.html)'ed memory
- `buf_size`: (out) gets the size of the [`mmap`](https://man7.org/linux/man-pages/man2/mmap.2.html)'ed region

**Return**

`0` on success, negative error code for failure

## Usage

This ring buffer can be used to implement a custom events consumer. The ring buffer starts with the [`struct perf_event_mmap_page`](https://elixir.bootlin.com/linux/v6.13.7/source/include/uapi/linux/perf_event.h#L580), which holds the ring buffer management fields, when accessing the header structure it's important to be <nospell>SMP</nospell> aware. You can refer to [`perf_event_read_simple`](https://github.com/libbpf/libbpf/blob/374036c9f1cdfe2a8df98d9d6a53c34fd02de14b/src/libbpf.c#L13172) for a simple example.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
