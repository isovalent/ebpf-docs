---
title: Functions
description: This page explains the concept of eBPF functions, the different ways to use them, and how their usage has changed over time.
---
# Functions

When we talk about a function, we are referring to a function you would write in C or comparable programming language. In eBPF, the term function, program, and sub-program are often used interchangeably. A program refers to a function with a single argument, that being the context. A program can be attached to a hook point. Sub-programs, also referred to as BPF-to-BPF functions, are functions with zero to five arguments, they cannot be attached to a hook point, and are called from a program or special mechanism.

## Calling convention

The eBPF instruction set defines the calling convention for functions, these include programs, sub-programs, helper functions and kfuncs. Every function no matter who defines it uses the same calling convention. The R0 register is uses as the return value, a function should set it before returning unless it is a void function. Registers R1-R5 are used for arguments, R1 for the first argument, R2 for the second, and so on. Unlike calling conventions on native architectures, arguments are never passed via the stack. So 5 arguments is a hard limit, structures must be used to work around this. Registers R1-5 are clobbered after a function call, the verifier will not allow you to read from them until they are set with a known value. R6-9 are called saved registers, they are preserved across function calls.

## BPF to BPF functions (sub-programs)

In [:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/cc8b0b92a1699bc32f7fec71daa2bfc90de43a4d) BPF to BPF function calls were added. This allows BPF programs to reuse logic within the same program. A function can take up to 5 arguments, limited by the calling convention. It gets a fresh stack frame, which like any stack is free-ed and reused after the program exits. Functions can also access memory on the stack of a caller if a pointer to it is passed as an argument. Functions are limited to a maximum call depth of 8, so recursion of any meaningful depth is not possible.

### Function inlining

By default, the compiler will chose inline a function or to keep it a separate function. Compilers can be encouraged to inline or not inline a function with arguments like `__attribute__((always_inline))`/[`__always_inline`](../../ebpf-library/libbpf/ebpf/__always_inline.md) or `__attribute__((noinline))`/[`__noinline`](../../ebpf-library/libbpf/ebpf/__noinline.md). Inlined functions do not incur the overhead of a function call as they will become part of the calling function. Inlined functions can also be optimized per call site since arguments are known.

### Function by function verification

Until [:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/51c39bb1d5d105a02e29aa7960f0a395086e6342) the verifier would re-verify that a function was safe for every call site. Meaning that if you have a function which is called 10 times, then the verifier would check for every call that with the given inputs the function was save. This defeats the purpose of functions somewhat since you still incur verifier complexity for every call. 

Since [:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/51c39bb1d5d105a02e29aa7960f0a395086e6342) a distinction is made between "static" and "global" functions. Static functions are still verified as usual. But global functions undergo "function by function verification". This means that the verifier will verify every function once, and even out of order. It will assume all possible input values are possible, since it will not check every call site anymore. Therefor, functions might require more input checking to pass the verifier. This change reduces verification times and complexity. Static functions are functions marked with the `static` keyword in the C code, global functions regular non-static functions.

### Global function replacement

Also added in [:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/be8704ff07d2374bcc5c675526f95e70c6459683) is the ability to replace global functions. The primary use case for this is libxdp which used this to implement XDP program chaining from a dispatcher program.

See [Program Type `BPF_PROG_TYPE_EXT`](../program-type/BPF_PROG_TYPE_EXT.md) for more information.

### Mixing tail calls and functions

Since [:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/e411901c0b775a3ae7f3e2505f8d2d90ac696178) it is possible to mix tail calls and functions. Before then you had to pick one or the other.

Mixing tail calls and functions causes the available stack size per function to shrink from `512` bytes to `256` bytes. The reasoning being that when a tail call is made, the current stack frame is reused, but if that tail call is made from a function, the stack of the caller can not be reused. By default, kernel threads are limited to 8k stack sizes. By decreasing the max stack size, it is harder to run out of stack space, though it is still possible.

Mixing tail calls and functions requires support from the JIT for a given architecture. This is because the tail call counter (which prevents doing more than `32` tail calls) has to be propagated to functions so they can pass it to the next tail call program. The practical result of this is that support for mixing tail calls and functions got added in stages. Here is a table of when support was added for each architecture:

| Architecture | Support added |
|--------------|---------------|
| x86          | [:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/e411901c0b775a3ae7f3e2505f8d2d90ac696178) |
| ARM64        | [:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/d4609a5d8c70d21b4a3f801cf896a3c16c613fe1) |
| s390         | [:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/dd691e847d28ac5f8b8e3005be44fd0e46722809) |
| LoongArch    | [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/bb035ef0cc91e115faa80187ac8886a7f1914d06) |

Architectures that are not listed do not support mixing tail calls and functions as of :octicons-tag-24: v6.15.

### Callbacks

In [:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/69c087ba6225b574afb6e505b72cb75242a3d844) the verifier was extended to allow for callbacks. Since then a number of helper function and kfuncs have been added that call back into given functions.

### Argument annotations

In [:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/94e1c70a34523b5e1529e4ec508316acc6a26a2b) global function argument annotation were added. These are a set of annotations (in practice these are BTF decl tags), which if added to an attribute, tell the verifier to restrict the input values to the function. Possible tags are:

* [`__arg_ctx`](../../ebpf-library/libbpf/ebpf/__arg_ctx.md) - The argument is a pointer to a program context.
* [`__arg_nonnull`](../../ebpf-library/libbpf/ebpf/__arg_nonnull.md) - The argument can not be NULL.
* [`__arg_nullable`](../../ebpf-library/libbpf/ebpf/__arg_nullable.md) - The argument can be NULL.
* [`__arg_trusted`](../../ebpf-library/libbpf/ebpf/__arg_trusted.md) - The argument must be a [trusted value](kfuncs.md#kf_trusted_args).
* [`__arg_arena`](../../ebpf-library/libbpf/ebpf/__arg_arena.md) - The argument must be a pointer to a [memory arena](../map-type/BPF_MAP_TYPE_ARENA.md).
