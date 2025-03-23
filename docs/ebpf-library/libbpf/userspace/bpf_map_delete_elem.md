---
title: "Libbpf userspace function 'bpf_map_delete_elem'"
description: "This page documents the 'bpf_map_delete_elem' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_delete_elem`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_DELETE_ELEM`](../../../linux/syscall/BPF_MAP_DELETE_ELEM.md) syscall command.

## Definition

`#!c int bpf_map_delete_elem(int fd, const void *key);`

**Parameters**

- `fd`: file descriptor of the map to delete element from
- `key`: pointer to memory containing bytes of the key

**Return**

`0`, on success; negative error, otherwise

## Usage

This function should only be used if you need precise control over the map element delete process. In most cases the [`bpf_map__delete_elem`](bpf_map__delete_elem.md) function should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
