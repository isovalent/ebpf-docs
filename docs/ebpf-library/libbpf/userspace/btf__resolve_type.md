---
title: "Libbpf userspace function 'btf__resolve_type'"
description: "This page documents the 'btf__resolve_type' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__resolve_type`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Get the type ID, the given type points to.

## Definition

`#!c int btf__resolve_type(const struct btf *btf, __u32 type_id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `type_id`: ID of the BTF type

**Return**

Return the type ID, the given type points to, or a negative error code on failure.

## Usage

Certain BTF kinds such as `BTF_KIND_PTR` point to another type. This function resolves the type ID, the given type points to.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
