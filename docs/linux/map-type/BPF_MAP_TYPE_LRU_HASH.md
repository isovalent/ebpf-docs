---
title: "Map Type 'BPF_MAP_TYPE_LRU_HASH'"
description: "This page documents the 'BPF_MAP_TYPE_LRU_HASH' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_LRU_HASH`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_LRU_HASH) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/29ba732acbeece1e34c68483d1ec1f3720fa1bb3)
<!-- [/FEATURE_TAG] -->

This map is the LRU (Least Recently Used) variant of the [`BPF_MAP_TYPE_HASH`](../map-type/BPF_MAP_TYPE_HASH.md). It is a generic map type that stores a fixed maximum number of key/value pairs. When the map starts to get at capacity, the approximately least recently used elements is removed to make room for new elements.

## LRU internals

The idea behind the LRU eviction scheme is that its better to evict elements that have not been updated or looked at in a while. Users should not rely on any accuracy on the part of the eviction algorithm, the algorithm is is more approximate than exact for the sake of performance.

In the default mode (no `BPF_F_NO_COMMON_LRU` flag set) the map uses a global LRU accounting for all elements in the map. This accounting is comprised of three maps, the "active", "inactive", and "free" list. At map creation all elements preallocated and put in the "free" list. When new space is needed, one is taken from the "free" list and added to the "active" list.

Every element has a "ref" bit that tracks if the element has been recently referenced. It is set every time that element is looked-up. When a update to the map takes place, the kernel checks if the "inactive" list is "running low" (meaning # in "inactive" < # in "active"). When this is the case, list rotation will take place. The "active" list is iterated from its tail, if the "ref" bit is set, it is moved to the head of the "active" list, if the "ref" bit is not set then it is moved to the head of the "inactive" list. Then the "inactive" list is iterated from its tail, if the "ref" bit is set, it is moved to the head of the "active" list, if the "ref" bit is not set then it is moved to the head of the "inactive" list. In all cases, the "ref" bit is cleared.

When the "free" list is empty, the "inactive" list is iterated from tail to head and all elements without a set "ref" bit are deleted and moved to the "free" list. If no elements were found, the element at the tail of the "inactive" or "active" list will be freed regardless of the "ref" bit.

On top of the global LRU mechanism explained above there are also two per-CPU lists which are used as buffers so the lock on the global LRU lists is taken less often. Every CPU has its own 
"free" list and "pending" list. When a new element is needed, the CPU local "free" list is consulted first. If it is empty, the algorithm attempts to transfer `LOCAL_FREE_TARGET`(typically `128`) elements from the global "free" list to the local one. If the global "free" list does not have enough elements it will shrink the "inactive" list as described above. When, after shrinking, there are no new free elements, we iterate over all other CPUs and attempt to "steal" a free element from their local "free" lists. When this is not possible we resort to forcefully deleting the elements from the global list.

The local "pending" list stores newly added elements before they end up on the "active" list. The contents of the "pending" list are moved to the "active" list right before attempting to move elements from the global to the local "free" list.

When the `BPF_F_NO_COMMON_LRU` is set, every CPU gets its own 3-list LRU accounting and no per CPU buffers. The map keys and values are still shared between CPUs, unlike a per-CPU map. But the LRU accounting is done per CPU. This means that a CPU will only ever evict elements that it has added in the first place. This mode is mostly intended for maximum performance in cases where related requests are handled by the same CPU such as traffic for the same 5-tuple in a network application.

## Attributes

While the size of the key and value are essentially unrestricted both [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) and [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must be at least larger than zero and their combined size plus implementation overhead no larger than `KMALLOC_MAX_SIZE`. `KMALLOC_MAX_SIZE` is the maximum size which can be allocated by the kernel memory allocator, its exact value being dependant on a number of factors. If this edge case is hit a `-E2BIG` [error number](https://man7.org/linux/man-pages/man3/errno.3.html) is returned to the [map create syscall](../syscall/BPF_MAP_CREATE.md).

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)

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

### `BPF_F_NO_COMMON_LRU`
<!-- [FEATURE_TAG](BPF_F_NO_COMMON_LRU) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/29ba732acbeece1e34c68483d1ec1f3720fa1bb3)
<!-- [/FEATURE_TAG] -->

When set, every CPU gets its own LRU accounting which increases performance but also makes evictions less accurate.

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

This random seed makes hash map operations inherently random in access time. This flag was introduced to make performance evaluation more consistent.

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

