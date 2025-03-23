---
title: "Libbpf userspace function 'btf__pointer_size'"
description: "This page documents the 'btf__pointer_size' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__pointer_size`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Return pointer size this BTF instance assumes.

## Definition

`#!c size_t btf__pointer_size(const struct btf *btf);`

**Parameters**

- `btf`: pointer to a `struct btf` object

**Return**

The size of the pointer in bytes.

## Usage

The size is heuristically determined by looking for 'long' or 'unsigned long' integer type and recording its size in bytes. If BTF type information doesn't have any such type, this function returns `0`. In the latter case, native architecture's pointer size is assumed, so will be either `4` or `8`, depending on architecture that libbpf was compiled for. It's possible to override guessed value by using [`btf__set_pointer_size`](btf__set_pointer_size.md) API.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
