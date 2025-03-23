---
title: "Libbpf userspace function 'btf__add_array'"
description: "This page documents the 'btf__add_array' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_array`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_ARRAY` type to BTF object.

## Definition

`#!c int btf__add_array(struct btf *btf, int index_type_id, int elem_type_id, __u32 nr_elems);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `index_type_id`: type ID of the type describing array index;
- `elem_type_id`: type ID of the type describing array element;
- `nr_elems`: the size of the array;

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
