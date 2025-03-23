---
title: "Libbpf userspace function 'user_ring_buffer__new'"
description: "This page documents the 'user_ring_buffer__new' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `user_ring_buffer__new`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.1.0](https://github.com/libbpf/libbpf/releases/tag/v1.1.0)
<!-- [/LIBBPF_TAG] -->

Creates a new instance of a user ring buffer.

## Definition

`#!c struct user_ring_buffer * user_ring_buffer__new(int map_fd, const struct user_ring_buffer_opts *opts);`

**Parameters**

- `map_fd`: A file descriptor to a [`BPF_MAP_TYPE_USER_RINGBUF`](../../../linux/map-type/BPF_MAP_TYPE_USER_RINGBUF.md) map.
- `opts`: Options for how the ring buffer should be created.

**Return**

A user ring buffer on success; `NULL` and [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) being set on a
failure.

### `struct user_ring_buffer_opts`

```c
struct user_ring_buffer_opts {
	size_t sz; /* size of this struct, for forward/backward compatibility */
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
