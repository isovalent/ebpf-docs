---
title: "Libbpf userspace function 'bpf_map__set_autoattach'"
description: "This page documents the 'bpf_map__set_autoattach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_autoattach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Sets whether libbpf has to auto-attach map during BPF skeleton attach phase.

## Definition

`#!c int bpf_map__set_autoattach(struct bpf_map *map, bool autoattach);`

**Parameters**

- `map`: the BPF map instance
- `autoattach`: whether to attach map during BPF skeleton attach phase

**Return**

0 on success; negative error code, otherwise

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
