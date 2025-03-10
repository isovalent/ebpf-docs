---
title: "Libbpf userspace function 'bpf_map__get_next_key'"
description: "This page documents the 'bpf_map__get_next_key' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__get_next_key`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

allows to iterate BPF map keys by fetching next key that follows current key.

## Definition

`#!c int bpf_map__get_next_key(const struct bpf_map *map, const void *cur_key, void *next_key, size_t key_sz);`

**Parameters**

- `map`: BPF map to fetch next key from
- `cur_key`: pointer to memory containing bytes of current key or `NULL` to
fetch the first key
- `next_key`: pointer to memory to write next key into
- `key_sz`: size in bytes of key data, needs to match BPF map definition's `key_size`

**Return**

`0`, on success; `-ENOENT` if `cur_key` is the last key in BPF map; negative error, otherwise

## Usage

[`bpf_map__get_next_key()`](bpf_map__get_next_key.md) is high-level equivalent of [`bpf_map_get_next_key()`](bpf_map_get_next_key.md) API with added check for key size.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
