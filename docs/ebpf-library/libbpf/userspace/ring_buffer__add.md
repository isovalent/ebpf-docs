---
title: "Libbpf userspace function 'ring_buffer__add'"
description: "This page documents the 'ring_buffer__add' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring_buffer__add`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Add a new ring buffer to a ring buffer manager.

## Definition

```c
typedef int (*ring_buffer_sample_fn)(void *ctx, void *data, size_t size);

int ring_buffer__add(struct ring_buffer *rb, int map_fd, ring_buffer_sample_fn sample_cb, void *ctx);
```

**Parameters**

- `rb`: ring buffer manager to add the new ring buffer to
- `map_fd`: file descriptor of the [`BPF_MAP_TYPE_RINGBUF`](../../../linux/map-type/BPF_MAP_TYPE_RINGBUF.md) map of the ring buffer to add
- `sample_cb`: callback function that will be called when a sample is ready for the ring buffer
- `ctx`: context that will be passed to the callback function

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
