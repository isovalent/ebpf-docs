---
title: "Libbpf userspace function 'bpf_program__unload'"
description: "This page documents the 'bpf_program__unload' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__unload`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Unload a BPF program.

## Definition

`#!c void bpf_program__unload(struct bpf_program *prog);`

## Usage

This function closes the file descriptor of the BPF program, decrementing its refcount. This will cause the program to be unloaded, if there are no other references to it.

A program might still stay loaded in the kernel if anything else maintains a reference like:

* A BPF link
* A subsystem that maintains a link (such as XDP or TC)
* A BPFFS pin

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
