---
title: "Libbpf userspace function 'bpf_obj_get'"
description: "This page documents the 'bpf_obj_get' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_obj_get`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_OBJ_GET`](../../../linux/syscall/BPF_OBJ_GET.md) syscall command.

## Definition

`#!c int bpf_obj_get(const char *pathname);`

**Parameters**

- `pathname`: path to the object to retrieve

**Return**

`>0`, file descriptor of the object; negative error code, otherwise

## Usage

This function should only be used if you need precise control over the object retrieval process. In most cases the [`bpf_object__open`](bpf_object__open.md) or similar high level API functions should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
