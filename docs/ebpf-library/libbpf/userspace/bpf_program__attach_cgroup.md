---
title: "Libbpf userspace function 'bpf_program__attach_cgroup'"
description: "This page documents the 'bpf_program__attach_cgroup' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_cgroup`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_CGROUP_*`](../../../linux/program-type/index.md#cgroup-program-types) program.

## Definition

`#!c struct bpf_link * bpf_program__attach_cgroup(const struct bpf_program *prog, int cgroup_fd);`

**Parameters**

- `prog`: BPF program to attach
- `cgroup_fd`: file descriptor of the cgroup to attach the program to

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
