---
title: "Libbpf userspace function 'bpf_raw_tracepoint_open'"
description: "This page documents the 'bpf_raw_tracepoint_open' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_raw_tracepoint_open`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_RAW_TRACEPOINT_OPEN`](../../../linux/syscall/BPF_RAW_TRACEPOINT_OPEN.md) syscall command.

## Definition

`#!c int bpf_raw_tracepoint_open(const char *name, int prog_fd);`

**Parameters**

- `name`: name of the raw tracepoint
- `prog_fd`: BPF program file descriptor

**Return**

`>0`, file descriptor of the raw tracepoint; negative error code, otherwise

## Usage

This function should only be used if you need precise control over the raw tracepoint opening process. In most cases the [`bpf_program__attach_raw_tracepoint`](bpf_program__attach_raw_tracepoint.md) or [`bpf_program__attach_raw_tracepoint_opts`](bpf_program__attach_raw_tracepoint_opts.md) function should be used instead.


### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
