---
title: "Helper Function 'bpf_sock_map_update'"
description: "This page documents the 'bpf_sock_map_update' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sock_map_update`

<!-- [FEATURE_TAG](bpf_sock_map_update) -->
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/174a79ff9515f400b9a6115643dafd62a635b7e6)
<!-- [/FEATURE_TAG] -->

Add an entry to, or update a `map` referencing sockets.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_sock_map_update)(struct bpf_sock_ops *skops, void *map, void *key, __u64 flags) = (void *) 53;`

## Usage

The `skops` is used as a new value for the entry associated to `key`. `flags` is one of:

 * `BPF_NOEXIST` - The entry for `key` must not exist in the map.
 * `BPF_EXIST` - The entry for `key` must already exist in the map.
 * `BPF_ANY` - No condition on the existence of the entry for `key`.

If the `map` has eBPF programs (parser and verdict), those will be inherited by the socket being added. If the socket is already attached to eBPF programs, this results in an error.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [BPF_MAP_TYPE_SOCKMAP](../map-type/BPF_MAP_TYPE_SOCKMAP.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
