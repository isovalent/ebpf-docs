---
title: "Libbpf userspace function 'bpf_linker__new_fd'"
description: "This page documents the 'bpf_linker__new_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_linker__new_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.6.0](https://github.com/libbpf/libbpf/releases/tag/v1.6.0)
<!-- [/LIBBPF_TAG] -->

Create a new `struct bpf_linker` object from a file.

## Definition

`#!c struct bpf_linker *bpf_linker__new_fd(int fd, struct bpf_linker_opts *opts);`

**Parameters**

- `fd`: The file descriptor of the first file to load the BPF object from.
- `opts`: A pointer to a `struct bpf_linker_opts` object that contains options for the linker.

**Return**

A pointer to a new `struct bpf_linker` object. On error, `NULL` is returned and [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set.

### `struct bpf_linker_opts`

```c
struct bpf_linker_opts {
    /* size of this struct, for forward/backward compatibility */
    size_t sz;
};
```

## Usage

This function creates a new BPF linker. The linker is used to link statically link multiple BPF objects together into a single ELF file.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
