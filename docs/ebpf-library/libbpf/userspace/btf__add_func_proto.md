---
title: "Libbpf userspace function 'btf__add_func_proto'"
description: "This page documents the 'btf__add_func_proto' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_func_proto`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_FUNC_PROTO` type to BTF object.

## Definition

`#!c int btf__add_func_proto(struct btf *btf, int ret_type_id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `ret_type_id`: type ID for return result of a function

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

Function prototype initially has no arguments, but they can be added by [`btf__add_func_param`](btf__add_func_param.md) one by one, immediately after `btf__add_func_proto` succeeded.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
