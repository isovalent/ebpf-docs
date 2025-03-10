---
title: "Libbpf userspace function 'bpf_map__set_key_size'"
description: "This page documents the 'bpf_map__set_key_size' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_key_size`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Set the size of the key of a BPF map.

## Definition

`#!c int bpf_map__set_key_size(struct bpf_map *map, __u32 size);`

**Parameters**

- `map`: Pointer to the BPF map.
- `size`: Size of the key of the BPF map in bytes.

**Return**

`0` on success, a negative error in case of failure.

## Usage

Changing the size of the key of a BPF map can only be done before it is loaded.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
