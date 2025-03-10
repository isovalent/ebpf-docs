---
title: "Libbpf userspace function 'bpf_map__btf_key_type_id'"
description: "This page documents the 'bpf_map__btf_key_type_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__btf_key_type_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Get the BTF type ID of the key of a BPF map.

## Definition

`#!c __u32 bpf_map__btf_key_type_id(const struct bpf_map *map);`

**Parameters**

- `map`: Pointer to the BPF map.

**Return**

BTF type ID of the key of the map. Or `0` if no BTF type ID is available.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
