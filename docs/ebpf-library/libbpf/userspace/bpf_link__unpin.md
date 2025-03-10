---
title: "Libbpf userspace function 'bpf_link__unpin'"
description: "This page documents the 'bpf_link__unpin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link__unpin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Unpins the BPF link from a file in the BPFFS specified by a path. This decrements the links reference count.

## Definition

`#!c int bpf_link__unpin(struct bpf_link *link);`

**Parameters**

- `prog`: BPF program to unpin
- `path`: file path to the pin in a BPF file system

**Return**

`0`, on success; negative error code, otherwise

## Usage

The file pinning the BPF link can also be unlinked by a different process in which case this function will return an error.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
