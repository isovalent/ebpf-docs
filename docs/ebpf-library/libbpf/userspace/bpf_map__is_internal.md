---
title: "Libbpf userspace function 'bpf_map__is_internal'"
description: "This page documents the 'bpf_map__is_internal' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__is_internal`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.3](https://github.com/libbpf/libbpf/releases/tag/v0.0.3)
<!-- [/LIBBPF_TAG] -->

Tells the caller whether or not the passed map is a special map created by libbpf automatically for things like global variables, `__ksym` externs, Kconfig values, etc

## Definition

`#!c bool bpf_map__is_internal(const struct bpf_map *map);`

**Parameters**

- `map`: the bpf_map

**Return**

`true`, if the map is an internal map; `false`, otherwise

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
