---
title: "Libbpf userspace function 'bpf_link__disconnect'"
description: "This page documents the 'bpf_link__disconnect' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link__disconnect`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Release "ownership" of underlying BPF resource

## Definition

`#!c void bpf_link__disconnect(struct bpf_link *link);`

**Parameters**

- `link`: BPF link to disconnect

## Usage

Disconnected link, when destructed through [`bpf_link__destroy`](bpf_link__destroy.md) call won't attempt to detach/unregister that BPF resource. This is useful in situations where, say, attached BPF program has to outlive userspace program that attached it in the system. Depending on type of BPF program, though, there might be additional steps (like pinning BPF program in BPFFS) necessary to ensure exit of userspace program doesn't trigger automatic detachment and clean up inside the kernel.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
