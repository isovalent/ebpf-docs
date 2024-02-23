---
title: "Syscall command 'BPF_MAP_CREATE'"
description: "This page documents the 'BPF_MAP_CREATE' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_CREATE` command

<!-- [FEATURE_TAG](BPF_MAP_CREATE) -->
[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/99c55f7d47c0dc6fc64729f37bf435abf43f4c60)
<!-- [/FEATURE_TAG] -->

The `BPF_MAP_CREATE` command is used to create a new BPF map.

## Return value

This command will return a file descriptor to the created map on success (positive integer) or an error number (negative integer) if something went wrong.

## Attributes

### `map_type`

This attribute specifies which type of map should be created, this should be one of the pre-defined [map types](../map-type/index.md).

### `key_size`

This attribute specifies the size of the key in bytes. 

!!! info
    Some map types have restrictions on which values are allowed, check the documentation of the specific map type for more details.

### `value_size`

This attribute specifies the size of the value in bytes. 

!!! info
    Some map types have restrictions on which values are allowed, check the documentation of the specific map type for more details.

### `max_entries`

This attribute specifies the maximum amount of entries the map can hold.

!!! info
    Some map types have restrictions on which values are allowed, check the documentation of the specific map type for more details.

### `map_flags`

This attribute is a bitmask of flags, see the [flags](#flags) section below for details.

### `inner_map_fd`
[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/56f668dfe00dcf086734f1c42ea999398fad6572)

This attribute should be set to the FD of another map when creating [map-in-map](../map-type/index.md#map-in-map) type maps. Doing so doesn't link the specified inner map to this new map we are creating, rather it is used as a mechanism to inform the kernel of the inner-maps attributes like type, key size, value size. When writing map references as values to this map, the kernel will verify that those maps are compatible with the attributes of the map given via this field.

A known technique is to create a pseudo/temporary map just for the purpose of informing this field and then releasing all references to it.

### `numa_node`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/96eabe7a40aa17e613cf3db2c742ee8b1fc764d0)

This attribute specifies on which [NUMA](https://en.wikipedia.org/wiki/Non-uniform_memory_access) node the map should be located. Memory access within the same node is typically faster, which can lead to optimization if applied correctly.

### `map_name`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/ad5b177bd73f5107d97c36f56395c4281fb6f089)

This attribute allows the map creator to give it a human readable name. The attribute is an array of 16 bytes in which a null terminated string can be placed (thus limiting the name to 15 actual characters). This name will stay associated with the map and is reported back in the results of `BPF_OBJ_GET_INFO_BY_*` syscall commands.

### `map_ifindex`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/a38845729ea3985db5d2544ec3ef3dc8f6313a27)

This attribute can be set to the index of a network interface to request that the map be offloaded to that network device. This does require that network interface to support eBPF offloading.

### `btf_fd`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/a26ca7c982cb576749cbdd01e8ecde4bf010d60a)

This attribute specifies the file descriptor of the [BTF](../../concepts/btf.md) object which contains the key and value type info which will be referenced in `btf_key_type_id` and `btf_key_value_id`.

<!-- TODO are there situations where BTF info on a map is required? spin-locks and timers for example -->

Adding BTF information about the key and value types of the map allows tools like `bpftool` to pretty-print the map keys and values instead of just the binary blobs.

### `btf_key_type_id`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/a26ca7c982cb576749cbdd01e8ecde4bf010d60a)

This attribute specifies the [BTF](../../concepts/btf.md) type ID of the map key within the [BTF](../../concepts/btf.md) object indicated by `btf_id`.

### `btf_key_value_id`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/a26ca7c982cb576749cbdd01e8ecde4bf010d60a)

This attribute specifies the [BTF](../../concepts/btf.md) type ID of the map value within the [BTF](../../concepts/btf.md) object indicated by `btf_id`.

### `btf_vmlinux_value_type_id`

[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/85d33df357b634649ddbe0a20fd2d0fc5732c3cb)

This attribute is specifically used for the `BPF_MAP_TYPE_STRUCT_OPS` map type to indicate which structure in the kernel we wish to replicate using eBPF. For more details please check the [struct ops map](../map-type/BPF_MAP_TYPE_STRUCT_OPS.md) page.

### `map_extra`

[:octicons-tag-24: v5.16](https://github.com/torvalds/linux/commit/9330986c03006ab1d33d243b7cfe598a7a3c1baa)

This attribute specifies additional settings, the meaning of which is map type specific.

It has the following meanings per map type:

* `BPF_MAP_TYPE_BLOOM_FILTER` - The lowest 4 bits indicate the number of hash functions (if 0, the bloom filter will default to using 5 hash functions).

## Flags

### `BPF_F_NO_PREALLOC`

<!-- [FEATURE_TAG](BPF_F_NO_PREALLOC) -->
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/6c90598174322b8888029e40dd84a4eb01f56afe)
<!-- [/FEATURE_TAG] -->

Before kernel version v4.6, [`BPF_MAP_TYPE_HASH`](../map-type/BPF_MAP_TYPE_HASH.md) and [`BPF_MAP_TYPE_PERCPU_HASH`](../map-type/BPF_MAP_TYPE_PERCPU_HASH.md) hash maps were lazily allocated. To improve performance, the default has been switched to pre-allocation of such map types. However, this means that for large `max_entries` values a lot of unused memory is kept in reserve. Setting this flag will not pre-allocate these maps.

Some map types require the loader to set this flag when creating maps to explicitly make clear that memory for such map types is always lazily allocated (also to guarantee stable behavior in case pre-allocation for those maps is ever added).

<!-- TODO list map types with support and link to specific pages -->

### `BPF_F_NO_COMMON_LRU`

<!-- [FEATURE_TAG](BPF_F_NO_COMMON_LRU) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/29ba732acbeece1e34c68483d1ec1f3720fa1bb3)
<!-- [/FEATURE_TAG] -->

By default, LRU maps have a single LRU list (even per-CPU LRU maps). When set, the an LRU map will use a percpu LRU list
which can scale and perform better.

!!! note
    The LRU nodes (including free nodes) cannot be moved across different LRU lists.

### `BPF_F_NUMA_NODE`

<!-- [FEATURE_TAG](BPF_F_NUMA_NODE) -->
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/96eabe7a40aa17e613cf3db2c742ee8b1fc764d0)
<!-- [/FEATURE_TAG] -->

When set, the [`numa_node`](#numa_node) attribute is respected during map creation.


### `BPF_F_RDONLY`

<!-- [FEATURE_TAG](BPF_F_RDONLY) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)
<!-- [/FEATURE_TAG] -->

Setting this flag will make it so the map can only be read via the [syscall](../syscall/index.md) interface, but not written to.

This flag is mutually exclusive with `BPF_F_WRONLY`, one of them can be used, not both.

### `BPF_F_WRONLY`

<!-- [FEATURE_TAG](BPF_F_WRONLY) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)
<!-- [/FEATURE_TAG] -->

