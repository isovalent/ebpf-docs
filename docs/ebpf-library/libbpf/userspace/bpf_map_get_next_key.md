---
title: "Libbpf userspace function 'bpf_map_get_next_key'"
description: "This page documents the 'bpf_map_get_next_key' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_get_next_key`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_GET_NEXT_KEY`](../../../linux/syscall/BPF_MAP_GET_NEXT_KEY.md) syscall command.

## Definition

`#!c int bpf_map_get_next_key(int fd, const void *key, void *next_key);`

**Parameters**

- `fd`: file descriptor of the map to get the next key from
- `key`: pointer to memory containing bytes of the key used for lookup
- `next_key`: pointer to memory in which the next key will be stored

**Return**

`0`, on success; negative error, otherwise

## Usage

This function should only be used if you need precise control over the map element lookup process. In most cases the [`bpf_map__get_next_key`](bpf_map__get_next_key.md) function should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
