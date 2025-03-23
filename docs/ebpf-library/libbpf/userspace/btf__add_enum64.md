---
title: "Libbpf userspace function 'btf__add_enum64'"
description: "This page documents the 'btf__add_enum64' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_enum64`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_ENUM64` type to BTF object.

## Definition

`#!c int btf__add_enum64(struct btf *btf, const char *name, __u32 bytes_sz, bool is_signed);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the enum, can be `NULL` or empty for anonymous enums;
- `bytes_sz`: size of the enum, in bytes.
- `is_signed`: whether the enum values are signed or not.

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

Enum initially has no enum values in it (and corresponds to enum forward declaration). Enumerator values can be added by [`btf__add_enum64_value`](btf__add_enum64_value.md) immediately after `btf__add_enum64` succeeds.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
