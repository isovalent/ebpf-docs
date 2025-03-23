---
title: "Libbpf userspace function 'btf__load_from_kernel_by_id'"
description: "This page documents the 'btf__load_from_kernel_by_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__load_from_kernel_by_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Load BTF object from the kernel by its ID.

## Definition

`#!c struct btf *btf__load_from_kernel_by_id(__u32 id);`

**Parameters**

- `id`: ID of the BTF object

**Return**

Return a pointer to a `struct btf` object on success, or `NULL` on failure. The caller is responsible for freeing the returned object with [`btf__free`](btf__free.md).

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
