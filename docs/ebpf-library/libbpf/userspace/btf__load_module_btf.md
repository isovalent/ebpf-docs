---
title: "Libbpf userspace function 'btf__load_module_btf'"
description: "This page documents the 'btf__load_module_btf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__load_module_btf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Load BTF object of a kernel module.

## Definition

`#!c struct btf *btf__load_module_btf(const char *module_name, struct btf *vmlinux_btf);`

**Parameters**

- `module_name`: name of the module
- `vmlinux_btf`: pointer to a `struct btf` object gotten from [`btf__load_vmlinux_btf`](btf__load_vmlinux_btf.md)

**Return**

Return a pointer to a `struct btf` object on success, or `NULL` on failure. The caller is responsible for freeing the returned object with [`btf__free`](btf__free.md).

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