Setting this flag will make it so the map can only be written to via the [syscall](../syscall/index.md) interface, but not read from.

This flag is mutually exclusive with `BPF_F_RDONLY`, one of them can be used, not both.

### `BPF_F_STACK_BUILD_ID`

<!-- [FEATURE_TAG](BPF_F_STACK_BUILD_ID) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/615755a77b2461ed78dfafb8a6649456201949c7)
<!-- [/FEATURE_TAG] -->

By default, `BPF_MAP_TYPE_STACK_TRACE` maps store address for each entry in the call trace. To map these addresses to user space files, it is necessary to maintain the mapping from these virtual address to symbols in the binary.

When setting this flag, the [stack trace map](../map-type/BPF_MAP_TYPE_STACK_TRACE.md) will instead store the variation stores ELF file build_id + offset.

For more details, check the [stack trace map](../map-type/BPF_MAP_TYPE_STACK_TRACE.md) map page.

### `BPF_F_ZERO_SEED`

<!-- [FEATURE_TAG](BPF_F_ZERO_SEED) -->
[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/96b3b6c9091d23289721350e32c63cc8749686be)
<!-- [/FEATURE_TAG] -->

This flag can be used in the following map types:

* [`BPF_MAP_TYPE_HASH`](../map-type/BPF_MAP_TYPE_HASH.md)
* [`BPF_MAP_TYPE_PERCPU_HASH`](../map-type/BPF_MAP_TYPE_PERCPU_HASH.md)
* [`BPF_MAP_TYPE_LRU_HASH`](../map-type/BPF_MAP_TYPE_LRU_HASH.md)
* [`BPF_MAP_TYPE_LRU_PERCPU_HASH`](../map-type/BPF_MAP_TYPE_LRU_PERCPU_HASH.md)
* [`BPF_MAP_TYPE_BLOOM_FILTER`](../map-type/BPF_MAP_TYPE_BLOOM_FILTER.md)

