---
title: "Libbpf userspace function 'bpf_program__set_flags'"
description: "This page documents the 'bpf_program__set_flags' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__set_flags`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Set flags with which the BPF program will be loaded.

## Definition

`#!c int bpf_program__set_flags(struct bpf_program *prog, __u32 flags);`

**Parameters**

- `prog`: BPF program to set the flags for
- `flags`: [flags](../../../linux/syscall/BPF_PROG_LOAD.md#flags) to set the BPF program to have

**Return**

error code; or `0` if no error. An error occurs if the object is already loaded.

## Usage

This must be called before the BPF object is loaded, otherwise it has no effect and an error is returned.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
