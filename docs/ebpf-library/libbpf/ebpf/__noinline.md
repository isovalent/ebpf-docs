---
title: "Libbpf eBPF macro '__noinline'"
description: "This page documents the '__noinline' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__noinline`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `__noinline` macros is used to tell the compiler to inline a function.

## Definition

`#!c #define __noinline __attribute__((noinline))`

## Usage

This macro is used to tell the compiler to not inline a function. This is mainly useful to force the use of a BPF-to-BPF function call. This might be important if you are at the limit of stack space in calling functions.

### Example

```c hl_lines="1"
static int __noinline add(int a, int b)
{
    return a + b;
}

SEC("xdp")
int example_prog(struct xdp_md *ctx)
{
    if (add(1, 2) == 3)
        return XDP_PASS;
    else
        return XDP_DROP;
}
```
