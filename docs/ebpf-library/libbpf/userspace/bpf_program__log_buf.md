---
title: "Libbpf userspace function 'bpf_program__log_buf'"
description: "This page documents the 'bpf_program__log_buf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__log_buf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Get the log buffer of a BPF program.

## Definition

`#!c const char *bpf_program__log_buf(const struct bpf_program *prog, size_t *log_size);`

**Parameters**

- `prog`: BPF program to get the log buffer of
- `log_size`: Pointer to a `size_t` variable to which the size of the log will be written

**Return**

Pointer to the log buffer of the BPF program. The log buffer is a null-terminated string containing the verifier log.

## Usage

Get the log buffer of the BPF program, which will contain the verifier log after loading. This log is important to understand why a BPF program failed when its rejected by the verifier.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
