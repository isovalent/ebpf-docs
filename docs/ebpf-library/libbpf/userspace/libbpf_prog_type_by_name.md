---
title: "Libbpf userspace function 'libbpf_prog_type_by_name'"
description: "This page documents the 'libbpf_prog_type_by_name' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_prog_type_by_name`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Resolve the program and attach type from the provided ELF section name.

## Definition

`#!c int libbpf_prog_type_by_name(const char *name, enum bpf_prog_type *prog_type, enum bpf_attach_type *expected_attach_type);`

**Parameters**

- `name`: The ELF section name.
- `prog_type`: Pointer to program type, will be set to the resolved program type.
- `expected_attach_type`: Pointer to expected attach type, will be set to the resolved expected attach type.

**Return**

`0` on success, negative error code on failure. 

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
