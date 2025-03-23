---
title: "Libbpf userspace function 'ring__size'"
description: "This page documents the 'ring__size' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring__size`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Returns the total size of the ring buffer's map data area (excluding special producer/consumer pages). Effectively this gives the amount of usable bytes of data inside the ring buffer.

## Definition

`#!c size_t ring__size(const struct ring *r);`

**Parameters**

- `r`: A ring buffer object.

**Return**

The total size of the ring buffer map data area.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
