---
title: "SCX eBPF macro 'unlikely'"
description: "This page documents the 'unlikely' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `unlikely`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `unlikely` macro hints to the compiler that an expression is unlikely to be true. This can help the compiler optimize the generated code.

## Definition

```c
#define unlikely(x) __builtin_expect(!!(x), 0)
```

## Usage

The `unlikely` macro can be used on boolean expressions, like in `if` statements, where you expect the expression to be false most of the time. This can help the compiler optimize the generated instructions.

An example of such an optimization would be to place instructions for the likely code path close to the conditional jump to improve instruction cache locality.

### Example

```c
if (unlikely(ptr == NULL)) {
    // Unlikely code path
    return -EINVAL;
}
// Likely code path
```
