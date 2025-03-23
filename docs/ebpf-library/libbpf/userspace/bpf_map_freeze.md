---
title: "Libbpf userspace function 'bpf_map_freeze'"
description: "This page documents the 'bpf_map_freeze' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_freeze`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.3](https://github.com/libbpf/libbpf/releases/tag/v0.0.3)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_FREEZE`](../../../linux/syscall/BPF_MAP_FREEZE.md) syscall command.

## Definition

`#!c int bpf_map_freeze(int fd);`

**Parameters**

- `fd`: file descriptor of the map to freeze

**Return**

`0`, on success; `-errno`, on error.

## Usage

This function allows you to freeze a map, preventing any further modifications to it. This is usually done on internal maps libbpf creates for constant global data. This function should only be used if the user knows what they are doing.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
