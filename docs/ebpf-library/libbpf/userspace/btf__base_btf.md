---
title: "Libbpf userspace function 'btf__base_btf'"
description: "This page documents the 'btf__base_btf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__base_btf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.3.0](https://github.com/libbpf/libbpf/releases/tag/v0.3.0)
<!-- [/LIBBPF_TAG] -->

Get the base BTF object of a BTF object.

## Definition

`#!c const struct btf *btf__base_btf(const struct btf *btf);`

**Parameters**

- `btf`: pointer to a `struct btf` object

**Return**

Return a pointer to a `struct btf` object on success, or `NULL` if the BTF object has no base.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
