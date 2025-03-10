---
title: "Libbpf userspace function 'bpf_program__attach_sockmap'"
description: "This page documents the 'bpf_program__attach_sockmap' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_sockmap`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_SK_SKB`](../../../linux/program-type/BPF_PROG_TYPE_SK_SKB.md) or [`BPF_PROG_TYPE_SK_MSG`](../../../linux/program-type/BPF_PROG_TYPE_SK_MSG.md) program to a [`BPF_MAP_TYPE_SOCKMAP`](../../../linux/map-type/BPF_MAP_TYPE_SOCKMAP.md) or [`BPF_MAP_TYPE_SOCKHASH`](../../../linux/map-type/BPF_MAP_TYPE_SOCKHASH.md) map.

## Definition

`#!c struct bpf_link * bpf_program__attach_sockmap(const struct bpf_program *prog, int map_fd);`

**Parameters**

- `prog`: BPF program to attach
- `map_fd`: file descriptor of the [`BPF_MAP_TYPE_SOCKMAP`](../../../linux/map-type/BPF_MAP_TYPE_SOCKMAP.md) or [`BPF_MAP_TYPE_SOCKHASH`](../../../linux/map-type/BPF_MAP_TYPE_SOCKHASH.md) map.

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
