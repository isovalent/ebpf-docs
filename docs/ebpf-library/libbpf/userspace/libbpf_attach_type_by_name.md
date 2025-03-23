---
title: "Libbpf userspace function 'libbpf_attach_type_by_name'"
description: "This page documents the 'libbpf_attach_type_by_name' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_attach_type_by_name`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Resolve the attach type from the provided ELF section name.

## Definition

`#!c int libbpf_attach_type_by_name(const char *name, enum bpf_attach_type *attach_type);`

**Parameters**

- `name`: The ELF section name.
- `attach_type`: Pointer to attach type, will be set to the resolved attach type.

**Return**

`0` on success, negative error code on failure.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
