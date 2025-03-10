---
title: "Libbpf userspace function 'bpf_object__find_map_by_name'"
description: "This page documents the 'bpf_object__find_map_by_name' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__find_map_by_name`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

returns BPF map of the given name, if it exists within the passed BPF object

## Definition

`#!c struct bpf_map * bpf_object__find_map_by_name(const struct bpf_object *obj, const char *name);`

**Parameters**

- `obj`: BPF object
- `name`: name of the BPF map

**Return**

BPF map instance, if such map exists within the BPF object; or `NULL` otherwise.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
