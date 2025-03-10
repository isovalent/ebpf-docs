---
title: "Libbpf userspace function 'bpf_program__autoload'"
description: "This page documents the 'bpf_program__autoload' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__autoload`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Query if the program will be automatically loaded.

## Definition

`#!c bool bpf_program__autoload(const struct bpf_program *prog);`

## Usage

If `true`, the program will be loaded when the BPF object is loaded. By default this value is determined by the ELF section name. Programs in ELF sections starting with `?` are not autoloaded.

This value can be updated with the [`bpf_program__set_autoload`](bpf_program__set_autoload.md) function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
