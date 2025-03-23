---
title: "Libbpf userspace function 'btf__parse_elf'"
description: "This page documents the 'btf__parse_elf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__parse_elf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.4](https://github.com/libbpf/libbpf/releases/tag/v0.0.4)
<!-- [/LIBBPF_TAG] -->

Parse BTF from an ELF file.

## Definition

`#!c struct btf *btf__parse_elf(const char *path, struct btf_ext **btf_ext);`

**Parameters**

- `path`: path to the ELF file containing BTF data
- `btf_ext`: double pointer, will be populated with BTF extension data if present

**Return**

Return a pointer to a `struct btf` object on success, or `NULL` on failure. The caller is responsible for freeing the returned object with [`btf__free`](btf__free.md).

If `btf_ext` is not `NULL`, it will be populated with a pointer to a `struct btf_ext` object on success, or `NULL` on failure. The caller is responsible for freeing the returned object with [`btf_ext__free`](btf_ext__free.md).

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
