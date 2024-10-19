---
title: "Libbpf eBPF macro '__type'"
description: "This page documents the '__type' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__type`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `__type` macros is used to define type properties of BTF maps.

## Definition

`#!c #define __type(name, val) typeof(val) *name`

## Usage

This macro is used to encode type properties in BTF map definitions. For example, to define the key or value type of a BTF map.

### Example

```c hl_lines="4 5"
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __type(key, __u32);
    __type(value, struct example_struct);
} SEC(".maps") my_map;
```
