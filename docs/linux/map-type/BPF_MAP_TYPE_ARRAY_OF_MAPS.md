---
title: "Map Type 'BPF_MAP_TYPE_ARRAY_OF_MAPS'"
description: "This page documents the 'BPF_MAP_TYPE_ARRAY_OF_MAPS' eBPF map type, including its defintion, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_ARRAY_OF_MAPS`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_ARRAY_OF_MAPS) -->
[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/56f668dfe00dcf086734f1c42ea999398fad6572)
<!-- [/FEATURE_TAG] -->

The array of maps map type contains references to other maps.

## Usage

This map type is a map-in-map type. The map values contain references to other BPF maps. We will refer to map-in-map as the "outer map" and the maps referenced as the "inner map(s)". The key advantage of using a map-in-map is that the outer map is directly referenced by any programs that use it, but the inner maps are not.

There are a couple of use cases for this indirection. First is to implement a form of RCU on the whole map, copying the existing map to a new map, updating multiple fields, then switching out the maps in the outer map. A second use case can be for statistics/metrics accuracy. It takes time for userspace to iterate and read the full contents of a map containing statistics/metrics. If accuracy is required, the outer map can contain multiple inner maps with counters, the program switches the inner map it writes to, giving userspace time to collect the values on the other maps, resulting in very accurate measurements. Lastly, since [:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/134fede4eecfcbe7900e789f625fa6f9c3a8cd0e) most inner map types can have varying [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) values from the reference map. This allows for dynamic resizing of a map without having to reload any programs.

!!! warning
    For inner maps of type `BPF_MAP_TYPE_ARRAY` the [`BPF_F_INNER_MAP`](../syscall/BPF_MAP_CREATE.md#bpf_f_inner_map) flag must be set on the inner map to allow varying [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries). And inner maps of type `BPF_MAP_TYPE_XSKMAP` must always have the same amount of [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) as the reference map.
    
!!! warning
    maps of type `BPF_MAP_TYPE_PERF_EVENT_ARRAY` are not allowed as inner maps.

Users should be aware of the read/write asymmetry of this map type:

* The [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md) syscall command takes *file descriptor* of the BPF map you wish to insert into the map.
* The [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md) syscall command returns the *ID* of the BPF map, which can be turned into a file descriptor with the [`BPF_MAP_GET_FD_BY_ID`](../syscall/BPF_MAP_GET_FD_BY_ID.md) syscall command.
* The [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md) returns a pointer to the inner map or `NULL`. This pointer can be used like any other in helpers that that map pointers.

## Attributes

Both the [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) and the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `4` indicating a 32-bit unsigned integer.

The [`inner_map_fd`](../syscall/BPF_MAP_CREATE.md#inner_map_fd) attribute must be set to the file descriptor of another map. This other map will serve as a template for the inner maps. After loading, during insertion of values, the kernel will verify that the spec of the inner map values you are attempting to insert match the spec of the map provided by this field. The map used to indicate the type is not linked to the map-in-map type in any way, it is just used to transfer type info. A common technique for loaders is to build a temporary map just for the purpose of providing the type info and freeing that map as soon as the outer map has been created.

!!! note
    If inner maps use the [`BPF_F_INNER_MAP`](../syscall/BPF_MAP_CREATE.md#bpf_f_inner_map) flag, the [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) field of the spec is ignored for the purposes of comparing the map spec.

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.

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
