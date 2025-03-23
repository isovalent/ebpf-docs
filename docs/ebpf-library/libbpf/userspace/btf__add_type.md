---
title: "Libbpf userspace function 'btf__add_type'"
description: "This page documents the 'btf__add_type' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_type`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)
<!-- [/LIBBPF_TAG] -->

Add a BTF type `src_type` found in the BTF object `src_btf` to the BTF object `btf`.

## Definition

`#!c int btf__add_type(struct btf *btf, const struct btf *src_btf, const struct btf_type *src_type);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `src_btf`: pointer to a `struct btf` object containing the type to add
- `src_type`: pointer to a `struct btf_type` object to add

**Return**

`0`, on success; `< 0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
