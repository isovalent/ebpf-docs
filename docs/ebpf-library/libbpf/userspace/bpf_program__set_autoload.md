---
title: "Libbpf userspace function 'bpf_program__set_autoload'"
description: "This page documents the 'bpf_program__set_autoload' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__set_autoload`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Changes the auto-load status of a BPF program.

## Definition

`#!c int bpf_program__set_autoload(struct bpf_program *prog, bool autoload);`

**Parameters**

- `prog`: The BPF program.
- `autoload`: The new autoload status.

**Returns**

`0` on success, or a negative error code on failure.

## Usage

This method changes the auto-load status of a BPF program. If `autoload` is `true`, the program will be loaded when the object is loaded.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
