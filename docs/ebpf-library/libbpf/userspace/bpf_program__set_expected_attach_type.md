---
title: "Libbpf userspace function 'bpf_program__set_expected_attach_type'"
description: "This page documents the 'bpf_program__set_expected_attach_type' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__set_expected_attach_type`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Sets the attach type of the passed BPF program. This is used for auto-detection of attachment when programs are loaded.

## Definition

`#!c int bpf_program__set_expected_attach_type(struct bpf_program *prog, enum bpf_attach_type type);`

**Parameters**

- `prog`: BPF program to set the attach type for
- `type`: attach type to set the BPF map to have

**Return**

Error code; or `0` if no error. An error occurs if the object is already loaded.

## Usage

This must be called before the BPF object is loaded, otherwise it has no effect and an error is returned.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
