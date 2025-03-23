---
title: "Libbpf userspace function 'btf__add_restrict'"
description: "This page documents the 'btf__add_restrict' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_restrict`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_RESTRICT` type to BTF object.

## Definition

`#!c int btf__add_restrict(struct btf *btf, int ref_type_id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `ref_type_id`: referenced type ID, it might not exist yet;

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
