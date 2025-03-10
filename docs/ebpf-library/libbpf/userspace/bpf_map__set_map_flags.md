---
title: "Libbpf userspace function 'bpf_map__set_map_flags'"
description: "This page documents the 'bpf_map__set_map_flags' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_map_flags`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Set the flags the map will be loaded with.

## Definition

`#!c int bpf_map__set_map_flags(struct bpf_map *map, __u32 flags);`

**Parameters**

- `map`: Pointer to the BPF map.
- `flags`: [Flags](../../../linux/syscall/BPF_MAP_CREATE.md#flags) the map will be loaded with.

**Return**

`0` on success, a negative error in case of failure.

## Usage

Flags can only be modified before the map is loaded.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
