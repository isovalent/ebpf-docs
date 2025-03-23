---
title: "Libbpf userspace function 'btf__add_type_tag'"
description: "This page documents the 'btf__add_type_tag' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_type_tag`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_TYPE_TAG` type to BTF object.

## Definition

`#!c int btf__add_type_tag(struct btf *btf, const char *value, int ref_type_id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `value`: tag value, can't be `NULL` or empty;
- `ref_type_id`: referenced type ID, it might not exist yet;

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

Set `info->kflag` to `1`, indicating this tag is an `__attribute__`.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
