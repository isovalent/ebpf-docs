---
title: "Libbpf userspace function 'bpf_link__update_map'"
description: "This page documents the 'bpf_link__update_map' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link__update_map`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/releases/tag/v1.2.0)
<!-- [/LIBBPF_TAG] -->

Swap out the struct ops map of the link.

## Definition

`#!c int bpf_link__update_map(struct bpf_link *link, const struct bpf_map *map);`

**Parameters**

- `link`: Pointer to the BPF link.

## Usage

This function swaps out the [`BPF_MAP_TYPE_STRUCT_OPS`](../../../linux/map-type/BPF_MAP_TYPE_STRUCT_OPS.md) map that is linked to the kernel.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
