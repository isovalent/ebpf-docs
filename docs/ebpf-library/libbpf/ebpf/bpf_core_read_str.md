---
title: "Libbpf eBPF macro 'bpf_core_read_str'"
description: "This page documents the 'bpf_core_read_str' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_read_str`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `bpf_core_read_str` macro abstracts away [`bpf_probe_read_kernel_str`](../../../linux/helper-function/bpf_probe_read_kernel_str.md) call and captures offset relocation.

## Definition

```c
#define bpf_core_read_str(dst, sz, src)					    \
	bpf_probe_read_kernel_str(dst, sz, (const void *)__builtin_preserve_access_index(src))
```

## Usage

`bpf_core_read_str` is the string variant of the [`bpf_core_read`](bpf_core_read.md) macro. It is better suited for reading NUL-terminated strings from kernel memory.

See the [`bpf_core_read`](bpf_core_read.md) documentation for more details on usage of it and its variants like this macros.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
