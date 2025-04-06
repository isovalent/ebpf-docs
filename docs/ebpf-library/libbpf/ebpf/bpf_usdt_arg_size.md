---
title: "Libbpf eBPF function 'bpf_usdt_arg_size'"
description: "This page documents the 'bpf_usdt_arg_size' libbpf eBPF function, including its definition, usage, and examples."
---
# Libbpf eBPF function `bpf_usdt_arg_size`

[:octicons-tag-24: v1.6.0](https://github.com/libbpf/libbpf/releases/tag/v1.6.0)

The `bpf_usdt_arg_size` function is used to get the size of a [USDT](../../../linux/concepts/usdt.md) argument.

## Definition

```c
static [__always_inline](__always_inline.md)
int bpf_usdt_arg_size(struct pt_regs *ctx, __u64 arg_num)
```

Returns the size in bytes of the #`arg_num` (zero-indexed) USDT argument. Returns negative error if argument is not found or `arg_num` is invalid.

## Usage

The `bpf_usdt_arg` function returns the size of a given argument.

This function uses specifiers that describe the size of the argument. These specifiers live in maps, defined in [`usdt.bpf.h`](https://github.com/libbpf/libbpf/blob/master/src/usdt.bpf.h) which should be populated by the loader, in particular the [`bpf_program__attach_usdt`](../userspace/bpf_program__attach_usdt.md) function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
