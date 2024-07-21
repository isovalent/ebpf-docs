---
title: "Map Type 'BPF_MAP_TYPE_TASK_STORAGE'"
description: "This page documents the 'BPF_MAP_TYPE_TASK_STORAGE' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_TASK_STORAGE`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_TASK_STORAGE) -->
[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/4cf1bc1f10452065a29d576fc5693fc4fab5b919)
<!-- [/FEATURE_TAG] -->

This map type stores data keyed on a task. The user can only create, update or delete entries for existing tasks. When a task is deleted, the entry for that task is automatically removed.

Userspace can read or update values for any task, granted they have a PID file descriptor for that task.

## Attributes

The [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `4` indicating the key is a 32-bit unsigned integer. The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) of the map may be any size within the limits of the kernel. [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) must be `0`, as the number of entries is determined by the number of sockets on the system.

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
 * [bpf_task_storage_get](../helper-function/bpf_task_storage_get.md)
 * [bpf_task_storage_delete](../helper-function/bpf_task_storage_delete.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.

### `BPF_F_NO_PREALLOC`

[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/6c90598174322b8888029e40dd84a4eb01f56afe)

This flag indicates that values for the map are not pre-allocated on creation of the map. This flag is required for the `BPF_MAP_TYPE_TASK_STORAGE` map.
