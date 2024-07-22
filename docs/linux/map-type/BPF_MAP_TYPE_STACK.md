---
title: "Map Type 'BPF_MAP_TYPE_STACK'"
description: "This page documents the 'BPF_MAP_TYPE_STACK' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_STACK`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_STACK) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/f1a2e44a3aeccb3ff18d3ccc0b0203e70b95bd92)
<!-- [/FEATURE_TAG] -->

The stack map type is a generic map type, resembling a stack data structure.

## Usage

This map type has no keys, only values. The size and type of the values can be specified by the user to fit a large variety of use cases. The typical use-case for this map type is for brace matching (`{`,`}`) when parsing JSON for example.

As apposed to most map types, this map type uses a custom set of helpers to pop, peek and push elements, noted in the [helper functions](#helper-functions) section below.

## Attributes

While the [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) is essentially unrestricted, the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `0` since this map type has no keys.

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_LOOKUP_AND_DELETE_ELEM`](../syscall/BPF_MAP_LOOKUP_AND_DELETE_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)

!!! note
    The `BPF_MAP_LOOKUP_ELEM` syscall command acts as `peek`, `BPF_MAP_LOOKUP_AND_DELETE_ELEM` as `pop`, and `BPF_MAP_UPDATE_ELEM` as `push`.

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
 * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
 * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

### `BPF_F_NUMA_NODE`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/96eabe7a40aa17e613cf3db2c742ee8b1fc764d0)

When set, the [`numa_node`](../syscall/BPF_MAP_CREATE.md#numa_node) attribute is respected during map creation.

### `BPF_F_RDONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be read via the [syscall](../syscall/index.md) interface, but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly).

### `BPF_F_WRONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be written to via the [syscall](../syscall/index.md) interface, but not read from.

### `BPF_F_RDONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be read via [helper functions](../helper-function/index.md), but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly_prog).

### `BPF_F_WRONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be written to via [helper functions](../helper-function/index.md), but not read from.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_wronly_prog).
