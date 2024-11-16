---
title: "Libbpf eBPF macro 'barrier'"
description: "This page documents the 'barrier' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `barrier`

[:octicons-tag-24: v0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

The `barrier` macro is used to prevent the compiler from reordering memory operations.

## Definition

`#!c #define barrier() asm volatile("" ::: "memory")`

## Usage

This macro inserts what is referred to as a "full memory barrier". Compilers such as GCC and Clang will sometimes reorder memory operations to optimize code, so your actual program may not execute in the order you wrote it. If the order of memory operations is important, you can use the `barrier` macro to tell the compiler to not reorder memory operations across the barrier.

This is a very specialized tool which you will likely not use often.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
