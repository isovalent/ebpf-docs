---
title: "Libbpf userspace function 'ring__avail_data_size'"
description: "This page documents the 'ring__avail_data_size' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `ring__avail_data_size`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Returns the number of bytes in the ring buffer not yet consumed. This has no locking associated with it, so it can be inaccurate if operations are ongoing while this is called. However, it should still show the correct trend over the long-term.

## Definition

`#!c size_t ring__avail_data_size(const struct ring *r);`

**Parameters**

- `r`: A ring buffer object.

**Return**

The number of bytes not yet consumed.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
