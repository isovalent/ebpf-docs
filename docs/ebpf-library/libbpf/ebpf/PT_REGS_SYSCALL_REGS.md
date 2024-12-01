---
title: "Libbpf eBPF macro 'PT_REGS_SYSCALL_REGS'"
description: "This page documents the 'PT_REGS_SYSCALL_REGS' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `PT_REGS_SYSCALL_REGS`

[:octicons-tag-24: v0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)

The `PT_REGS_SYSCALL_REGS` macro that ensures consistent access to `struct pt_regs` when using a kprobe on a syscall.


## Definition

```c
#define PT_REGS_SYSCALL_REGS(ctx) ((struct pt_regs *)PT_REGS_PARM1(ctx))
```

### Usage

This macro is useful when placing kprobes on syscalls. On some CPU architectures the kernel will use a so called syscall wrapper (indicated by the `CONFIG_ARCH_HAS_SYSCALL_WRAPPER` kernel config). These wrappers change the syscall calling convention, so the actual `struct pt_regs` are in a different location then normally expected.

This macro corrects for this based on the target architecture and returns the expected `struct pt_regs`.

The architecture for which the eBPF program is compiled is determined by setting one of the `__TARGET_ARCH_{arch}` macros. These are typically set by passing a flag to the compiler, such as `-D__TARGET_ARCH_x86` for x86. This allows for easy cross-compilation of eBPF programs for different architectures by changing the compiler invocation.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
