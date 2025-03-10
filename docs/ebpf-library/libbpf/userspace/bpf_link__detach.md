---
title: "Libbpf userspace function 'bpf_link__detach'"
description: "This page documents the 'bpf_link__detach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link__detach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Detach a BPF link from the kernel.

## Definition

`#!c int bpf_link__detach(struct bpf_link *link);`

**Parameters**

- `link`: BPF link to detach

## Usage

This detaches the BPF link from the kernel but does delete the pin or free resources associated.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
