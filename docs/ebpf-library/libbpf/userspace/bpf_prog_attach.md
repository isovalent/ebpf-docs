---
title: "Libbpf userspace function 'bpf_prog_attach'"
description: "This page documents the 'bpf_prog_attach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_attach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_PROG_ATTACH`](../../../linux/syscall/BPF_PROG_ATTACH.md) syscall command.

## Definition

`#!c int bpf_prog_attach(int prog_fd, int attachable_fd, enum bpf_attach_type type, unsigned int flags);`

**Parameters**

- `prog_fd`: file descriptor of the program to attach
- `attachable_fd`: file descriptor of the attachable object
- `type`: type of the attachment
- `flags`: flags for the attachment

**Return**

`0`, on success; negative error code, otherwise

## Usage

This function should only be used for specific program types that need to be attached via the `BPF_PROG_ATTACH` syscall command and you need specific control over this process. In most cases, the [`bpf_program__attach`](bpf_program__attach.md) or specific `bpf_program__attach_*` functions should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
