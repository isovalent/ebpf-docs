---
title: "Libbpf userspace function 'bpf_linker__add_buf'"
description: "This page documents the 'bpf_linker__add_buf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_linker__add_buf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.6.0](https://github.com/libbpf/libbpf/releases/tag/v1.6.0)
<!-- [/LIBBPF_TAG] -->

Add a ELF file as memory buffer to the linker.

## Definition

`#!c int bpf_linker__add_buf(struct bpf_linker *linker, void *buf, size_t buf_sz, const struct bpf_linker_file_opts *opts);`

**Parameters**

- `linker`: pointer to a `struct bpf_linker` object
- `buf`: pointer to the ELF file buffer
- `buf_sz`: size of the ELF file buffer
- `opts`: pointer to a `struct bpf_linker_file_opts` object that contains options for the file

**Return**

`0`, on success; On error negative error code, and [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set to the error code.

### `struct bpf_linker_file_opts`

```c
struct bpf_linker_file_opts {
    /* size of this struct, for forward/backward compatibility */
    size_t sz;
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
