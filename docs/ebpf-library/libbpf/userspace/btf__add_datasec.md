---
title: "Libbpf userspace function 'btf__add_datasec'"
description: "This page documents the 'btf__add_datasec' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_datasec`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_DATASEC` type to BTF object.

## Definition

`#!c int btf__add_datasec(struct btf *btf, const char *name, __u32 byte_sz);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the data section, can't be `NULL` or empty;
- `byte_sz`: size of the data section, in bytes.

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

Data section is initially empty. Variables info can be added with [`btf__add_datasec_var_info`](btf__add_datasec_var_info.md) calls, after `btf__add_datasec` succeeds.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
