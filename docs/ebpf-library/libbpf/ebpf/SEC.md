---
title: "Libbpf eBPF macro 'SEC'"
description: "This page documents the 'SEC' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `SEC`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `SEC` macros is used to tell the compiler in which ELF section to place a symbol.

## Definition

`#!c #define SEC(NAME) __attribute__((section(NAME), used))`

## Usage

This macro is used to place a symbol in a specific ELF section. You will typically see this on eBPF programs and map definitions, however, they can be used on any symbol. The `NAME` argument is the name of the ELF section. The section in which a symbol is placed often has implicit meaning and can change how a loader (library) such as libbpf will interpret the contents of that section.

For example the `maps` section is used for legacy map definitions, and BTF based maps have to be placed in the `.maps` section. All programs in the `xdp` section will we loaded as programs of type `BPF_PROG_TYPE_XDP` and attach type `BPF_XDP`, in the `xdp.frags` will also have the `BPF_F_XDP_HAS_FRAGS` flag set for example.

A table of well-known program sections can be found [here](https://docs.kernel.org/bpf/libbpf/program_types.html). Note that these do not include non-program sections such as `.maps`, `license`, `ksym`, etc.

### Example


```c hl_lines="1"
SEC("xdp")
int example_prog(struct xdp_md *ctx)
{
    void *data_end = (void *)(long)ctx->data_end;
    void *data = (void *)(long)ctx->data;

    if (data + sizeof(struct ethhdr) > data_end)
        return XDP_DROP;

    struct ethhdr *eth = data;
    if (eth->h_proto == htons(ETH_P_IP))
        return XDP_PASS;
}
```
