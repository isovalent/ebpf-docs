---
title: "Libbpf userspace function 'btf__add_enum64_value'"
description: "This page documents the 'btf__add_enum64_value' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_enum64_value`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)
<!-- [/LIBBPF_TAG] -->

Append new enum value for the current `ENUM64` type.

## Definition

`#!c int btf__add_enum64_value(struct btf *btf, const char *name, __u64 value);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the enumerator value, can't be `NULL` or empty;
- `value`: integer value corresponding to enum value `name`.

**Return**

`0`, on success; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
