---
title: "Libbpf userspace function 'bpf_map__unpin'"
description: "This page documents the 'bpf_map__unpin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__unpin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Removes the file that serves as a 'pin' for the BPF map.

## Definition

`#!c int bpf_map__unpin(struct bpf_map *map, const char *path);`

**Parameters**

- `map`: The bpf_map to unpin
- `path`: A file path for the 'pin'

**Return**

`0`, on success; negative error, otherwise

## Usage

The `path` parameter can be `NULL`, in which case the `pin_path` map attribute is unpinned. If both the `path` parameter and `pin_path` map attribute are set, they must be equal.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
