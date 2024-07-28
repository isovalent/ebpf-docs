---
title: "Map Type 'BPF_MAP_TYPE_SOCKMAP'"
description: "This page documents the 'BPF_MAP_TYPE_SOCKMAP' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_SOCKMAP`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_SOCKMAP) -->
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/174a79ff9515f400b9a6115643dafd62a635b7e6)
<!-- [/FEATURE_TAG] -->

The socket map is a specialized map type which hold network sockets as value.

## Usage

This map type can be use too lookup a pointer to a socket with the [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md) helper function, which then can be passed to helpers such as [`bpf_sk_assign`](../helper-function/bpf_sk_assign.md) or the a map reference can be used directly in a range of helpers such as [`bpf_sk_redirect_map`](../helper-function/bpf_sk_redirect_map.md), [`bpf_msg_redirect_map`](../helper-function/bpf_msg_redirect_map.md) and [`bpf_sk_select_reuseport`](../helper-function/bpf_sk_select_reuseport.md). All of the above cases redirect a packet or connection in some way, the details differ depending on the program type and the helper function, so please visit the specific pages for details.

!!! note
    Sockets returned by `bpf_map_lookup_elem` are ref-counted, so the caller must call [`bpf_sk_release`](../helper-function/bpf_sk_release.md) in all code paths where the returned socket is not NULL before exiting the program. This is enforced by the verifier which will throw a `Unreleased reference` error if socket pointers are not released.

This map can also be manipulated from kernel space, the main use-case for doing so seems to be to manage the contents of the map automatically from program types that trigger on socket events. This would allow 1 program to manage the contents of the map, and another to do the actual redirecting on packet events.

[`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md) and [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md) programs can be attached to this map type. When a socket is inserted into the map, its socket callbacks will be replaced with these programs.

The attach types for the map programs are:

* `msg_parser` program - [`BPF_SK_MSG_VERDICT`](../syscall/BPF_LINK_CREATE.md#bpf_sk_msg_verdict).
* `stream_parser` program - [`BPF_SK_SKB_STREAM_PARSER`](../syscall/BPF_LINK_CREATE.md#bpf_sk_skb_stream_parser).
* `stream_verdict` program - [`BPF_SK_SKB_STREAM_VERDICT`](../syscall/BPF_LINK_CREATE.md#bpf_sk_skb_stream_verdict).
* `skb_verdict` program - [`BPF_SK_SKB_VERDICT`](../syscall/BPF_LINK_CREATE.md#bpf_sk_skb_verdict).


A sock object may be in multiple maps, but can only inherit a single parse or verdict program. If adding a sock object to a map would result in having multiple parser programs the update will return an `EBUSY` error.

!!! warning
    Users are not allowed to attach stream_verdict and skb_verdict programs to the same map.

## Attributes

The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) must always be `4` and the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `8`. 

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [`bpf_sk_redirect_map`](../helper-function/bpf_sk_redirect_map.md)
 * [`bpf_sock_map_update`](../helper-function/bpf_sock_map_update.md)
 * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
 * [`bpf_msg_redirect_map`](../helper-function/bpf_msg_redirect_map.md)
 * [`bpf_sk_select_reuseport`](../helper-function/bpf_sk_select_reuseport.md)
 * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
 * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
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
