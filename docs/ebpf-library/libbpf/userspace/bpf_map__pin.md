---
title: "Libbpf userspace function 'bpf_map__pin'"
description: "This page documents the 'bpf_map__pin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__pin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Creates a file that serves as a 'pin' for the BPF map. This increments the reference count on the BPF map which will keep the BPF map loaded even after the userspace process which loaded it has exited.

## Definition

`#!c int bpf_map__pin(struct bpf_map *map, const char *path);`

**Parameters**

- `map`: The bpf_map to pin
- `path`: A file path for the 'pin'

**Return**

`0`, on success; negative error, otherwise

## Usage

If `path` is `NULL` the maps `pin_path` attribute will be used. If this is also `NULL`, an error will be returned and the map will not be pinned.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
