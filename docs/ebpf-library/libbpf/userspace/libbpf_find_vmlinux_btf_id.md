---
title: "Libbpf userspace function 'libbpf_find_vmlinux_btf_id'"
description: "This page documents the 'libbpf_find_vmlinux_btf_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_find_vmlinux_btf_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)
<!-- [/LIBBPF_TAG] -->

Find the BTF ID of a kernel symbol.

## Definition

`#!c int libbpf_find_vmlinux_btf_id(const char *name, enum bpf_attach_type attach_type);`

**Parameters**

- `name`: The kernel symbol name.
- `attach_type`: The attach type.

**Return**

The BTF ID of the kernel symbol, or a negative error code on failure.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
