---
title: "Libbpf userspace function 'bpf_program__insn_cnt'"
description: "This page documents the 'bpf_program__insn_cnt' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__insn_cnt`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Returns number of `struct bpf_insn`'s that form specified BPF program.

## Definition

`#!c size_t bpf_program__insn_cnt(const struct bpf_program *prog);`

**Parameters**

- `prog`: BPF program for which to return number of BPF instructions

## Usage

See [`bpf_program__insns`](bpf_program__insns.md) documentation for notes on how libbpf can change instructions and their count during different phases of `bpf_object` lifetime.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
