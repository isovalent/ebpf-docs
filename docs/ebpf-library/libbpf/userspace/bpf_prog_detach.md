---
title: "Libbpf userspace function 'bpf_prog_detach'"
description: "This page documents the 'bpf_prog_detach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_detach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_PROG_DETACH`](../../../linux/syscall/BPF_PROG_DETACH.md) syscall command.

## Definition

`#!c int bpf_prog_detach(int attachable_fd, enum bpf_attach_type type);`

**Parameters**

- `attachable_fd`: file descriptor of the attachable object
- `type`: type of the attachment

**Return**

`0`, on success; negative error code, otherwise

## Usage

This function should only be used for specific program types that need to be detached via the `BPF_PROG_DETACH` syscall command and you need specific control over this process. In most cases, the [`bpf_link__detach`](bpf_link__detach.md) or [`bpf_link__destroy`](bpf_link__destroy.md) function should be used.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
