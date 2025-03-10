---
title: "Libbpf userspace function 'bpf_object__next_map'"
description: "This page documents the 'bpf_object__next_map' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__next_map`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Iterate over the maps in a BPF object.

## Definition

`#!c struct bpf_map * bpf_object__next_map(const struct bpf_object *obj, const struct bpf_map *map);`

**Parameters**

- `obj`: The BPF object to iterate over.
- `map`: The current map, or `NULL` to start iteration.

**Returns**

The next map in the object, or `NULL` if there are no more maps.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
