---
title: "Libbpf eBPF macro '__bpf_unreachable'"
description: "This page documents the '__bpf_unreachable' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__bpf_unreachable`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `__bpf_unreachable` macro is used to cause a compile time error if the code path is reached.

## Definition

`#!c #define __bpf_unreachable() __builtin_trap()`

## Usage

This macro asks the compiler to generate a trap instruction, which in a userspace program would crash the program on runtime. However, the eBPF compiler backend does not implement this, so it actually will cause the compiler to emit an error if the code path is reached.

This is useful if you want to assert at compile time that a certain code path is never reachable to prevent accidental bugs.

### Example

In this example, we use the default clause of the switch statement to assert that `my_func` is always called with one of the enum values handled. If `my_enum` is ever extended without updating the switch statement, and the new value is passed to `my_func`, the compiler will emit an error. This will also happen if user input is used to call `my_func` and the input is not properly validated.

```c hl_lines="15"
enum my_enum {
    MY_ENUM_A,
    MY_ENUM_B,
};

static void __always_inline my_func(enum my_enum e) {
    switch (e) {
    case MY_ENUM_A:
        // Do something
        break;
    case MY_ENUM_B:
        // Do something
        break;
    default:
        __bpf_unreachable();
    }
}
```
