---
title: "Libbpf userspace function 'bpf_iter_create'"
description: "This page documents the 'bpf_iter_create' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_iter_create`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_ITER_CREATE`](../../../linux/syscall/BPF_ITER_CREATE.md) syscall command.

## Definition

`#!c int bpf_iter_create(int link_fd);`

**Parameters**

- `link_fd`: file descriptor of the link to create the iterator for

**Return**

`0`, on success; negative error code, otherwise

## Usage

This function should only be used if you require specific control over the iterator creation process. In most cases, the [`bpf_program__attach_iter`](bpf_program__attach_iter.md) function should be used.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
