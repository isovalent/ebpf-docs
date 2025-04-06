---
title: "Libbpf eBPF function 'bpf_usdt_cookie'"
description: "This page documents the 'bpf_usdt_cookie' libbpf eBPF function, including its definition, usage, and examples."
---
# Libbpf eBPF function `bpf_usdt_cookie`

[:octicons-tag-24: v0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

The `bpf_usdt_cookie` function returns the cookie associated with the attach point.

## Definition

```c
[__weak](__weak.md) [__hidden](__hidden.md)
long bpf_usdt_cookie(struct pt_regs *ctx)
```

Retrieve user-specified cookie value provided during attach as [`bpf_usdt_opts.usdt_cookie`](../userspace/bpf_program__attach_usdt.md#usdt_cookie). This serves the same purpose as BPF cookie returned by [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). Libbpf's support for USDT is itself
utilizing BPF cookies internally, so user can't use BPF cookie directly  for USDT programs and has to use `bpf_usdt_cookie` API instead.

## Usage

This function uses specifiers that describe the tracepoint. These specifiers live in maps, defined in [`usdt.bpf.h`](https://github.com/libbpf/libbpf/blob/master/src/usdt.bpf.h) which should be populated by the loader, in particular the [`bpf_program__attach_usdt`](../userspace/bpf_program__attach_usdt.md) function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
