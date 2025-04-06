---
title: "Libbpf eBPF function 'bpf_usdt_arg_cnt'"
description: "This page documents the 'bpf_usdt_arg_cnt' libbpf eBPF function, including its definition, usage, and examples."
---
# Libbpf eBPF function `bpf_usdt_arg_cnt`

[:octicons-tag-24: v0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

The `bpf_usdt_arg_cnt` function is used to get the number of [USDT](../../../linux/concepts/usdt.md) arguments available.

## Definition

```c
[__weak](__weak.md) [__hidden](__hidden.md)
int bpf_usdt_arg_cnt(struct pt_regs *ctx)
```

Return number of USDT arguments defined for currently traced USDT.

## Usage

The `bpf_usdt_arg_cnt` function returns the number of arguments a tracepoint has.

This function uses specifiers that describe the arguments. These specifiers live in maps, defined in [`usdt.bpf.h`](https://github.com/libbpf/libbpf/blob/master/src/usdt.bpf.h) which should be populated by the loader, in particular the [`bpf_program__attach_usdt`](../userspace/bpf_program__attach_usdt.md) function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
