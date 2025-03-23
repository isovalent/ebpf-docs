---
title: "Libbpf userspace function 'btf__add_enum'"
description: "This page documents the 'btf__add_enum' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_enum`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new BTF_KIND_ENUM type to BTF object.

## Definition

`#!c int btf__add_enum(struct btf *btf, const char *name, __u32 bytes_sz);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the enum, can be `NULL` or empty for anonymous enums;
- `bytes_sz`: size of the enum, in bytes.

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

Enum initially has no enum values in it (and corresponds to enum forward declaration). Enumerator values can be added by [`btf__add_enum_value`](btf__add_enum_value.md) immediately after `btf__add_enum` succeeds.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
