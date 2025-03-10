---
title: "Libbpf userspace function 'bpf_map__delete_elem'"
description: "This page documents the 'bpf_map__delete_elem' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__delete_elem`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Allows to delete element in BPF map that corresponds to provided key.

## Definition

`#!c int bpf_map__delete_elem(const struct bpf_map *map, const void *key, size_t key_sz, __u64 flags);`

**Parameters**

- `map`: BPF map to delete element from
- `key`: pointer to memory containing bytes of the key
- `key_sz`: size in bytes of key data, needs to match BPF map definition's `key_size`
- `flags`: flags passed to kernel for this operation

**Return**

`0`, on success; negative error, otherwise

## Usage

[`bpf_map__delete_elem()`](bpf_map__delete_elem.md) is high-level equivalent of [`bpf_map_delete_elem()`](bpf_map_delete_elem.md) API with added check for key size.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
