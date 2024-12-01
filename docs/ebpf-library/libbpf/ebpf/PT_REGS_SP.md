---
title: "Libbpf eBPF macro 'PT_REGS_SP'"
description: "This page documents the 'PT_REGS_SP' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `PT_REGS_SP`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `PT_REGS_SP` macro make it easy to extract the stack pointer from `struct pt_regs` style contexts in an architecture-independent way.

## Usage

Since the `struct pt_regs` type represents the state of the CPU registers, it is different for every architecture. The `PT_REGS_SP` macro picks the correct register in the `struct pt_regs` type depending on the calling convention of the architecture.

The stack pointer tells you the address of the top of the stack, which can be used to read the value of variables on the stack (if you know the correct offsets).

The architecture for which the eBPF program is compiled is determined by setting one of the `__TARGET_ARCH_{arch}` macros. These are typically set by passing a flag to the compiler, such as `-D__TARGET_ARCH_x86` for x86. This allows for easy cross-compilation of eBPF programs for different architectures by changing the compiler invocation.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
