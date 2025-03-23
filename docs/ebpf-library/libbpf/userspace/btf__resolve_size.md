---
title: "Libbpf userspace function 'btf__resolve_size'"
description: "This page documents the 'btf__resolve_size' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__resolve_size`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Resolve the size of a BTF type.

## Definition

`#!c __s64 btf__resolve_size(const struct btf *btf, __u32 type_id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `type_id`: ID of the BTF type

**Return**

Return the size of the BTF type in bytes on success, or a negative error code on failure.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
