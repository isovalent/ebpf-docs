---
title: "Libbpf userspace function 'bpf_map__reuse_fd'"
description: "This page documents the 'bpf_map__reuse_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__reuse_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Let a `bpf_map` reuse a file descriptor.

## Definition

`#!c int bpf_map__reuse_fd(struct bpf_map *map, int fd);`

**Parameters**

- `map`: Pointer to the `bpf_map` object.
- `fd`: File descriptor to reuse.

**Return**

`0` on success, a negative error code on failure.

## Usage

This function assigns an existing file descriptor to a `bpf_map` object. When a program references this map, it will use the file descriptor provided by the user instead of creating a new one. If the map is already used by other programs, it will cause the map to be shared.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
