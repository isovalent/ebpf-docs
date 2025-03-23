---
title: "Libbpf userspace function 'perf_buffer__buffer_fd'"
description: "This page documents the 'perf_buffer__buffer_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `perf_buffer__buffer_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Return perf_event FD of a ring buffer in `buf_idx` slot of [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../../../linux/map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) BPF map. This file descriptor can be polled for new data using [`select`](https://man7.org/linux/man-pages/man2/select.2.html)/[`poll`](https://man7.org/linux/man-pages/man2/poll.2.html)/[`epoll`](https://man7.org/linux/man-pages/man7/epoll.7.html) Linux syscalls.

## Definition

`#!c int perf_buffer__buffer_fd(const struct perf_buffer *pb, size_t buf_idx);`

**Parameters**

- `pb`: perf buffer manager to get file descriptor from
- `buf_idx`: index of the perf ring buffer to get file descriptor from

**Returns**

File descriptor of the perf ring buffer in the `buf_idx` slot of the perf buffer, or negative number on error. [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) will be set to the error code.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
