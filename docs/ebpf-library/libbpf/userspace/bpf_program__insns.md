---
title: "Libbpf userspace function 'bpf_program__insns'"
description: "This page documents the 'bpf_program__insns' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__insns`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Gives read-only access to BPF program's underlying BPF instructions.

## Definition

`#!c const struct bpf_insn *bpf_program__insns(const struct bpf_program *prog);`

**Parameters**

- `prog`: BPF program for which to return instructions

**Return**

A pointer to an array of BPF instructions that belong to the specified BPF program

## Usage

Returned pointer is always valid and not `NULL`. Number of `struct bpf_insn` pointed to can be fetched using [`bpf_program__insn_cnt`](bpf_program__insn_cnt.md) API.

Keep in mind, libbpf can modify and append/delete BPF program's instructions as it processes BPF object file and prepares everything for uploading into the kernel. So depending on the point in BPF object lifetime, [`bpf_program__insns`](bpf_program__insns.md) can return different sets of instructions. As an example, during BPF object load phase BPF program instructions will be CO-RE-relocated, BPF subprograms instructions will be appended, `ldimm64` instructions will have file descriptors embedded, etc. So instructions returned before [`bpf_object__load`](bpf_object__load.md) and after it might be quite different.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
