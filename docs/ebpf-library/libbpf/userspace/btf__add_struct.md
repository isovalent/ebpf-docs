---
title: "Libbpf userspace function 'btf__add_struct'"
description: "This page documents the 'btf__add_struct' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_struct`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_STRUCT` type to BTF object.

## Definition

`#!c int btf__add_struct(struct btf *btf, const char *name, __u32 sz);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the struct, can be NULL or empty for anonymous structs;
- `sz`: size of the struct, in bytes;

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

Struct initially has no fields in it. Fields can be added by [`btf__add_field`](btf__add_field.md) right after `btf__add_struct` succeeds.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
