---
title: "Libbpf userspace function 'btf__add_field'"
description: "This page documents the 'btf__add_field' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_field`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new field for the current `STRUCT`/`UNION` type in BTF object.

## Definition

`#!c int btf__add_field(struct btf *btf, const char *name, int field_type_id, __u32 bit_offset, __u32 bit_size);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the field, can be `NULL` or empty for anonymous field;
- `field_type_id`: type ID for the type describing field type;
- `bit_offset`: bit offset of the start of the field within struct/union;
- `bit_size`: bit size of a bitfield, `0` for non-bitfield fields;

**Return**

`0`, on success; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
