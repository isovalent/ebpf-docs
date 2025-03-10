---
title: "Libbpf userspace function 'bpf_map__numa_node'"
description: "This page documents the 'bpf_map__numa_node' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__numa_node`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Get the NUMA node the map is or will be created on.

## Definition

`#!c __u32 bpf_map__numa_node(const struct bpf_map *map);`

**Parameters**

- `map`: Pointer to the BPF map.

**Return**

NUMA node the map is or will be created on. Or `0` if no preference is given.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
