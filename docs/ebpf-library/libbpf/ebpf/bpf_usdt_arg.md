---
title: "Libbpf eBPF function 'bpf_usdt_arg'"
description: "This page documents the 'bpf_usdt_arg' libbpf eBPF function, including its definition, usage, and examples."
---
# Libbpf eBPF function `bpf_usdt_arg`

[:octicons-tag-24: v0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

The `bpf_usdt_arg` function is used to get a [USDT](../../../linux/concepts/usdt.md) argument from the process.

## Definition

```c
[__weak](__weak.md) [__hidden](__hidden.md)
int bpf_usdt_arg(struct pt_regs *ctx, __u64 arg_num, long *res)
```

Fetch USDT argument #`arg_num` (zero-indexed) and put its value into `*res`. Returns `0` on success; negative error, otherwise. On error `*res` is guaranteed to be set to zero.

## Usage

The `bpf_usdt_arg` function extracts a USDT argument from the process. Unlike arguments to a function which follow an ABI, the arguments to a tracepoint can live anywhere in the process. The tracepoint describes them using [GAS(GNU assembler) operands](https://en.wikibooks.org/wiki/X86_Assembly/GNU_assembly_syntax). This macro allows you to define your programs as these arguments were actually passed to the program like normal arguments.

This function uses specifiers that describe where to find a given argument. These specifiers live in maps, defined in [`usdt.bpf.h`](https://github.com/libbpf/libbpf/blob/master/src/usdt.bpf.h) which should be populated by the loader, in particular the [`bpf_program__attach_usdt`](../userspace/bpf_program__attach_usdt.md) function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
