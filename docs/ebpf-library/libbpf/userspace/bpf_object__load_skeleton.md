---
title: "Libbpf userspace function 'bpf_object__load_skeleton'"
description: "This page documents the 'bpf_object__load_skeleton' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__load_skeleton`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Load a skeleton into the kernel.

## Definition

`#!c int bpf_object__load_skeleton(struct bpf_object_skeleton *s);`

**Parameters**

- `s`: The skeleton to load.

## Usage

This function loads a skeleton into the kernel. It is typically not called directly by users, but rather via the `<name>__load` or `<name>__open_and_load` function that is generated as part of the skeleton header file.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
