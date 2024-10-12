---
title: "Helper Function 'bpf_ntohl'"
description: "This page documents the 'bpf_ntohl' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_ntohl`

## Definition

<!-- [HELPER_FUNC_DEF] -->

!!! info 
    LLVM's BPF target selects the endianness of the CPU it compiles on, or the user specifies (bpfel/bpfeb), respectively. The used __BYTE_ORDER__ is defined by the compiler, we cannot rely on __BYTE_ORDER from libc headers, since it doesn't reflect the actual requested byte order.

Used to converts the unsigned integer from network byte order to host byte order.

### Returns

A 32-bit unsigned integer converted from network byte order to host byte order.
<!-- [/HELPER_FUNC_DEF] -->

## Usage

```c
uint32_t bpf_ntohl(uint32_t);
```

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
