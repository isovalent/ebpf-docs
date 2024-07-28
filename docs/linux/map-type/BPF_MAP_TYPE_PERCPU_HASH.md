---
title: "Map Type 'BPF_MAP_TYPE_PERCPU_HASH'"
description: "This page documents the 'BPF_MAP_TYPE_PERCPU_HASH' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_PERCPU_HASH`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_PERCPU_HASH) -->
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/824bd0ce6c7c43a9e1e210abf124958e54d88342)
<!-- [/FEATURE_TAG] -->

This is the per-CPU variant of the [`BPF_MAP_TYPE_HASH`](BPF_MAP_TYPE_HASH.md) map type. 

This map type is a generic map type with no restrictions on the structure of the key and value. Hash-maps are implemented using a hash table, allowing for lookups with arbitrary keys. 

This per-CPU version has a separate hash map for each logical CPU. When accessing the map using most [helper function](../helper-function/index.md), the hash map assigned to the CPU the eBPF program is currently running on is accessed implicitly. 

Since preemption is disabled during program execution, no other programs will be able to concurrently access the same memory. This guarantees there will never be any race conditions and improves the performance due to the lack of congestion and synchronization logic, at the cost of having a large memory footprint.

<!-- TODO: On newer kernels CPU migration is disabled, not preemption, check the implications of that against the above statements -->
<!-- TODO: "preemption" need a link -->

!!! note
    The `bpf_map_lookup_percpu_elem` helper can be used to access maps assigned to other logical CPUs which can negate the above mentioned advantages.

## Attributes

While the size of the key and value are essentially unrestricted both [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) and [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must be at least zero and their combined size no larger than `KMALLOC_MAX_SIZE`. `KMALLOC_MAX_SIZE` is the maximum size which can be allocated by the kernel memory allocator, its exact value being dependant on a number of factors. If this edge case is hit a `-E2BIG` [error number](https://man7.org/linux/man-pages/man3/errno.3.html) is returned to the [map create syscall](../syscall/BPF_MAP_CREATE.md).

The [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) attribute indicates the max entries per-CPU so the actual memory size consumed is also dependant on the logical CPU count of the host.

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
 * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
 * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
 * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
 * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
 * [`bpf_map_lookup_percpu_elem`](../helper-function/bpf_map_lookup_percpu_elem.md)
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

While settings this flag is allowed, only a value of `-1` is allowed in the [`numa_node`](../syscall/BPF_MAP_CREATE.md#numa_node) attribute, which indicates no specific NUMA node. Since each logical CPU has its own hash table, it is impossible to allocate on only a single NUMA node.

*[NUMA]: Non-Uniform Memory Access

### `BPF_F_RDONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be read via the [syscall](../syscall/index.md) interface, but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#BPF_F_RDONLY).

### `BPF_F_WRONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be written to via the [syscall](../syscall/index.md) interface, but not read from.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#BPF_F_WRONLY).

### `BPF_F_ZERO_SEED`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/96b3b6c9091d23289721350e32c63cc8749686be)

Setting this flag will initialize the hash table with a seed of 0.

The hashing algorithm used by the hash table is seeded with a random number by default. This seeding is meant as a mitigation against Denial of Service attacks which could exploit the predictability of hashing implementations.

This random seed makes hash map operations inherently random in access time. This flag was introduced to make performance evaluation more consistent.

!!! warning
    It is not recommended to use this flag in production due to the vulnerability to Denial of Service attacks.

!!! info
    Only users with the `CAP_SYS_ADMIN` capability can use this flag, `CAP_BPF` is not enough due to the security risk associated with the flag.

### `BPF_F_RDONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be read via [helper functions](../helper-function/index.md), but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#BPF_F_RDONLY_PROG).

### `BPF_F_WRONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be written to via [helper functions](../helper-function/index.md), but not read from.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#BPF_F_RDONLY_PROG).

<!-- ## Internals -->
<!-- TODO locking / implementations -->
