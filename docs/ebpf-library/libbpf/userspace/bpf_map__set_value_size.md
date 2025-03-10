---
title: "Libbpf userspace function 'bpf_map__set_value_size'"
description: "This page documents the 'bpf_map__set_value_size' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_value_size`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Sets map value size.

## Definition

`#!c int bpf_map__set_value_size(struct bpf_map *map, __u32 size);`

**Parameters**

- `map`: the BPF map instance

**Return**

`0`, on success; negative error, otherwise

## Usage

There is a special case for maps with associated memory-mapped regions, like the global data section maps (`bss`, `data`, `rodata`). When this function is used on such a map, the mapped region is resized. Afterward, an attempt is made to adjust the corresponding BTF info. This attempt is best-effort and can only succeed if the last variable of the data section map is an array. The array BTF type is replaced by a new BTF array type with a different length. Any previously existing pointers returned from [`bpf_map__initial_value`](bpf_map__initial_value.md) or corresponding data section skeleton pointer must be reinitialized.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
