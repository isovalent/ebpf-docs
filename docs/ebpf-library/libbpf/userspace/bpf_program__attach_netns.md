---
title: "Libbpf userspace function 'bpf_program__attach_netns'"
description: "This page documents the 'bpf_program__attach_netns' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_netns`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_SK_LOOKUP`](../../../linux/program-type/BPF_PROG_TYPE_SK_LOOKUP.md) or [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../../../linux/program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md) program to a network namespace.

## Definition

`#!c struct bpf_link * bpf_program__attach_netns(const struct bpf_program *prog, int netns_fd);`

**Parameters**

- `prog`: BPF program to attach
- `netns_fd`: file descriptor of the network namespace to attach the program to

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
