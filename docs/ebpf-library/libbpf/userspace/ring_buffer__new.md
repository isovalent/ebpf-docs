---
title: "Libbpf userspace function 'ring_buffer__new'"
description: "This page documents the 'ring_buffer__new' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring_buffer__new`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Create a new ring buffer manager.

## Definition

```c
typedef int (*ring_buffer_sample_fn)(void *ctx, void *data, size_t size); 

struct ring_buffer * ring_buffer__new(int map_fd, ring_buffer_sample_fn sample_cb, void *ctx, const struct ring_buffer_opts *opts);
```

**Parameters**

- `map_fd`: file descriptor of the [`BPF_MAP_TYPE_RINGBUF`](../../../linux/map-type/BPF_MAP_TYPE_RINGBUF.md) map of the first ring buffer that will be part of the ring buffer manager
- `sample_cb`: callback function that will be called when a sample is ready for the initial ring buffer
- `ctx`: context that will be passed to the callback function
- `opts`: options for the ring buffer

**Returns**

A pointer to the newly created ring buffer manager, or `NULL` on error. [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set to the error code on error.

### `struct ring_buffer_opts`

```c
struct ring_buffer_opts {
	size_t sz; /* size of this struct, for forward/backward compatibility */
};
```

## Usage

A ring buffer in this context is a circular buffer where eBPF programs are producers and userspace is the consumer. Even tough the returned value is `struct ring_buffer *`, it actually is a "manager" for multiple ring buffers (`struct ring`).

All ring buffers that are part of this ring buffer manager can be [`epoll`](https://man7.org/linux/man-pages/man7/epoll.7.html)-ed together, to get notified of pending data on any of the rings.

A ring buffer manager always contains at least one ring, the `map_fd`, `sample_cb`, and `ctx` parameters are used to create the first ring buffer. It internally calls [`ring_buffer__add`](ring_buffer__add.md) to add the first ring buffer.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
