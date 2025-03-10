---
title: "Libbpf userspace function 'bpf_map__set_map_extra'"
description: "This page documents the 'bpf_map__set_map_extra' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_map_extra`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Set the `map_extra` with which the map will be created.

## Definition

`#!c int bpf_map__set_map_extra(struct bpf_map *map, __u64 map_extra);`

**Parameters**

- `map`: Pointer to the BPF map.
- `map_extra`: [`map_extra`](../../../linux/syscall/BPF_MAP_CREATE.md#map_extra) info to set.

**Return**

`0` on success, or a negative error in case of failure.

## Usage

The `map_extra` can only be set before the map is created.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
