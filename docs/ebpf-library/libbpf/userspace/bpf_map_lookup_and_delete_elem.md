---
title: "Libbpf userspace function 'bpf_map_lookup_and_delete_elem'"
description: "This page documents the 'bpf_map_lookup_and_delete_elem' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_lookup_and_delete_elem`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_LOOKUP_AND_DELETE_ELEM`](../../../linux/syscall/BPF_MAP_LOOKUP_AND_DELETE_ELEM.md) syscall command.

## Definition

`#!c int bpf_map_lookup_and_delete_elem(int fd, const void *key, void *value);`

**Parameters**

- `fd`: file descriptor of the map to lookup element in
- `key`: pointer to memory containing bytes of the key used for lookup
- `value`: pointer to memory in which looked up value will be stored

**Return**

`0`, on success; negative error, otherwise

## Usage

This function should only be used if you need precise control over the map element lookup and delete process. In most cases the [`bpf_map__lookup_and_delete_elem`](bpf_map__lookup_and_delete_elem.md) function should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
