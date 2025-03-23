---
title: "Libbpf userspace function 'bpf_linker__finalize'"
description: "This page documents the 'bpf_linker__finalize' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_linker__finalize`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)
<!-- [/LIBBPF_TAG] -->

Finalize the linker and write the linked ELF file to the path specified in [`bpf_linker__new`](bpf_linker__new.md).

## Definition

`#!c int bpf_linker__finalize(struct bpf_linker *linker);`

**Parameters**

- `linker`: pointer to a `struct bpf_linker` object

**Return**

`0`, on success; On error negative error code, and and [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set to the error code.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
