---
title: "Libbpf userspace function 'btf__new'"
description: "This page documents the 'btf__new' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__new`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Creates a new instance of a BTF object from the raw bytes of an ELF's BTF section

## Definition

`#!c struct btf *btf__new(const void *data, __u32 size);`

**Parameters**

- `data`: raw bytes
- `size`: number of bytes passed in `data`

**Return**

New BTF object instance which has to be eventually freed with [`btf__free`](btf__free.md)

On error, error-code-encoded-as-pointer is returned, not a `NULL`. To extract error code from such a pointer [`libbpf_get_error`](libbpf_get_error.md) should be used. If [`libbpf_set_strict_mode(LIBBPF_STRICT_CLEAN_PTRS)`](libbpf_set_strict_mode.md) is enabled, `NULL` is returned on error instead. In both cases thread-local [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) variable is always set to error code as well.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
