---
title: "Libbpf userspace function 'btf__raw_data'"
description: "This page documents the 'btf__raw_data' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__raw_data`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Get the raw data of the BTF object.

## Definition

`#!c const void *btf__raw_data(const struct btf *btf, __u32 *size);`

**Parameters**

- `btf`: BTF object
- `size`: Size of the raw data, written to by the function

**Return**

A `const` pointer to raw serialized BTF. Including a header, types and strings. The caller is responsible for freeing the memory.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
