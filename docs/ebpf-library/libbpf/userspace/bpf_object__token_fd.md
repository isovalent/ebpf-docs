---
title: "Libbpf userspace function 'bpf_object__token_fd'"
description: "This page documents the 'bpf_object__token_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__token_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

is an accessor for BPF token FD associated with BPF object.

## Definition

`#!c int bpf_object__token_fd(const struct bpf_object *obj);`

**Parameters**

- `obj`: Pointer to a valid BPF object

**Return**

BPF token FD or `-1`, if it wasn't set

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
