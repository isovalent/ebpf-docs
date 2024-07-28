---
title: "Helper Function 'bpf_sock_from_file'"
description: "This page documents the 'bpf_sock_from_file' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_sock_from_file`

<!-- [FEATURE_TAG](bpf_sock_from_file) -->
[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/4f19cab76136e800a3f04d8c9aa4d8e770e3d3d8)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
If the given file represents a socket, returns the associated socket.

### Returns

A pointer to a struct socket on success or NULL if the file is not a socket.

`#!c static struct socket *(* const bpf_sock_from_file)(struct file *file) = (void *) 162;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
