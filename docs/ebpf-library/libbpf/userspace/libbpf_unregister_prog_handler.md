---
title: "Libbpf userspace function 'libbpf_unregister_prog_handler'"
description: "This page documents the 'libbpf_unregister_prog_handler' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_unregister_prog_handler`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Unregister a previously registered custom BPF program [`SEC`](../ebpf/SEC.md) handler.

## Definition

`#!c int libbpf_unregister_prog_handler(int handler_id);`

**Parameters**

- `handler_id`: handler ID returned by [`libbpf_register_prog_handler`](libbpf_register_prog_handler.md) after successful registration

**Return**

0 on success, negative error code if handler isn't found

## Usage

!!! note
    like much of global libbpf APIs (e.g., [`libbpf_set_print`](libbpf_set_print.md), [`libbpf_set_strict_mode`](libbpf_set_strict_mode.md), etc) these APIs are not thread-safe. User needs to ensure synchronization if there is a risk of running this API from multiple threads simultaneously.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
