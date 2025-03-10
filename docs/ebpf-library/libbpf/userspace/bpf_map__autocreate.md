---
title: "Libbpf userspace function 'bpf_map__autocreate'"
description: "This page documents the 'bpf_map__autocreate' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__autocreate`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Returns a boolean value indicating whether libbpf has to auto-create BPF map during BPF object load phase.

## Definition

`#!c bool bpf_map__autocreate(const struct bpf_map *map);`

**Parameters**

- `map`: Pointer to the BPF map.

**Return**

`true` if the BPF map should be auto-created, `false` otherwise.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
