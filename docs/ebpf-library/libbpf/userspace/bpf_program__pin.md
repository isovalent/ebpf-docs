---
title: "Libbpf userspace function 'bpf_program__pin'"
description: "This page documents the 'bpf_program__pin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__pin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

[Pins](../../../linux/concepts/pinning.md) the BPF program to a file in the BPFFS specified by a path.

## Definition

`#!c int bpf_program__pin(struct bpf_program *prog, const char *path);`

**Parameters**

- `prog`: BPF program to pin, must already be loaded
- `path`: file path in a BPF file system

**Return**

`0`, on success; negative error code, otherwise

## Usage

Pinning a program increments the programs reference count, allowing it to stay loaded after the process which loaded it has exited.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
