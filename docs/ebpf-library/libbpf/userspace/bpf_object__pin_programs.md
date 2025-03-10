---
title: "Libbpf userspace function 'bpf_object__pin_programs'"
description: "This page documents the 'bpf_object__pin_programs' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__pin_programs`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Pins each program contained within the BPF object at the passed directory.

## Definition

`#!c int bpf_object__pin_programs(struct bpf_object *obj, const char *path);`

**Parameters**

- `obj`: Pointer to a valid BPF object
- `path`: A directory where programs should be pinned.

**Return**

`0`, on success; negative error code, otherwise

## Usage

Each program will be pinned in this directory, `path` must be set to a valid directory. The name of the pin-file will be the name of the program.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
