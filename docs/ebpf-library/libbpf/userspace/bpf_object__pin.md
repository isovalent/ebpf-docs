---
title: "Libbpf userspace function 'bpf_object__pin'"
description: "This page documents the 'bpf_object__pin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__pin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Pin all maps and programs contained within the BPF object to the passed directory.

## Definition

`#!c int bpf_object__pin(struct bpf_object *object, const char *path);`

**Parameters**

- `obj`: Pointer to a valid BPF object
- `path`: A directory where programs should be pinned.

**Return**

`0`, on success; negative error code, otherwise

## Usage

Calls both [`bpf_object__pin_maps`](bpf_object__pin_maps.md) and [`bpf_object__pin_programs`](bpf_object__pin_programs.md) with the same `path` argument.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
