---
title: "Libbpf userspace function 'bpf_object__gen_loader'"
description: "This page documents the 'bpf_object__gen_loader' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__gen_loader`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Instruct the BPF object to generate a loader program instead of actually loading the object.

## Definition

`#!c int bpf_object__gen_loader(struct bpf_object *obj, struct gen_loader_opts *opts);`

## Usage

This method instructs the BPF object to generate a loader program instead of actually loading the object. It records all steps it would take from userspace and translates them into a program of type [`BPF_PROG_TYPE_SYSCALL`](../../../linux/program-type/BPF_PROG_TYPE_SYSCALL.md). 

This is part of an attempt to create cryptographically signed BPF programs. BPF programs are commonly modified by the loader before being loaded into the kernel which makes it impossible to sign the ELF, since the actual bytecode to be loaded will be different.  So the idea was to generate a loader program at compile time, it could be signed.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
