---
title: "Libbpf userspace function 'btf__add_fwd'"
description: "This page documents the 'btf__add_fwd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_fwd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_FWD` type to BTF object.

## Definition

`#!c int btf__add_fwd(struct btf *btf, const char *name, enum btf_fwd_kind fwd_kind);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the forward declaration, can't be `NULL` or empty;
- `fwd_kind`: kind of forward declaration, one of `BTF_FWD_STRUCT`, `BTF_FWD_UNION`, or `BTF_FWD_ENUM`.

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
