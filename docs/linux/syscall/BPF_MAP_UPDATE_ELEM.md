---
title: "Syscall command 'BPF_MAP_UPDATE_ELEM'"
description: "This page documents the 'BPF_MAP_UPDATE_ELEM' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_UPDATE_ELEM` command

<!-- [FEATURE_TAG](BPF_MAP_UPDATE_ELEM) -->
[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/db20fd2b01087bdfbe30bce314a198eefedcc42e)
<!-- [/FEATURE_TAG] -->

The `BPF_MAP_UPDATE_ELEM` command is used to insert or update a single value in a map.

## Return value

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Attributes
### `map_fd`

This attribute specifies the file descriptor of the map in which you wish to update.

### `key`

This attribute holds a **pointer** to the key you wish to update. The size of the key is derived from the [`key_size`](BPF_MAP_CREATE.md#key_size) attribute of the map you specified with `map_fd`.

The memory indicated by this field will not be modified.

### `value`

This attribute holds a **pointer** to a memory location of the value you wish to write to the map. This value should be [`value_size`](BPF_MAP_CREATE.md#value_size) bytes large for regular maps and be an array of #-logical-CPUs elements of [`value_size`](BPF_MAP_CREATE.md#value_size) bytes for per-CPU maps.

### `flags`

This attribute is a bitmask of flags.

#### `BPF_ANY`

This flag has a value of `0`, so setting it together with another flag has no impact. It is meant to be used if no other flags are specified to explicitly state that the command should update the map regardless of if the key already exists or not.

#### `BPF_NOEXIST`

If this flag is set, the command will make sure that the given key doesn't exist yet. If the same key already exists when this command is executed the `-EEXIST` error number will be returned.

!!! note
    Array map types are always pre-allocated and have all keys from `0` to `max_entries`-`1` set to zero, so commands with this flag set will always fail.

#### `BPF_EXISTS`

If this flag is set, the command will make sure that the given key already exists. If no entry for this key exists, the `-ENOENT` error number will be returned.

#### `BPF_F_LOCK`

If this flag is set, the command will acquire the spin-lock of the map value we are updating. If the map contains no spin-lock in its value, `-EINVAL` will be returned by the command.
