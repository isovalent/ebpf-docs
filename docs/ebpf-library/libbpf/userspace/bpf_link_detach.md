---
title: "Libbpf userspace function 'bpf_link_detach'"
description: "This page documents the 'bpf_link_detach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link_detach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_LINK_DETACH`](../../../linux/syscall/BPF_LINK_DETACH.md) syscall command.

## Definition

`#!c int bpf_link_detach(int link_fd);`

**Parameters**

- `link_fd`: file descriptor of the link to detach

**Return**

`0`, on success; `-errno`, on error.

## Usage

This function should only be used if you require specific control over this process. In most cases, the [`bpf_link__detach`](bpf_link__detach.md) or [`bpf_link__destroy`](bpf_link__destroy.md) function should be used.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
