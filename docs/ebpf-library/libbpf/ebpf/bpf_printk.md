---
title: "Libbpf eBPF macro 'bpf_printk'"
description: "This page documents the 'bpf_printk' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_printk`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `bpf_printk` macro is used to make printing to the kernel trace log easier.

## Definition

The short version (left out the implementation)
```c
/* Helper macro to print out debug messages */
#define bpf_printk(fmt, args...) ___bpf_pick_printk(args)(fmt, ##args)
```

??? example "The full version"
    ```c
    #ifdef BPF_NO_GLOBAL_DATA
    #define BPF_PRINTK_FMT_MOD
    #else
    #define BPF_PRINTK_FMT_MOD static const
    #endif

    #define __bpf_printk(fmt, ...)				\
    ({							\
        BPF_PRINTK_FMT_MOD char ____fmt[] = fmt;	\
        bpf_trace_printk(____fmt, sizeof(____fmt),	\
                ##__VA_ARGS__);		\
    })

    /*
    * __bpf_vprintk wraps the bpf_trace_vprintk helper with variadic arguments
    * instead of an array of u64.
    */
    #define __bpf_vprintk(fmt, args...)				\
    ({								\
        static const char ___fmt[] = fmt;			\
        unsigned long long ___param[___bpf_narg(args)];		\
                                    \
        _Pragma("GCC diagnostic push")				\
        _Pragma("GCC diagnostic ignored \"-Wint-conversion\"")	\
        ___bpf_fill(___param, args);				\
        _Pragma("GCC diagnostic pop")				\
                                    \
        bpf_trace_vprintk(___fmt, sizeof(___fmt),		\
                ___param, sizeof(___param));		\
    })

    /* Use __bpf_printk when bpf_printk call has 3 or fewer fmt args
    * Otherwise use __bpf_vprintk
    */
    #define ___bpf_pick_printk(...) \
        ___bpf_nth(_, ##__VA_ARGS__, __bpf_vprintk, __bpf_vprintk, __bpf_vprintk,	\
            __bpf_vprintk, __bpf_vprintk, __bpf_vprintk, __bpf_vprintk,		\
            __bpf_vprintk, __bpf_vprintk, __bpf_printk /*3*/, __bpf_printk /*2*/,\
            __bpf_printk /*1*/, __bpf_printk /*0*/)

    /* Helper macro to print out debug messages */
    #define bpf_printk(fmt, args...) ___bpf_pick_printk(args)(fmt, ##args)
    ```

## Usage

This is a macro makes writing to the kernel trace log easier. It does two things, the first is it picks between the [`bpf_trace_printk`](../../../linux/helper-function/bpf_trace_printk.md) and [`bpf_trace_vprintk`](../../../linux/helper-function/bpf_trace_vprintk.md) helper functions. The [`bpf_trace_printk`](../../../linux/helper-function/bpf_trace_printk.md) helper only supports up to 3 arguments besides the format string, but [`bpf_trace_vprintk`](../../../linux/helper-function/bpf_trace_vprintk.md) supports any number of arguments (the macro supports up to 12 arguments).

!!! note
    While the [`bpf_trace_printk`](../../../linux/helper-function/bpf_trace_printk.md) helper is supported on most kernels, the [`bpf_trace_vprintk`](../../../linux/helper-function/bpf_trace_vprintk.md) helper is only supported on kernel versions [:octicons-tag-24: v5.16](https://github.com/torvalds/linux/commit/10aceb629e198429c849d5e995c3bb1ba7a9aaa3) and later.

    So using this macro with more than 3 arguments on older kernels might cause the verifier to reject your program.

The second thing it does is it turns the literal format string into a `char` array. If `BPF_NO_GLOBAL_DATA` is defined the char array will live on the stack, otherwise it will be turned into a global variable (this only works up to 4 arguments, the [`bpf_trace_vprintk`](../../../linux/helper-function/bpf_trace_vprintk.md) variant always make the format string into a global variable). Without this `char` array the compiler will place strings in a dedicated string ELF section unrecognized by loaders and not emit the proper global variable relocations.

### Example
```c
SEC("tc")
int example_prog(struct __sk_buff *ctx)
{
    // Will use bpf_trace_printk
    bpf_printk(
        "Got a packet from interface %d, src: %pi4, dst: %pi4\\n", 
        ctx->ingress_ifindex, 
        ctx->remote_ip4, 
        ctx->local_ip4,
    );

    // Will use bpf_trace_vprintk
    bpf_printk(
        "Got a packet from interface %d, src: %pi4, dst: %pi4, src port: %d, dst port: %d\\n", 
        ctx->ingress_ifindex, 
        ctx->remote_ip4, 
        ctx->local_ip4,
        ctx->remote_port,
        ctx->local_port,
    );

    return TC_ACT_OK;
}
```
