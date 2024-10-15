---
title: "Libbpf userspace function 'libbpf_strerror'"
description: "This page documents the 'libbpf_strerror' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_strerror`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Convert an error code into a human-readable string.

## Definition

`#!c int libbpf_strerror(int err, char *buf, size_t size)`

`err` - error code to convert into a string

`buf` - buffer to store the string

`size` - size of the buffer

## Usage

This function converts an error code into a human-readable string. It is useful for debugging and logging purposes.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
