---
title: "Libbpf userspace function 'bpf_object__name'"
description: "This page documents the 'bpf_object__name' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__name`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Get the name of the BPF object.

## Definition

`#!c const char *bpf_object__name(const struct bpf_object *obj);`

**Returns**

A string containing the name of the BPF object.

## Usage

The name of the object can be set via options when opening the object. This allows you to get that identifier later on.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
