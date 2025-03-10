---
title: "Libbpf userspace function 'bpf_object__open_file'"
description: "This page documents the 'bpf_object__open_file' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__open_file`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)
<!-- [/LIBBPF_TAG] -->

creates a `struct bpf_object` by opening the BPF ELF object file pointed to by the passed path and loading it into memory.

## Definition

`#!c struct bpf_object * bpf_object__open_file(const char *path, const struct bpf_object_open_opts *opts);`

**Parameters**

- `path`: BPF object file path
- `opts`: options for how to load the bpf object, this parameter is optional and can be set to `NULL`

**Return**

pointer to the new `struct bpf_object`; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
