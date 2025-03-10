---
title: "Libbpf userspace function 'bpf_program__attach_xdp'"
description: "This page documents the 'bpf_program__attach_xdp' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_xdp`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_XDP`](../../../linux/program-type/BPF_PROG_TYPE_XDP.md) program to a network interface.

## Definition

`#!c struct bpf_link * bpf_program__attach_xdp(const struct bpf_program *prog, int ifindex);`

**Parameters**

- `prog`: BPF program to attach
- `ifindex`: index of the network interface to attach the program to

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
