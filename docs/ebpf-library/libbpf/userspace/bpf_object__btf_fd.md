---
title: "Libbpf userspace function 'bpf_object__btf_fd'"
description: "This page documents the 'bpf_object__btf_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__btf_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Returns the file descriptor of the BTF loaded into the kernel for a given BPF object.

## Definition

`#!c int bpf_object__btf_fd(const struct bpf_object *obj);`

**Parameters**

- `obj`: Pointer to a valid BPF object

**Returns**

The file descriptor of the BTF loaded into the kernel for the BPF object, or `-1` if the BPF object has no BTF or the BTF has not yet been loaded into the kernel.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
