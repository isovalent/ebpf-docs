---
title: "Libbpf userspace function 'bpf_link__update_program'"
description: "This page documents the 'bpf_link__update_program' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link__update_program`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Updates the BPF program associated with a BPF link.

## Definition

`#!c int bpf_link__update_program(struct bpf_link *link, struct bpf_program *prog);`

**Parameters**

- `link`: BPF link to update
- `prog`: BPF program to associate with the link

**Return**

`0`, on success; negative error code, otherwise

## Usage

This functions allows you to swap out the existing program associated with a BPF link with a new one. For cases where you do not want to miss any events that might occur between removing and re-creating a link.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
