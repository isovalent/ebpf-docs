---
title: "Libbpf userspace function 'btf__find_str'"
description: "This page documents the 'btf__find_str' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__find_str`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Find an offset in BTF string section that corresponds to a given string `s`.

## Definition

`#!c int btf__find_str(struct btf *btf, const char *s);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `s`: string to find

**Return**

`>0` offset into string section, if string is found; `-ENOENT`, if string is not in the string section; `<0`, on any other error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
