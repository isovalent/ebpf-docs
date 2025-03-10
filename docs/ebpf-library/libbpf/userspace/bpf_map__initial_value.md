---
title: "Libbpf userspace function 'bpf_map__initial_value'"
description: "This page documents the 'bpf_map__initial_value' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__initial_value`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Get the initial value of a BPF map.

## Definition

`#!c void *bpf_map__initial_value(const struct bpf_map *map, size_t *psize);`

**Parameters**

- `map`: Pointer to the BPF map.
- `psize`: Pointer to the size of the initial value.

**Return**

Pointer to the initial value of the BPF map. `NULL` on error.

## Usage

Libbpf automatically creates array maps for certain ELF sections in which global variables are stored. This function allows you to get the initial value of these maps and [arena maps](../../../linux/map-type/BPF_MAP_TYPE_ARENA.md).

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
