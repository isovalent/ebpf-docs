---
title: "Libbpf userspace function 'bpf_map__map_extra'"
description: "This page documents the 'bpf_map__map_extra' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__map_extra`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Get the `map_extra` with which the map is or will be created.

## Definition

`#!c __u64 bpf_map__map_extra(const struct bpf_map *map);`

**Parameters**

- `map`: Pointer to the BPF map.

**Return**

[`map_extra`](../../../linux/syscall/BPF_MAP_CREATE.md#map_extra) info of the map.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