### `BPF_F_RDONLY_PROG`

<!-- [FEATURE_TAG](BPF_F_RDONLY_PROG) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)
<!-- [/FEATURE_TAG] -->

Setting this flag will make it so the map can only be read via [helper functions](../helper-function/index.md), but not written to.

This flag is mutually exclusive with `BPF_F_WRONLY_PROG`, one of them can be used, not both.

### `BPF_F_WRONLY_PROG`

<!-- [FEATURE_TAG](BPF_F_WRONLY_PROG) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)
<!-- [/FEATURE_TAG] -->

Setting this flag will make it so the map can only be written to via [helper functions](../helper-function/index.md), but not read from.

This flag is mutually exclusive with `BPF_F_RDONLY_PROG`, one of them can be used, not both.

### `BPF_F_CLONE`

<!-- [FEATURE_TAG](BPF_F_CLONE) -->
[:octicons-tag-24: v5.4](https://github.com/torvalds/linux/commit/8f51dfc73bf181f2304e1498f55d5f452e060cbe)
<!-- [/FEATURE_TAG] -->

This flag specifically applies to `BPF_MAP_TYPE_SK_STORAGE` maps. Sockets can be cloned. Setting this flag on the socket storage allows it to be cloned along with the socket itself when this happens. By default the storage is not cloned and the socket storage on the cloned socket will stay empty.

### `BPF_F_MMAPABLE`

<!-- [FEATURE_TAG](BPF_F_MMAPABLE) -->
[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/fc9702273e2edb90400a34b3be76f7b08fa3344b)
<!-- [/FEATURE_TAG] -->

Setting this flag on a `BPF_MAP_TYPE_ARRAY` will allow userspace programs to [mmap](https://man7.org/linux/man-pages/man2/mmap.2.html) the array values into the userspace process, effectively making a shared memory region between eBPF programs and a userspace program.

This can significantly improve read and write performance since there is no sycall overhead to access the map.

Using this flag is only supported on `BPF_MAP_TYPE_ARRAY` maps, for more details check the [array map page](../map-type/BPF_MAP_TYPE_ARRAY.md).

### `BPF_F_PRESERVE_ELEMS`

<!-- [FEATURE_TAG](BPF_F_PRESERVE_ELEMS) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/792caccc4526bb489e054f9ab61d7c024b15dea2)
<!-- [/FEATURE_TAG] -->

Maps of type `BPF_MAP_TYPE_PERF_EVENT_ARRAY` by default will clear all unread perf events when the original map file descriptor is closed, even if the map still exists. Setting this flag will make it so any pending elements will stay until explicitly removed or the map is freed. This makes sharing the perf event array between userspace programs easier.

### `BPF_F_INNER_MAP`

<!-- [FEATURE_TAG](BPF_F_INNER_MAP) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/4a8f87e60f6db40e640f1db555d063b2c4dea5f1)
<!-- [/FEATURE_TAG] -->

[Map-in-Map](../map-type/index.md#map-in-map) maps normally require that all inner maps have the same [`max_entries`](#max_entries) value and that this value matches the `max_entries` of the map specified by [`inner_map_fd`](#inner_map_fd). Setting this flag on the inner map value when loading will allow you to assign that map to the outer map even if it has a different `max_entries` value. This is at the cost of a slight hit to performance during lookups.

<!-- TODO this flag seems to only be allowed for array maps -->

*[BTF]: BPF Type Format 
*[FD]: File Descriptor

## Example

```c
union bpf_attr my_map {
    .map_type = BPF_MAP_TYPE_HASH,
    .key_size = sizeof(int),
    .value_size = sizeof(int),
    .max_entries = 100,
    .map_flags = BPF_F_NO_PREALLOC,
};
int fd = bpf(BPF_MAP_CREATE, &my_map, sizeof(my_map));
```
