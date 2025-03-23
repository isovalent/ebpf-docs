---
title: "Libbpf userspace function 'btf__name_by_offset'"
description: "This page documents the 'btf__name_by_offset' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__name_by_offset`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Return the name of a BTF type by its offset. Alias of [`btf__str_by_offset`](btf__str_by_offset.md).

## Definition

`#!c const char *btf__name_by_offset(const struct btf *btf, __u32 offset);`

**Parameters**

- `btf`: BTF object
- `offset`: Offset of the type

**Return**

A `const` pointer to the name of the type.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
