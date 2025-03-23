---
title: "Libbpf userspace function 'ring__consume'"
description: "This page documents the 'ring__consume' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring__consume`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Consumes available ring buffer data without event polling.

## Definition

`#!c int ring__consume(struct ring *r);`

**Parameters**

- `r`: A ring buffer object.

**Return**

The number of records consumed (or `INT_MAX`, whichever is less), or a negative number if any of the callbacks return an error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
