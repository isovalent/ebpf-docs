---
title: "Libbpf userspace function 'btf__distill_base'"
description: "This page documents the 'btf__distill_base' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__distill_base`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Creates new versions of the split BTF `src_btf` and its base BTF.

## Definition

`#!c int btf__distill_base(const struct btf *src_btf, struct btf **new_base_btf, struct btf **new_split_btf);`

**Parameters**

- `src_btf`: The split BTF to distill the base BTF from.
- `new_base_btf`: The new base BTF distilled from `src_btf`.
- `new_split_btf`: The new split BTF distilled from `src_btf`.

**Return**

If successful, 0 is returned and `new_base_btf` and `new_split_btf` will point at new base/split BTF. Both the new split and its associated new base BTF must be freed by the caller.

A negative value is returned on error and the thread-local [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) variable is set to the error code as well.

## Usage

The new base BTF will only contain the types needed to improve robustness of the split BTF to small changes in base BTF. When that split BTF is loaded against a (possibly changed) base, this distilled base BTF will help update references to that (possibly changed) base BTF.

Both the new split and its associated new base BTF must be freed by the caller.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
