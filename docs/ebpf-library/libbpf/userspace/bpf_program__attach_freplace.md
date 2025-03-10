---
title: "Libbpf userspace function 'bpf_program__attach_freplace'"
description: "This page documents the 'bpf_program__attach_freplace' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_freplace`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_EXT`](../../../linux/program-type/BPF_PROG_TYPE_EXT.md) program to a global eBPF function, thereby replacing it.

## Definition

`#!c struct bpf_link * bpf_program__attach_freplace(const struct bpf_program *prog, int target_fd, const char *attach_func_name);`

**Parameters**

- `prog`: BPF program to attach
- `target_fd`: file descriptor of the eBPF program containing the global function to replace
- `attach_func_name`: name of the global function to replace

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
