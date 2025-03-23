---
title: "Libbpf userspace function 'btf__parse_raw'"
description: "This page documents the 'btf__parse_raw' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__parse_raw`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Parse BTF from a file containing raw BTF data.

## Definition

`#!c struct btf *btf__parse_raw(const char *path);`

**Parameters**

- `path`: path to the file

**Return**

Return a pointer to a `struct btf` object on success, or `NULL` on failure. The caller is responsible for freeing the returned object with [`btf__free`](btf__free.md).

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
