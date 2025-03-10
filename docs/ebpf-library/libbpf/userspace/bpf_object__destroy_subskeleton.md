---
title: "Libbpf userspace function 'bpf_object__destroy_subskeleton'"
description: "This page documents the 'bpf_object__destroy_subskeleton' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__destroy_subskeleton`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Destroy a sub-skeleton.

## Definition

`#!c void bpf_object__destroy_subskeleton(struct bpf_object_subskeleton *s);`

## Usage

This function destroys the resources associated with a sub-skeleton. It does not detach programs.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
