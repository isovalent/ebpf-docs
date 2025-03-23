---
title: "Libbpf userspace function 'btf__add_str'"
description: "This page documents the 'btf__add_str' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_str`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Add a string s to the BTF string section.

## Definition

`#!c int btf__add_str(struct btf *btf, const char *s);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `s`: string to add

**Return**

`> 0` offset into string section, on success; `< 0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
