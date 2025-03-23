---
title: "Libbpf userspace function 'bpf_obj_get_info_by_fd'"
description: "This page documents the 'bpf_obj_get_info_by_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_obj_get_info_by_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_OBJ_GET_INFO_BY_FD`](../../../linux/syscall/BPF_OBJ_GET_INFO_BY_FD.md) syscall command.

## Definition

`#!c int bpf_obj_get_info_by_fd(int bpf_fd, void *info, __u32 *info_len);`

**Parameters**

- `bpf_fd`: file descriptor of the BPF object
- `info`: buffer to store the information about the BPF object
- `info_len`: size of the `info` buffer

**Return**

`>0`, amount of information written to the `info` buffer; negative error code, otherwise

## Usage

This function allows you to retrieve information about a BPF object by its file descriptor. The information will be written to `info`, the structure of the info depends on the underlying object type.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
