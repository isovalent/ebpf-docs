---
title: "Libbpf userspace function 'bpf_map_delete_elem_flags'"
description: "This page documents the 'bpf_map_delete_elem_flags' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_delete_elem_flags`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_DELETE_ELEM`](../../../linux/syscall/BPF_MAP_DELETE_ELEM.md) syscall command.

## Definition

`#!c int bpf_map_delete_elem_flags(int fd, const void *key, __u64 flags);`

**Parameters**

- `fd`: file descriptor of the map to delete element from
- `key`: pointer to memory containing bytes of the key
- `flags`: flags passed to kernel for this operation

**Return**

`0`, on success; negative error, otherwise

## Usage

This function should only be used if you need precise control over the map element delete process. In most cases the [`bpf_map__delete_elem`](bpf_map__delete_elem.md) function should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
