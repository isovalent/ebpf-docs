---
title: "Libbpf userspace function 'bpf_object__unpin_maps'"
description: "This page documents the 'bpf_object__unpin_maps' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__unpin_maps`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

unpins each map contained within the BPF object found in the passed directory.

## Definition

`#!c int bpf_object__unpin_maps(struct bpf_object *obj, const char *path);`

**Parameters**

- `obj`: Pointer to a valid BPF object
- `path`: A directory where pinned maps should be searched for.

**Return**

`0`, on success; negative error code, otherwise

## Usage


If `path` is `NULL` `bpf_map__unpin` (which is being used on each map) will use the `pin_path` attribute of each map. In this case, maps that don't have a `pin_path` set will be ignored.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
