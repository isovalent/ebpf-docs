---
title: "Libbpf userspace function 'btf__add_datasec_var_info'"
description: "This page documents the 'btf__add_datasec_var_info' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_datasec_var_info`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new data section variable information entry for current `DATASEC` type.

## Definition

`#!c int btf__add_datasec_var_info(struct btf *btf, int var_type_id, __u32 offset, __u32 byte_sz);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `var_type_id`: type ID describing the type of the variable;
- `offset`: variable offset within data section, in bytes;
- `byte_sz`: variable size, in bytes.

**Return**

`0`, on success; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
