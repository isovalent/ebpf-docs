---
title: "SCX eBPF macro 'likely'"
description: "This page documents the 'likely' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `likely`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `likely` macro hints to the compiler that an expression is likely to be true. This can help the compiler optimize the generated code.

## Definition

```c
#define likely(x) __builtin_expect(!!(x), 1)
```

## Usage

The `likely` macro can be used on boolean expressions, like in `if` statements, where you expect the expression to be true most of the time. This can help the compiler optimize the generated instructions.

An example of such an optimization would be to place instructions for the likely code path close to the conditional jump to improve instruction cache locality.

### Example

```c
if (likely(x > 0)) {
    // Likely code path
} else {
    // Unlikely code path
}
```
