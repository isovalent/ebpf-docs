---
title: "Map Type 'BPF_MAP_TYPE_CGRP_STORAGE'"
description: "This page documents the 'BPF_MAP_TYPE_CGRP_STORAGE' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_CGRP_STORAGE`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_CGRP_STORAGE) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/c4bcfb38a95edb1021a53f2d0356a78120ecfbe4)
<!-- [/FEATURE_TAG] -->

This is a specialized map, an improved version of [`BPF_MAP_TYPE_CGROUP_STORAGE`](BPF_MAP_TYPE_CGROUP_STORAGE.md). This map type stores data keyed on a cGroup. When a cGroup is deleted, the entry for that cGroup is automatically removed.

Unlike the deprecated `BPF_MAP_TYPE_CGROUP_STORAGE`, this map type can be used with any program type. BPF programs can use the [bpf_cgrp_storage_get](../helper-function/bpf_cgrp_storage_get.md) and [bpf_cgrp_storage_delete](../helper-function/bpf_cgrp_storage_delete.md) helper functions to access any map value as long as they can obtain a pointer to a cGroup.

Userspace can read or update values for any cGroup value, granted they have a file descriptor for that cGroup.

## Attributes

The [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `4` indicating the key is a 32-bit unsigned integer. The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) of the map may be any size within the limits of the kernel. [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) must be `0`, as the number of entries is determined by the number of cgroups on the system.

This map type also requires the usage of BTF key and value types.

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
 * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.


### `BPF_F_NO_PREALLOC`

[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/6c90598174322b8888029e40dd84a4eb01f56afe)

This flag indicates that values for the map are not pre-allocated on creation of the map. This flag is required for the `BPF_MAP_TYPE_CGRP_STORAGE` map.


