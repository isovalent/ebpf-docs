---
title: "Libbpf userspace function 'bpf_object__btf'"
description: "This page documents the 'bpf_object__btf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__btf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.2](https://github.com/libbpf/libbpf/releases/tag/v0.0.2)
<!-- [/LIBBPF_TAG] -->

Returns the BTF of the BPF object.

## Definition

`#!c struct btf *bpf_object__btf(const struct bpf_object *obj);`

**Parameter**

- `obj`: Pointer to a valid BPF object

**Returns**

A pointer to the BTF of the BPF object; `NULL` if the BPF object has no BTF.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
