---
title: "Libbpf userspace function 'btf__add_decl_tag'"
description: "This page documents the 'btf__add_decl_tag' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_decl_tag`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_DECL_TAG` type with a given value to BTF object.

## Definition

`#!c int btf__add_decl_tag(struct btf *btf, const char *value, int ref_type_id, int component_idx);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `value`: tag value, can't be `NULL` or empty;
- `ref_type_id`: referenced type ID, it might not exist yet;
- `component_idx`: -1 for tagging reference type, otherwise struct/union member or function argument index;

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
