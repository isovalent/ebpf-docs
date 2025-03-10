---
title: "Libbpf userspace function 'bpf_program__unpin'"
description: "This page documents the 'bpf_program__unpin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__unpin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

[Unpins](../../../linux/concepts/pinning.md) the BPF program from a file in the BPFFS specified by a path. 

## Definition

`#!c int bpf_program__unpin(struct bpf_program *prog, const char *path);`

**Parameters**

- `prog`: BPF program to unpin
- `path`: file path to the pin in a BPF file system

**Return**

`0`, on success; negative error code, otherwise

## Usage

Unpinning a program decrements the programs reference count. The file pinning the BPF program can also be unlinked by a different process in which case this function will return an error.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
