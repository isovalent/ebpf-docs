---
title: "Libbpf userspace function 'ring__consume_n'"
description: "This page documents the 'ring__consume_n' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring__consume_n`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Consumes up to a requested amount of items from a ring buffer without event polling.

## Definition

`#!c int ring__consume_n(struct ring *r, size_t n);`

**Parameters**

- `r`: A ring buffer object.
- `n`: Maximum amount of items to consume.

**Return**

The number of items consumed, or a negative number if any of the callbacks return an error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
