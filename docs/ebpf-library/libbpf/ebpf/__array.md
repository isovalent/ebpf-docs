---
title: "Libbpf eBPF macro '__array'"
description: "This page documents the '__array' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__array`

[:octicons-tag-24: v0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)

The `__array` macros is used to define array properties of BTF maps.

## Definition

`#!c #define __array(name, val) typeof(val) *name[]`

## Usage

This macro is used to encode array properties in BTF map definitions. This is useful when defining maps of maps or tail call maps in eBPF programs.

### Example

```c hl_lines="13 15 16 17 18"
struct inner_map {
        __uint(type, BPF_MAP_TYPE_ARRAY);
        __uint(max_entries, 1);
        __type(key, int);
        __type(value, int);
} inner_map1 SEC(".maps"),
inner_map2 SEC(".maps");

struct outer_hash {
        __uint(type, BPF_MAP_TYPE_HASH_OF_MAPS);
        __uint(max_entries, 5);
        __uint(key_size, sizeof(int));
        __array(values, struct inner_map);
} outer_hash SEC(".maps") = {
        .values = {
                [0] = &inner_map2,
                [4] = &inner_map1,
        },
};
```
