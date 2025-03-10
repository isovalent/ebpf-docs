---
title: "Libbpf userspace function 'bpf_object__attach_skeleton'"
description: "This page documents the 'bpf_object__attach_skeleton' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__attach_skeleton`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Auto attach the programs in a skeleton to hooks in the kernel.

## Definition

`#!c int bpf_object__attach_skeleton(struct bpf_object_skeleton *s);`

**Parameters**

- `s`: The skeleton to attach

## Usage

This function is typically not called directly but via the generated `<name>__attach` function. This function will auto attach any programs in the skeleton that can be auto attached (based on the section they are in). However, this is optional, and you have more control if you manually attach program. 

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
