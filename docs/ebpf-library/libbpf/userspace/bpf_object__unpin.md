---
title: "Libbpf userspace function 'bpf_object__unpin'"
description: "This page documents the 'bpf_object__unpin' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__unpin`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Unpins each program and map contained within the BPF object at the passed directory.

## Definition

`#!c int bpf_object__unpin(struct bpf_object *object, const char *path);`

**Parameters**

- `obj`: Pointer to a valid BPF object
- `path`: A directory where programs should be pinned.

**Return**

`0`, on success; negative error code, otherwise

## Usage

Calls both [`bpf_object__unpin_maps`](bpf_object__unpin_maps.md) and [`bpf_object__unpin_programs`](bpf_object__unpin_programs.md) with the same `path` argument.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
