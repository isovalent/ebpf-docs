---
title: "Libbpf userspace function 'bpf_program__set_log_buf'"
description: "This page documents the 'bpf_program__set_log_buf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__set_log_buf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Set the log buffer of a BPF program.

## Definition

`#!c int bpf_program__set_log_buf(struct bpf_program *prog, char *log_buf, size_t log_size);`

**Parameters**

- `prog`: BPF program to set the log buffer for
- `log_buf`: Buffer to set as the log buffer
- `log_size`: Size of the log buffer

**Return**

error code; or `0` if no error. An error occurs if the object is already loaded.

## Usage

Set the log buffer. The verifier will write its log into this buffer when loading the program. This log is important to understand why a BPF program failed when its rejected by the verifier.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
