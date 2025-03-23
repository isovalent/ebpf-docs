---
title: "Libbpf userspace function 'btf__find_by_name'"
description: "This page documents the 'btf__find_by_name' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__find_by_name`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Find the type id of a type by its name, in the BTF object.

## Definition

`#!c __s32 btf__find_by_name(const struct btf *btf, const char *type_name);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `type_name`: name of the type

**Return**

Return the type id of the type on success, or a negative error code on failure.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
