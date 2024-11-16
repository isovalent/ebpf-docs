---
title: "Libbpf eBPF macro 'barrier_var'"
description: "This page documents the 'barrier_var' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `barrier_var`

[:octicons-tag-24: v0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

The `barrier_var` macro is used to prevent the compiler from reordering memory operations for a specific variable.

## Definition

`#!c #define barrier_var(var) asm volatile("" : "+r"(var))`

## Usage

This macro is a variable-specific compiler (optimization) barrier. It's a no-op which makes compiler believe that there is some black box modification of a given variable and thus prevents compiler from making extra assumption about its value and potential simplifications and optimizations on this variable.

E.g., compiler might often delay or even omit 32-bit to 64-bit casting of a variable, making some code patterns unverifiable. Putting barrier_var() in place will ensure that cast is performed before the barrier_var() invocation, because compiler has to pessimistically assume that embedded <nospell>asm</nospell> section might perform some extra operations on that variable.

This is a variable-specific variant of more global [`barrier()`](barrier.md).

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
