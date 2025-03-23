---
title: "Libbpf userspace function 'user_ring_buffer__discard'"
description: "This page documents the 'user_ring_buffer__discard' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `user_ring_buffer__discard`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.1.0](https://github.com/libbpf/libbpf/releases/tag/v1.1.0)
<!-- [/LIBBPF_TAG] -->

Discards a previously reserved sample.

## Definition

`#!c void user_ring_buffer__discard(struct user_ring_buffer *rb, void *sample);`

**Parameters**

- `rb`: The user ring buffer.
- `sample`: A reserved sample.

## Usage

It is not necessary to synchronize amongst multiple producers when invoking this function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
