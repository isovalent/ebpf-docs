---
title: "Libbpf userspace function 'libbpf_bpf_map_type_str'"
description: "This page documents the 'libbpf_bpf_map_type_str' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_bpf_map_type_str`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)
<!-- [/LIBBPF_TAG] -->

Converts the provided map type value into a textual representation.

## Definition

`#!c const char *libbpf_bpf_map_type_str(enum bpf_map_type t);`

**Parameters**

- `t`: The map type.

**Return**

Pointer to a static string identifying the map type. `NULL` is returned for unknown `bpf_map_type` values.

## Usage

This function allows you to take a ELF section name and query to which program type and attach type it corresponds to, if libbpf knows about it. Conversion happens following [this table](../../../linux/program-type/index.md#index-of-section-names)

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
