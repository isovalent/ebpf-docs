---
title: "Libbpf userspace function 'bpf_map__attach_struct_ops'"
description: "This page documents the 'bpf_map__attach_struct_ops' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__attach_struct_ops`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_MAP_TYPE_STRUCT_OPS`](../../../linux/map-type/BPF_MAP_TYPE_STRUCT_OPS.md) map.

## Definition

`#!c struct bpf_link *bpf_map__attach_struct_ops(const struct bpf_map *map);`

**Parameters**

- `map`: Pointer to the BPF map.

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
