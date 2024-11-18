---
title: "Libbpf eBPF macro 'bpf_ksym_exists'"
description: "This page documents the 'bpf_ksym_exists' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_ksym_exists`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `bpf_ksym_exists` macro is used to check if a [`__weak`](__weak.md) symbol has been defined.

## Definition

```c
#define bpf_ksym_exists(sym) ({ \
    _Static_assert(!__builtin_constant_p(!!sym), #sym " should be marked as __weak"); \
    !!sym; \
})
```

## Usage

This macro is used to check if a weak symbol has been defined. If the symbol is defined, it is better practice to use this macro to check instead of doing a `== 0` check. This is because this macro includes a static assertion to see if the symbol is marked as weak, this can avoid unintended bugs.

### Example
```c hl_lines="6"
extern int bpf_dynptr_from_xdp(struct xdp_md *x, u64 flags, struct bpf_dynptr *ptr__uninit) __weak __ksym;

SEC("xdp.frags")
int example_prog(struct xdp_md *ctx)
{
    if (bpf_ksym_exists(bpf_dynptr_from_xdp)) {
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
