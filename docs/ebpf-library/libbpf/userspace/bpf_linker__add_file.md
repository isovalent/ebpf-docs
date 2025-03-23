---
title: "Libbpf userspace function 'bpf_linker__add_file'"
description: "This page documents the 'bpf_linker__add_file' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_linker__add_file`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)
<!-- [/LIBBPF_TAG] -->

Add a file to the linker.

## Definition

`#!c int bpf_linker__add_file(struct bpf_linker *linker, const char *filename, const struct bpf_linker_file_opts *opts);`

**Parameters**

- `linker`: pointer to a `struct bpf_linker` object
- `filename`: name of the file to add to the linker
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

Add another file to the linker. This file will be linked together with the other files that have been added to the linker. The linker is used to link statically link multiple BPF objects together into a single ELF file.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
