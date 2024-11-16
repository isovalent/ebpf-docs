---
title: "Libbpf eBPF macro 'offsetof'"
description: "This page documents the 'offsetof' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `offsetof`

[:octicons-tag-24: v0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)

The `offsetof` macro is used to calculate the offset of a member within a struct.

## Definition

`#!c #define offsetof(type, member)	((unsigned long)&((type *)0)->member)`

## Usage

This macro can be useful in cases where you need to calculate the offset of a field. This comes up when you need to provide a relative offset or need to do pointer arithmetic.

### Example

```c hl_lines="7"
#define ETH_HLEN 14

SEC("tc")
int example_prog(struct __sk_buff *ctx)
{
    __u32 assigned_ip = 0x0a000001; // 10.0.0.1
    bpf_skb_store_bytes(skb, ETH_HLEN + offsetof(struct iphdr, daddr), &assigned_ip, sizeof(__u32), 0) < 0)
}
```
