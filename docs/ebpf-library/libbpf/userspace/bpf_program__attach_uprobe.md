---
title: "Libbpf userspace function 'bpf_program__attach_uprobe'"
description: "This page documents the 'bpf_program__attach_uprobe' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_uprobe`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.4](https://github.com/libbpf/libbpf/releases/tag/v0.0.4)
<!-- [/LIBBPF_TAG] -->

Attaches a BPF program to the userspace function which is found by binary path and offset. You can optionally specify a particular process to attach to. You can also optionally attach the program to the function exit instead of entry.

## Definition

`#!c struct bpf_link * bpf_program__attach_uprobe(const struct bpf_program *prog, bool retprobe, pid_t pid, const char *binary_path, size_t func_offset);`

**Parameters**

- `prog`: BPF program to attach
- `retprobe`: Attach to function exit
- `pid`: Process ID to attach the uprobe to, 0 for self (own process), `-1` for all processes
- `binary_path`: Path to binary that contains the function symbol
- `func_offset`: Offset within the binary of the function symbol

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
