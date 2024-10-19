---
title: "Libbpf eBPF macro '__uint'"
description: "This page documents the '__uint' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__uint`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `__uint` macros is used to define unsigned integer properties of BTF maps.

## Definition

`#!c #define __uint(name, val) int (*name)[val]`

## Usage

This macro is used to encode unsigned integer properties in BTF map definitions. BTF does not have a notion of literal values, we encode them as pointers to arrays of size `X`, were `X` is the actual number we want to communicate.

### Example

```c
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __uint(key_size, sizeof(int));
    __uint(value_size, sizeof(long));
} SEC(".maps") my_map;
```
