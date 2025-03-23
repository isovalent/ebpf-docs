---
title: "Libbpf userspace function 'btf_ext__new'"
description: "This page documents the 'btf_ext__new' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf_ext__new`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.2](https://github.com/libbpf/libbpf/releases/tag/v0.0.2)
<!-- [/LIBBPF_TAG] -->

Create a new BTF extension object from raw data.

## Definition

`#!c struct btf_ext *btf_ext__new(const __u8 *data, __u32 size);`

**Parameters**

- `data`: Raw data of the BTF extension
- `size`: Size of the raw data

**Return**

A pointer to the BTF extension object. Or `NULL` is returned and [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
