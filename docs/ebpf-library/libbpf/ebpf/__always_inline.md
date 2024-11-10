---
title: "Libbpf eBPF macro '__always_inline'"
description: "This page documents the '__always_inline' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__always_inline`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `__always_inline` macros is used to tell the compiler to inline a function.

## Definition

`#!c #define __always_inline inline __attribute__((always_inline))`

## Usage

This macro is used to tell the compiler to inline a function. This is a best attempt at "forcing" the compiler to inline a function. The combination of the C `inline` keyword and the `__attribute__((always_inline))` attribute gives the strongest hint possible to the compiler that the function should be inlined.

This is particularly useful when writing eBPF programs, for kernel version that do not support BPF-to-BPF functions or when they are unwanted for performance reasons.

It should be noted that the Clang docs states that the `always_inline` attribute is not a guarantee that the function will be inlined.

### Example

```c hl_lines="1"
static int __always_inline add(int a, int b)
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
