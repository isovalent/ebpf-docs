---
title: "Libbpf eBPF macro '__weak'"
description: "This page documents the '__weak' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__weak`

[:octicons-tag-24: v0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)

The `__weak` macros is used to mark a symbol is weak.

## Definition

`#!c #define __weak __attribute__((weak))`

## Usage

This macro is used mark a symbol as being "weak" as apposed to normal symbols which are "strong". Originally, this property was used to inform the linker when linking together multiple object files. A weak symbol can be overridden by a strong symbol from another object file. So in a sense it is a way to provide a default implementation that can be overridden depending on which files are linked together.

In the context of eBPF we do not typically link multiple object files together. Here it is used to tell the loader (library) to do a best effort to provide some sort of information. For example a kernel symbol or kfunc. Normally, a loader will throw an error if it cannot find a symbol, but by marking a symbol as weak, the user can tell the loader to tolerate the symbol not being found and continue anyway. The eBPF program is expected to handle the case where the symbol is not found.

Another side effect of marking a symbol as weak is that the compiler cannot make assumptions about the contents of the symbol. This makes inlining or other optimizations impossible. And since a linker is supposed to know about the existence of weak symbols, they are always emitted in the object file, even if they are not used. Sometimes developers might mark a symbol as weak for these reasons.

### Example

In this example we mark the `bpf_dynptr_from_xdp` function as weak. If the function is not found, its value is `NULL`. We can check with an if statement if the function is available and use it if it is.
Otherwise we fall back to an older way of accessing the data.

```c hl_lines="1"
extern int bpf_dynptr_from_xdp(struct xdp_md *x, u64 flags, struct bpf_dynptr *ptr__uninit) __weak __ksym;

SEC("xdp.frags")
int example_prog(struct xdp_md *ctx)
{
    if (bpf_dynptr_from_xdp) {
        struct bpf_dynptr ptr;
        if (bpf_dynptr_from_xdp(ctx, 0, &ptr) < 0)
            return XDP_DROP;

        __u8 buf[sizeof(struct ethhdr)];
        struct ethhdr *eth = bpf_dynptr_slice(&ptr, buf, sizeof(buf));
        if (!eth)
            return XDP_DROP;

        if (eth->h_proto == htons(ETH_P_IP))
            return XDP_PASS;
    } else {
        void *data_end = (void *)(long)ctx->data_end;
        void *data = (void *)(long)ctx->data;

        if (data + sizeof(struct ethhdr) > data_end)
            return XDP_DROP;

        struct ethhdr *eth = data;
        if (eth->h_proto == htons(ETH_P_IP))
            return XDP_PASS;
    }
}
```
