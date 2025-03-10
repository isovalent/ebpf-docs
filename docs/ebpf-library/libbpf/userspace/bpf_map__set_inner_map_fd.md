---
title: "Libbpf userspace function 'bpf_map__set_inner_map_fd'"
description: "This page documents the 'bpf_map__set_inner_map_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_inner_map_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Set the file descriptor of an inner map for a map.

## Definition

`#!c int bpf_map__set_inner_map_fd(struct bpf_map *map, int fd);`

**Parameters**

- `map`: The bpf_map
- `fd`: The file descriptor of the inner map

**Return**

`0`, on success; negative error, otherwise

## Usage

When loading map-in-map maps, such as [`BPF_MAP_TYPE_ARRAY_OF_MAPS`](../../../linux/map-type/BPF_MAP_TYPE_ARRAY_OF_MAPS.md) or [`BPF_MAP_TYPE_HASH_OF_MAPS`](../../../linux/map-type/BPF_MAP_TYPE_HASH_OF_MAPS.md.md), the verifier needs what sort of maps you will be putting into it. To communicate that, a map with the same attributes as will be used as values must be loaded first, and then its file descriptor passed to the outer map before loading.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
