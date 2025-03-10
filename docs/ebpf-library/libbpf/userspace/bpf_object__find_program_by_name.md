---
title: "Libbpf userspace function 'bpf_object__find_program_by_name'"
description: "This page documents the 'bpf_object__find_program_by_name' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__find_program_by_name`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Query a BPF object for a program by name.

## Definition

`#!c struct bpf_program * bpf_object__find_program_by_name(const struct bpf_object *obj, const char *name);`

**Parameters**

- `obj`: Pointer to a valid BPF object
- `name`: Name of the program to find

**Returns**

A pointer to the program with the given name, or `NULL` if no program with that name is found.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
