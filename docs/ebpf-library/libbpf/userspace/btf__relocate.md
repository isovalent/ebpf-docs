---
title: "Libbpf userspace function 'btf__relocate'"
description: "This page documents the 'btf__relocate' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__relocate`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Will check the split BTF `btf` for references to base BTF kinds, and verify those references are compatible with `base_btf`; if they are, `btf` is adjusted such that is re-parented to `base_btf` and type ids and strings are adjusted to accommodate this.

## Definition

`#!c int btf__relocate(struct btf *btf, const struct btf *base_btf);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `base_btf`: pointer to a `struct btf` object

**Return**

If successful, 0 is returned and `btf` now has `base_btf` as its base.

A negative value is returned on error and the thread-local [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) variable is set to the error code as well.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
