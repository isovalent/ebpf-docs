---
title: "Libbpf userspace function 'btf__add_func_param'"
description: "This page documents the 'btf__add_func_param' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_func_param`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new function parameter for current `FUNC_PROTO` type.

## Definition

`#!c int btf__add_func_param(struct btf *btf, const char *name, int type_id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the parameter, can be `NULL` or empty;
- `type_id`: type ID describing the type of the parameter.

**Return**

`0`, on success; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
