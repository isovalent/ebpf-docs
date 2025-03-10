---
title: "Libbpf userspace function 'bpf_object__open_mem'"
description: "This page documents the 'bpf_object__open_mem' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__open_mem`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)
<!-- [/LIBBPF_TAG] -->

creates a `struct bpf_object` by reading the BPF objects raw bytes from a memory buffer containing a valid BPF ELF object file.

## Definition

`#!c struct bpf_object * bpf_object__open_mem(const void *obj_buf, size_t obj_buf_sz, const struct bpf_object_open_opts *opts);`

**Parameters**

- `obj_buf`: pointer to the buffer containing ELF file bytes
- `obj_buf_sz`: number of bytes in the buffer
- `opts`: options for how to load the bpf object

**Return**

pointer to the new `struct bpf_object`; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
