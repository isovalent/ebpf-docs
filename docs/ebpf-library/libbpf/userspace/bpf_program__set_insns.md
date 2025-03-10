---
title: "Libbpf userspace function 'bpf_program__set_insns'"
description: "This page documents the 'bpf_program__set_insns' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__set_insns`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Set BPF program's underlying BPF instructions.

## Definition

`#!c int bpf_program__set_insns(struct bpf_program *prog, struct bpf_insn *new_insns, size_t new_insn_cnt);`

**Parameters**

- `prog`: BPF program for which to return instructions
- `new_insns`: a pointer to an array of BPF instructions
- `new_insn_cnt`: number of `struct bpf_insn`'s that form
specified BPF program

**Return**

`0`, on success; negative error code, otherwise

## Usage

This function allows a user to modify or replace the BPF instructions of a BPF program right before its about to be loaded into the kernel.

!!! warning
    This is a very advanced libbpf API and users need to know what they are doing. This should be used from [`prog_prepare_load_fn`](struct-libbpf_prog_handler_opts.md#prog_prepare_load_fn) callback only.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
