---
title: "Libbpf userspace function 'btf__new_empty_split'"
description: "This page documents the 'btf__new_empty_split' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__new_empty_split`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.3.0](https://github.com/libbpf/libbpf/releases/tag/v0.3.0)
<!-- [/LIBBPF_TAG] -->

Creates an unpopulated BTF object from an ELF BTF section except with a base BTF on top of which split BTF should be based

## Definition

`#!c struct btf *btf__new_empty_split(struct btf *base_btf);`

**Parameters**

- `base_btf`: The base BTF to use for the split BTF.

**Return**

New BTF object instance which has to be eventually freed with [`btf__free()`](btf__free.md)

On error, error-code-encoded-as-pointer is returned, not a `NULL`. To extract error code from such a pointer [`libbpf_get_error`](libbpf_get_error.md) should be used. If [`libbpf_set_strict_mode(LIBBPF_STRICT_CLEAN_PTRS)`](libbpf_set_strict_mode.md) is enabled, `NULL` is returned on error instead. In both cases thread-local [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) variable is always set to error code as well.

## Usage

If `base_btf` is `NULL`, `btf__new_empty_split` is equivalent to [`btf__new_empty`](btf__new_empty.md) and creates non-split BTF.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
