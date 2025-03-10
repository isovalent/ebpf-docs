---
title: "Libbpf userspace function 'bpf_program__attach_lsm'"
description: "This page documents the 'bpf_program__attach_lsm' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_lsm`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_LSM`](../../../linux/program-type/BPF_PROG_TYPE_LSM.md) program.

## Definition

`#!c struct bpf_link * bpf_program__attach_lsm(const struct bpf_program *prog);`

**Parameters**

- `prog`: BPF program to attach

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
