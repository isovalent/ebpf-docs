---
title: "Syscall command 'BPF_MAP_LOOKUP_AND_DELETE_ELEM'"
description: "This page documents the 'BPF_MAP_LOOKUP_AND_DELETE_ELEM' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_LOOKUP_AND_DELETE_ELEM` command

<!-- [FEATURE_TAG](BPF_MAP_LOOKUP_AND_DELETE_ELEM) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/bd513cd08f10cbe28856f99ae951e86e86803861)
<!-- [/FEATURE_TAG] -->

The `BPF_MAP_LOOKUP_AND_DELETE_ELEM` command is used to lookup values in map and deleting them at the same time.

## Return value

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Attributes
### `map_fd`

This attribute specifies the file descriptor of the map in which you wish to lookup a value.

### `key`

This attribute holds a **pointer** to the key you wish to lookup. The size of the key is derived from the [`key_size`](BPF_MAP_CREATE.md#key_size) attribute of the map you specified with `map_fd`.

The memory indicated by this field will not be modified.

!!! note
    For `BPF_MAP_TYPE_QUEUE` and `BPF_MAP_TYPE_STACK` map types this syscall command acts as the `pop` operation, in which case this field should be `NULL`.

### `value`

This attribute holds a **pointer** to a memory location where the kernel will write the value to. The user should make sure that this memory is at least [`value_size`](BPF_MAP_CREATE.md#value_size) bytes for non-per-CPU maps and `value_size` * # of logical CPUs for per-CPU maps.

For per-CPU maps, the value is an array indexed by the logical CPU number.

!!! note
    In the case of per-CPU maps, the `value_size` is rounded up to the nearest multiple of 8 bytes. So a 12-byte value will have 4-bytes of padding between each value.

The kernel will write the value(s) to the memory indicated by this field.

!!! warning
    The kernel may overwrite other neighboring memory if incorrectly sized.

### `flags`

This attribute is a bitmask of flags.

!!! note
    For `BPF_MAP_TYPE_QUEUE` and `BPF_MAP_TYPE_STACK` map types this syscall command acts as the `pop` operation, in which case this field should be `0`. 

#### `BPF_F_LOCK`

If this flag is set, the command will acquire the spin-lock of the map value we are looking up. If the map contains no spin-lock in its value, `-EINVAL` will be returned by the command.
