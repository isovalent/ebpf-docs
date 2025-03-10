---
title: "Libbpf userspace function 'bpf_program__attach'"
description: "This page documents the 'bpf_program__attach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

This is a generic function for attaching a BPF program based on auto-detection of program type, attach type, and extra parameters, where applicable.

## Definition

`#!c struct bpf_link * bpf_program__attach(const struct bpf_program *prog);`

**Parameters**

- `prog`: BPF program to attach

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

This is supported for:
  - [kprobe/kretprobe](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) (depends on [`SEC`](../ebpf/SEC.md) definition)
  - [uprobe/uretprobe](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) (depends on [`SEC`](../ebpf/SEC.md) definition)
  - [tracepoint](../../../linux/program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
  - [raw tracepoint](../../../linux/program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
  - [tracing programs](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md) (typed `raw TP`/`fentry`/`fexit`/`fmod_ret`)

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
