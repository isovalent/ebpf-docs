---
title: "Libbpf userspace function 'bpf_link_update'"
description: "This page documents the 'bpf_link_update' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link_update`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_LINK_UPDATE`](../../../linux/syscall/BPF_LINK_UPDATE.md) syscall command.

## Definition

`#!c int bpf_link_update(int link_fd, int new_prog_fd, const struct bpf_link_update_opts *opts);`

**Parameters**

- `link_fd`: file descriptor of the link to update
- `new_prog_fd`: file descriptor of the new program to attach
- `opts`: options for configuring the update

**Return**

`0`, on success; negative error code, otherwise

## Usage

This function should only be used if you require specific control over this process. In most cases, the [`bpf_program__attach`](bpf_program__attach.md) or specific `bpf_program__attach_*` functions should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
