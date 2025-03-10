---
title: "Libbpf userspace function 'bpf_object__load'"
description: "This page documents the 'bpf_object__load' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__load`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

loads BPF object into kernel.

## Definition

`#!c int bpf_object__load(struct bpf_object *obj);`

**Parameters**

- `obj`: Pointer to a valid BPF object instance returned by [`bpf_object__open`](bpf_object__open.md), [`bpf_object__open_file`](bpf_object__open_file.md), or [`bpf_object__open_mem`](bpf_object__open_mem.md)

**Return**

`0`, on success; negative error code, otherwise, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
