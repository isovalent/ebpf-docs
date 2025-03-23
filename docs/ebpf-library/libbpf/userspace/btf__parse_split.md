---
title: "Libbpf userspace function 'btf__parse_split'"
description: "This page documents the 'btf__parse_split' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__parse_split`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.3.0](https://github.com/libbpf/libbpf/releases/tag/v0.3.0)
<!-- [/LIBBPF_TAG] -->

Parse BTF from a file. The file can contain raw BTF data, or BTF data embedded in an ELF file.

## Definition

`#!c struct btf *btf__parse_split(const char *path, struct btf *base_btf);`

**Parameters**

- `path`: path to the file
- `base_btf`: pointer to a `struct btf` object that will be used as the base BTF object

**Return**

Return a pointer to a `struct btf` object on success, or `NULL` on failure. The caller is responsible for freeing the returned object with [`btf__free`](btf__free.md).

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
