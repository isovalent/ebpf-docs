---
title: "Libbpf userspace function 'bpf_prog_get_fd_by_id'"
description: "This page documents the 'bpf_prog_get_fd_by_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_get_fd_by_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_PROG_GET_FD_BY_ID`](../../../linux/syscall/BPF_PROG_GET_FD_BY_ID.md) syscall command.

## Definition

`#!c int bpf_prog_get_fd_by_id(__u32 id);`

**Parameters**

- `id`: BPF program ID

**Return**

`>0`, a file descriptor for the BPF program; negative error code, otherwise

## Usage

This function returns a file descriptor for a BPF program by its ID. This causes the current process to hold a reference to the BPF program, preventing the kernel from unloading it. The file descriptor should be closed with [`close`](https://man7.org/linux/man-pages/man2/close.2.html) when it is no longer needed.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
