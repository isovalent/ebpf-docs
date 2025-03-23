---
title: "Libbpf userspace function 'bpf_link_get_fd_by_id'"
description: "This page documents the 'bpf_link_get_fd_by_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link_get_fd_by_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_LINK_GET_FD_BY_ID`](../../../linux/syscall/BPF_LINK_GET_FD_BY_ID.md) syscall command.

## Definition

`#!c int bpf_link_get_fd_by_id(__u32 id);`

**Parameters**

- `id`: BPF link ID

**Return**

`>0`, a file descriptor for the BPF link; negative error code, otherwise

## Usage

This function returns the file descriptor of the BPF link with the given `id`. This causes the current process to hold a reference to the link, preventing the kernel from unloading it. The file descriptor should be closed with [`close`](https://man7.org/linux/man-pages/man2/close.2.html) when it is no longer needed.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
