---
title: "Libbpf userspace function 'bpf_map__is_pinned'"
description: "This page documents the 'bpf_map__is_pinned' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__is_pinned`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)
<!-- [/LIBBPF_TAG] -->

Tells the caller whether or not the passed map has been pinned via a 'pin' file.

## Definition

`#!c bool bpf_map__is_pinned(const struct bpf_map *map);`

**Parameters**

- `map`: The bpf_map

**Return**

`true`, if the map is pinned; `false`, otherwise

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
