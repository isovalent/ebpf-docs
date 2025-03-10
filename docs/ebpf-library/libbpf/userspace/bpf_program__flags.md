---
title: "Libbpf userspace function 'bpf_program__flags'"
description: "This page documents the 'bpf_program__flags' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__flags`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Get flags which with the BPF program is or will be loaded.

## Definition

`#!c __u32 bpf_program__flags(const struct bpf_program *prog);`

**Parameters**

- `prog`: BPF program to get the flags of.

**Return**

The [flags](../../../linux/syscall/BPF_PROG_LOAD.md#flags) of the BPF program. Not every flag can be set for every program type.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
