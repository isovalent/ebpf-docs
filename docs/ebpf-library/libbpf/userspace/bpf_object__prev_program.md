---
title: "Libbpf userspace function 'bpf_object__prev_program'"
description: "This page documents the 'bpf_object__prev_program' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__prev_program`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Iterate over the programs in a BPF object in reverse order.

## Definition

`#!c struct bpf_program * bpf_object__prev_program(const struct bpf_object *obj, struct bpf_program *prog);`

**Parameters**

- `obj`: The BPF object to iterate over.
- `prog`: The current program, or `NULL` to start iteration.

**Returns**

The previous program in the object, or `NULL` if there are no more programs.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
