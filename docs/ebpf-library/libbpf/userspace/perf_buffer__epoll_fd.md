---
title: "Libbpf userspace function 'perf_buffer__epoll_fd'"
description: "This page documents the 'perf_buffer__epoll_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `perf_buffer__epoll_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Get an file descriptor that can be used to sleep until data is available in the perf buffer.

## Definition

`#!c int perf_buffer__epoll_fd(const struct perf_buffer *pb);`

**Parameters**

- `pb`: perf buffer to get the [`epoll`](https://man7.org/linux/man-pages/man7/epoll.7.html) set file descriptor for.

**Returns**

File descriptor of an [`epoll`](https://man7.org/linux/man-pages/man7/epoll.7.html) set that can be used to sleep until data is available in the perf buffer.

## Usage

The returned file descriptor is created with [`epoll_create`](https://man7.org/linux/man-pages/man2/epoll_create.2.html). The file descriptor can be used with [`epoll_wait`](https://man7.org/linux/man-pages/man2/epoll_wait.2.html) to sleep until data is available in any of the perf buffer.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
