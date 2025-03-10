---
title: "Libbpf userspace function 'bpf_map__set_initial_value'"
description: "This page documents the 'bpf_map__set_initial_value' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_initial_value`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Set the initial value of a <nospell>mmap-ed</nospell> BPF map.

## Definition

`#!c int bpf_map__set_initial_value(struct bpf_map *map, const void *data, size_t size);`

**Parameters**

- `map`: Pointer to the BPF map.
- `data`: Pointer to the initial value.
- `size`: Size of the initial value.

**Return**

`0` on success, or a negative error in case of failure.

## Usage

Libbpf automatically creates array maps for certain ELF sections in which global variables are stored. This function allows you to set the initial value of these maps and [arena maps](../../../linux/map-type/BPF_MAP_TYPE_ARENA.md).

This function should be used after creating the map, but before loading the programs that use it.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
