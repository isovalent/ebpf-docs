---
title: "Libbpf userspace function 'btf__align_of'"
description: "This page documents the 'btf__align_of' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__align_of`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Get the alignment of a type in the BTF object.

## Definition

`#!c int btf__align_of(const struct btf *btf, __u32 id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `id`: ID of the BTF type

**Return**

Return the alignment of the type on success, or a negative error code on failure.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
