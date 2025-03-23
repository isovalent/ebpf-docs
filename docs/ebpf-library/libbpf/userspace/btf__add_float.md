---
title: "Libbpf userspace function 'btf__add_float'"
description: "This page documents the 'btf__add_float' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_float`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_FLOAT` type to BTF object.

## Definition

`#!c int btf__add_float(struct btf *btf, const char *name, size_t byte_sz);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the type; non-empty, non-NULL type name;
- `byte_sz`: size of the type, in bytes;

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
