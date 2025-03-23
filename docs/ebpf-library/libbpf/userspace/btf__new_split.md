---
title: "Libbpf userspace function 'btf__new_split'"
description: "This page documents the 'btf__new_split' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__new_split`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)
<!-- [/LIBBPF_TAG] -->

Create a new instance of a BTF object from the provided raw data bytes. It takes another BTF instance, `base_btf`, which serves as a base BTF, which is extended by types in a newly created BTF instance

## Definition

`#!c struct btf *btf__new_split(const void *data, __u32 size, struct btf *base_btf);`

**Parameters**

- `data`: raw bytes
- `size`: length of raw bytes
- `base_btf`: the base BTF object

**Return**

New BTF object instance which has to be eventually freed with [`btf__free`](btf__free.md)

On error, error-code-encoded-as-pointer is returned, not a `NULL`. To extract error code from such a pointer [`libbpf_get_error`](libbpf_get_error.md) should be used. If [`libbpf_set_strict_mode(LIBBPF_STRICT_CLEAN_PTRS)`](libbpf_set_strict_mode.md) is enabled, `NULL` is returned on error instead. In both cases thread-local [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) variable is always set to error code as well.

## Usage

If `base_btf` is `NULL`, `btf__new_split` is equivalent to [`btf__new`](btf__new.md) and creates non-split BTF.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
