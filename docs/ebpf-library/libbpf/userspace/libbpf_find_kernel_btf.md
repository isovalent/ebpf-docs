---
title: "Libbpf userspace function 'libbpf_find_kernel_btf'"
description: "This page documents the 'libbpf_find_kernel_btf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_find_kernel_btf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

This function has been renamed to [`btf__load_vmlinux_btf`](btf__load_vmlinux_btf.md).

## Definition

`#!c struct btf *libbpf_find_kernel_btf(void);`

## Usage

This functions is simply an alias for [`btf__load_vmlinux_btf`](btf__load_vmlinux_btf.md) as of v0.5.0.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
