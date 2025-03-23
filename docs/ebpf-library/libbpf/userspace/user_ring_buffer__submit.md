---
title: "Libbpf userspace function 'user_ring_buffer__submit'"
description: "This page documents the 'user_ring_buffer__submit' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `user_ring_buffer__submit`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.1.0](https://github.com/libbpf/libbpf/releases/tag/v1.1.0)
<!-- [/LIBBPF_TAG] -->

Submits a previously reserved sample into the ring buffer.

## Definition

`#!c void user_ring_buffer__submit(struct user_ring_buffer *rb, void *sample);`

**Parameters**

- `rb`: The user ring buffer.
- `sample`: A reserved sample.

## Usage

It is not necessary to synchronize amongst multiple producers when invoking this function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
