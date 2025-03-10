---
title: "Libbpf userspace function 'bpf_program__attach_perf_event'"
description: "This page documents the 'bpf_program__attach_perf_event' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_perf_event`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.4](https://github.com/libbpf/libbpf/releases/tag/v0.0.4)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_PERF_EVENT`](../../../linux/program-type/BPF_PROG_TYPE_PERF_EVENT.md) program to a perf event.

## Definition

`#!c struct bpf_link * bpf_program__attach_perf_event(const struct bpf_program *prog, int pfd);`

**Parameters**

- `prog`: BPF program to attach
- `pfd`: File descriptor of the perf event to attach to

## Usage

The `pfd` is obtained by first creating a perf event using the [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall.

This function calls [`bpf_program__attach_perf_event_opts`](bpf_program__attach_perf_event_opts.md) with default values.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
