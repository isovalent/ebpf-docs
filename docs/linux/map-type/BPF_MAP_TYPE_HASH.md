---
title: "Map Type 'BPF_MAP_TYPE_HASH'"
description: "This page documents the 'BPF_MAP_TYPE_HASH' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_HASH`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_HASH) -->
[:octicons-tag-24: v3.19](https://github.com/torvalds/linux/commit/0f8e4bd8a1fc8c4185f1630061d0a1f2d197a475)
<!-- [/FEATURE_TAG] -->

The hash map type is a generic map type with no restrictions on the structure of the key and value. Hash-maps are implemented using a hash table, allowing for lookups with arbitrary keys.

## Attributes

While the size of the key and value are essentially unrestricted both [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) and [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must be at least zero and their combined size no larger than `KMALLOC_MAX_SIZE`. `KMALLOC_MAX_SIZE` is the maximum size which can be allocated by the kernel memory allocator, its exact value being dependant on a number of factors. If this edge case is hit a `-E2BIG` [error number](https://man7.org/linux/man-pages/man3/errno.3.html) is returned to the [map create syscall](../syscall/BPF_MAP_CREATE.md).

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_LOOKUP_AND_DELETE_ELEM`](../syscall/BPF_MAP_LOOKUP_AND_DELETE_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)
* [`BPF_MAP_LOOKUP_BATCH`](../syscall/BPF_MAP_LOOKUP_BATCH.md)
* [`BPF_MAP_LOOKUP_AND_DELETE_BATCH`](../syscall/BPF_MAP_LOOKUP_AND_DELETE_BATCH.md)

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
 * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
 * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
 * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.

### `BPF_F_NO_PREALLOC`
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/6c90598174322b8888029e40dd84a4eb01f56afe)

Hash maps are pre-allocated by default, this means that even a completely empty hash map will use the same amount of
kernel memory as a full map. 

If this flag is set, pre-allocation is disabled. Users might consider this for large maps since allocating large amounts of memory takes a lot of time during creation and might be undesirable.

!!! warning
    The patch set[^1] does note that not pre-allocating may cause issues in some edge-cases, which was the original reason for defaulting to pre-allocation.

[^1]: [https://lwn.net/Articles/679074/](https://lwn.net/Articles/679074/)

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

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_wronly).

### `BPF_F_ZERO_SEED`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/96b3b6c9091d23289721350e32c63cc8749686be)

Setting this flag will initialize the hash table with a seed of 0.

The hashing algorithm used by the hash table is seeded with a random number by default. This seeding is meant as a mitigation against Denial of Service attacks which could exploit the predictability of hashing implementations.

This random seed makes hash map operations inherently random at access time. This flag was introduced to make performance evaluation more consistent.

!!! warning
    It is not recommended to use this flag in production due to the vulnerability to Denial of Service attacks.

!!! info
    Only users with the `CAP_SYS_ADMIN` capability can use this flag, `CAP_BPF` is not enough due to the security risk associated with the flag.

### `BPF_F_RDONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be read via [helper functions](../helper-function/index.md), but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly_prog).

<!-- TODO:  -->

### `BPF_F_WRONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be written to via [helper functions](../helper-function/index.md), but not read from.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_wronly_prog).

<!-- TODO usage examples -->

<!-- ## Internals -->
<!-- TODO locking / implementations -->
