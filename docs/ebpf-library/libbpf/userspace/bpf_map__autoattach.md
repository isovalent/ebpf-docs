---
title: "Libbpf userspace function 'bpf_map__autoattach'"
description: "This page documents the 'bpf_map__autoattach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__autoattach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Returns whether BPF map is configured to auto-attach during BPF skeleton attach phase.

## Definition

`#!c bool bpf_map__autoattach(const struct bpf_map *map);`

**Parameters**

- `map`: the BPF map instance

**Return**

`true` if map is set to auto-attach during skeleton attach phase; `false`, otherwise

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
