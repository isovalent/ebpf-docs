---
title: "Libbpf userspace function 'bpf_map_get_fd_by_id'"
description: "This page documents the 'bpf_map_get_fd_by_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_get_fd_by_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_GET_FD_BY_ID`](../../../linux/syscall/BPF_MAP_GET_FD_BY_ID.md) syscall command.

## Definition

`#!c int bpf_map_get_fd_by_id(__u32 id);`

**Parameters**

- `id`: BPF map ID

**Return**

File descriptor of the BPF map with the given `id`, on success; negative error code, otherwise

## Usage

This function returns the file descriptor of the BPF map with the given `id`. This causes the current process to hold a reference to the map, preventing the kernel from unloading it. The file descriptor should be closed with [`close`](https://man7.org/linux/man-pages/man2/close.2.html) when it is no longer needed.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
