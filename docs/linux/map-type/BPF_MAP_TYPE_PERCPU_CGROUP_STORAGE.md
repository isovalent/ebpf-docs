---
title: "Map Type 'BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE'"
description: "This page documents the 'BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/b741f1630346defcbc8cc60f1a2bdae8b3b0036f)
<!-- [/FEATURE_TAG] -->

This is a specialized map, the per-CPU variant of [`BPF_MAP_TYPE_CGROUP_STORAGE`](BPF_MAP_TYPE_CGROUP_STORAGE.md). This map type stores data keyed on a cGroup. The user can't create or delete entries in this map directly, only read and update keys for existing cGroups. When a cGroup is deleted, all entries for that cGroup are automatically removed.

This map type can only be used with `BPF_PROG_TYPE_CGROUP_*` program types. Upon attaching a BPF program with a cGroup storage map, the cGroup becomes associated with that map. Userspace can only read or set values for cGroups associated with the map. BPF programs can use the [bpf_get_local_storage](../helper-function/bpf_get_local_storage.md) helper function to access the map value for the cGroup the program is attached to. Due to these limitations, this map type has been deprecated in favor of the [`BPF_MAP_TYPE_CGRP_STORAGE`](BPF_MAP_TYPE_CGRP_STORAGE.md) map type.

## Attributes

The [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must be `12` bytes and has the following type:
```c
struct bpf_cgroup_storage_key {
        __u64 cgroup_inode_id;
        __u32 attach_type;
};
```
After [:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/7d9c3427894fe70d1347b4820476bf37736d2ff0) the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) may also be `8` bytes, which will be interpreted as just the `cgroup_inode_id`.

The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) of the map may be any size within the limits of the kernel. [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) must be `0`, as the number of entries is determined by the number of cgroups on the system.

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [`bpf_get_local_storage`](../helper-function/bpf_get_local_storage.md)
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
