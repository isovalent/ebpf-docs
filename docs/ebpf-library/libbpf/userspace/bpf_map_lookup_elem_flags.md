---
title: "Libbpf userspace function 'bpf_map_lookup_elem_flags'"
description: "This page documents the 'bpf_map_lookup_elem_flags' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_lookup_elem_flags`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.2](https://github.com/libbpf/libbpf/releases/tag/v0.0.2)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_LOOKUP_ELEM`](../../../linux/syscall/BPF_MAP_LOOKUP_ELEM.md) syscall command.

## Definition

`#!c int bpf_map_lookup_elem_flags(int fd, const void *key, void *value, __u64 flags);`

**Parameters**

- `fd`: file descriptor of the map to lookup element in
- `key`: pointer to memory containing bytes of the key used for lookup
- `value`: pointer to memory in which looked up value will be stored
- `flags`: flags passed to kernel for this operation

**Return**

`0`, on success; negative error, otherwise

## Usage

This function should only be used if you need precise control over the map element lookup process. In most cases the [`bpf_map__lookup_elem`](bpf_map__lookup_elem.md) function should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
