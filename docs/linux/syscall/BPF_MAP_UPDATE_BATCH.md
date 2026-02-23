---
title: "Syscall command 'BPF_MAP_UPDATE_BATCH'"
description: "This page documents the 'BPF_MAP_UPDATE_BATCH' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_UPDATE_BATCH` command

<!-- [FEATURE_TAG](BPF_MAP_UPDATE_BATCH) -->
[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/aa2e93b8e58e18442edfb2427446732415bc215e)
<!-- [/FEATURE_TAG] -->

The `BPF_MAP_UPDATE_BATCH` command is used to insert or update a group of values in a map

## Return value

This command will return `0` on success or a error number (negative integer) if something went wrong.


## Attributes
### `map_fd`

This attribute specifies the file descriptor of the map in which you wish to update.


### `count`
This attributes specifies the number of elements that you wish to update

### `keys`

This attribute holds a **pointer** to memory location of the keys you wish to update. It must point to memory large enough to hold count items based on the key size. The size of the key is derived from the [`key_size`](BPF_MAP_CREATE.md#key_size) attribute of the map you specified with `map_fd`.

Each element specified in keys is sequentially updated to the value in the corresponding index in values

The memory indicated by this field will not be modified.

### `values`

This attribute holds a **pointer** to a memory location of the value you wish to write to the map. This value should be `count` * [`value_size`](BPF_MAP_CREATE.md#value_size) bytes large for regular maps and be an array of #-logical-CPUs elements of `count` * [`value_size`](BPF_MAP_CREATE.md#value_size) bytes for per-CPU maps.

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
