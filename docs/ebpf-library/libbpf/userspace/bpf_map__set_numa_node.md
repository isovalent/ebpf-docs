---
title: "Libbpf userspace function 'bpf_map__set_numa_node'"
description: "This page documents the 'bpf_map__set_numa_node' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_numa_node`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Set the NUMA node the map is or will be created on.

## Definition

`#!c int bpf_map__set_numa_node(struct bpf_map *map, __u32 numa_node);`

**Parameters**

- `map`: Pointer to the BPF map.
- `numa_node`: NUMA node the map is or will be created on.

**Return**

`0` on success, a negative error in case of failure.

## Usage

Setting a NUMA node might improve performance if the map is expected to be accessed a lot from that NUMA domain. This is an advanced option and should only be used if you know what you are doing.

This value can only be modified before the map is loaded.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
