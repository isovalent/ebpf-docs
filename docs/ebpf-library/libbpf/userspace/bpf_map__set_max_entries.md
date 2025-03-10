---
title: "Libbpf userspace function 'bpf_map__set_max_entries'"
description: "This page documents the 'bpf_map__set_max_entries' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_max_entries`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Set the maximum number of entries in a BPF map.

## Definition

`#!c int bpf_map__set_max_entries(struct bpf_map *map, __u32 max_entries);`

**Parameters**

- `map`: Pointer to the BPF map.
- `max_entries`: Maximum number of entries in the BPF map.

**Return**

`0` on success, a negative error in case of failure.

## Usage

The maximum number of entries can only be modified before the map is loaded.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
