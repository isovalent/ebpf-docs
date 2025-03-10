---
title: "Libbpf userspace function 'bpf_program__set_ifindex'"
description: "This page documents the 'bpf_program__set_ifindex' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__set_ifindex`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Set the interface index for a BPF program.

## Definition

`#!c void bpf_program__set_ifindex(struct bpf_program *prog, __u32 ifindex);`

**Parameters**

- `prog`: The BPF program.
- `ifindex`: The interface index to set.

## Usage

This method associates a BPF program with a network interface. This is only relevant for programs of type [`BPF_PROG_TYPE_XDP`](../../../linux/program-type/BPF_PROG_TYPE_XDP.md) that are offloaded to hardware. Offloading to hardware loops in the hardware driver in the verification process, so the kernel needs to know the network device the program will be attached to at load time.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
