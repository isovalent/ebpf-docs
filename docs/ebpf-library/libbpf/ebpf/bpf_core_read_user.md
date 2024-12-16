---
title: "Libbpf eBPF macro 'bpf_core_read_user'"
description: "This page documents the 'bpf_core_read' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_read_user`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

The `bpf_core_read_user` macro abstracts away [`bpf_probe_read_user`](../../../linux/helper-function/bpf_probe_read_user.md) call and captures offset relocation.

## Definition

```c
#define bpf_core_read_user(dst, sz, src)				    \
	bpf_probe_read_user(dst, sz, (const void *)__builtin_preserve_access_index(src))
```

## Usage

The `bpf_core_read_user` is the userspace variant of [`bpf_core_read`](bpf_core_read.md). It wraps the [`bpf_probe_read_user`](../../../linux/helper-function/bpf_probe_read_user.md) helper function instead of [`bpf_probe_read`](../../../linux/helper-function/bpf_probe_read.md).

Please refer to the [`bpf_core_read`](bpf_core_read.md) documentation for more details on usage of it and its variants like this macros.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
