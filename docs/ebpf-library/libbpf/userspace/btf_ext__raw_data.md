---
title: "Libbpf userspace function 'btf_ext__raw_data'"
description: "This page documents the 'btf_ext__raw_data' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf_ext__raw_data`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)
<!-- [/LIBBPF_TAG] -->

Get the raw data of the BTF extension object.

## Definition

`#!c const void *btf_ext__raw_data(const struct btf_ext *btf_ext, __u32 *size);`

**Parameters**

- `btf_ext`: BTF extension object
- `size`: Size of the raw data, written to by the function

**Return**

A `const` pointer to raw serialized BTF extension. Including a header, types and strings. The caller is responsible for freeing the memory.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
