---
title: "Libbpf userspace function 'libbpf_bpf_link_type_str'"
description: "This page documents the 'libbpf_bpf_link_type_str' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_bpf_link_type_str`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)
<!-- [/LIBBPF_TAG] -->

Converts the provided link type value into a textual representation.

## Definition

`#!c const char *libbpf_bpf_link_type_str(enum bpf_link_type t);`

**Parameters**

- `t`: The link type.

**Return**

Pointer to a static string identifying the link type. `NULL` is returned for unknown `bpf_link_type` values.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
