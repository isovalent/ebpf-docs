---
title: "Libbpf userspace function 'bpf_obj_pin'"
description: "This page documents the 'bpf_obj_pin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_obj_pin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_OBJ_PIN`](../../../linux/syscall/BPF_OBJ_PIN.md) syscall command.

## Definition

`#!c int bpf_obj_pin(int fd, const char *pathname);`

**Parameters**

- `fd`: file descriptor of the object to pin
- `pathname`: path to the directory where the object will be pinned

**Return**

`0`, on success; negative error code, otherwise

## Usage

This function should only be used if you need precise control over the object pinning process. In most cases the [`bpf_object__pin`](bpf_object__pin.md) function should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
