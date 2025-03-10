---
title: "Libbpf userspace function 'bpf_map__set_pin_path'"
description: "This page documents the 'bpf_map__set_pin_path' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_pin_path`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)
<!-- [/LIBBPF_TAG] -->

Sets the path attribute that tells where the BPF map should be pinned. This does not actually create the 'pin'.

## Definition

`#!c int bpf_map__set_pin_path(struct bpf_map *map, const char *path);`

**Parameters**

- `map`: The bpf_map
- `path`: The path

**Return**

`0`, on success; negative error, otherwise

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
