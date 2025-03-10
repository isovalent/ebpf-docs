---
title: "Libbpf userspace function 'bpf_program__attach_kprobe'"
description: "This page documents the 'bpf_program__attach_kprobe' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_kprobe`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.4](https://github.com/libbpf/libbpf/releases/tag/v0.0.4)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_KPROBE`](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) program.

## Definition

`#!c struct bpf_link * bpf_program__attach_kprobe(const struct bpf_program *prog, bool retprobe, const char *func_name);`

**Parameters**

- `prog`: BPF program to attach
- `retprobe`: `true` if attaching a return probe, `false` if attaching an entry probe.
- `func_name`: name of the kernel function to attach the probe to.

## Usage

Force libbpf to attach kprobe/uprobe in specific mode, `-ENOTSUP` will be returned if it is not supported by the kernel.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
