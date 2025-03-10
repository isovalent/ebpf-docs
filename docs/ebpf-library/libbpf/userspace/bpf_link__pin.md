---
title: "Libbpf userspace function 'bpf_link__pin'"
description: "This page documents the 'bpf_link__pin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link__pin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Pins the BPF link to a file in the BPFFS specified by a path.

## Definition

`#!c int bpf_link__pin(struct bpf_link *link, const char *path);`

**Parameters**

- `link`: BPF link to pin, must already be loaded
- `path`: file path in a BPF file system

**Return**

`0`, on success; negative error code, otherwise

## Usage

This increments the links reference count, allowing it to stay loaded after the process which loaded it has exited.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
