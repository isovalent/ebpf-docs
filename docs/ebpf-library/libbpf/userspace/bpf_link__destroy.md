---
title: "Libbpf userspace function 'bpf_link__destroy'"
description: "This page documents the 'bpf_link__destroy' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link__destroy`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.4](https://github.com/libbpf/libbpf/releases/tag/v0.0.4)
<!-- [/LIBBPF_TAG] -->

Destroy a BPF link.

## Definition

`#!c int bpf_link__destroy(struct bpf_link *link);`

**Parameters**

- `link`: Pointer to the BPF link.

**Return**

`0` on success, `-1` on error.

## Usage

Destroying a BPF link will detach it from the kernel and free all resources associated with it.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
