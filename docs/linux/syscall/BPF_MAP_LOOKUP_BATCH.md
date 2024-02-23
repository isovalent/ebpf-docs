---
title: "Syscall command 'BPF_MAP_LOOKUP_BATCH' - eBPF Docs"
description: "This page documents the 'BPF_MAP_LOOKUP_BATCH' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_LOOKUP_BATCH` command

<!-- [FEATURE_TAG](BPF_MAP_LOOKUP_BATCH) -->
[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/cb4d03ab499d4c040f4ab6fd4389d2b49f42b5a5)
<!-- [/FEATURE_TAG] -->

The `BPF_MAP_LOOKUP_BATCH` command is used to iterate over maps in batches. This can be significantly faster than doing [`BPF_MAP_GET_NEXT_KEY`](BPF_MAP_GET_NEXT_KEY.md) and [`BPF_MAP_LOOKUP_ELEM`](BPF_MAP_LOOKUP_ELEM.md)  calls.

!!! note
    This command is less widely supported, than the normal single-key lookup command. Check the page of a given map type to check for compatibility.

## Return value

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Attributes
### `map_fd`

This attribute specifies the file descriptor of the map in which you wish to lookup a value.

### `in_batch`

This attribute is set by userspace. At the start of iteration it should be set to `0`/`NULL`. After the first call, the result of `out_batch` can me moved to this attribute to get the next batch of keys and values. The value is opaque, the kernel as internally assigned meaning, but userspace should just copy the values without modification.

### `out_batch`

This attribute is set by the kernel during the command exectution. At the start of iteration it should be set to `0`/`NULL`. The kernel will 

### `keys`

This attribute is an output with this command unlike with the non-batched lookup command. It holds a **pointer** to memory where a batch of keys will be written to. The size of this memory is should at least be the [`key_size`](BPF_MAP_CREATE.md#key_size) attribute of the map you specified with `map_fd` times the [`count`](#count) specified.

### `values`

This attribute holds a **pointer** to a memory location where the kernel will write the batch of values to. The user should make sure that this memory is at least [`value_size`](BPF_MAP_CREATE.md#value_size) times [`count`](#count) bytes for non-per-CPU maps and [`value_size`](BPF_MAP_CREATE.md#value_size) * # of logical CPUs * [`count`](#count) for per-CPU maps.

For per-CPU maps, the value is a two dimensional array indexed by batch index and then the logical CPU number.

!!! note
    In the case of per-CPU maps, the [`value_size`](BPF_MAP_CREATE.md#value_size) is rounded up to the nearest multiple of 8 bytes. So a 12-byte value will have 4-bytes of padding between each value.

The kernel will write the value(s) to the memory indicated by this field.

!!! warning
    The kernel may overwrite other neighboring memory if incorrectly sized.

### `count`

This attribute is both an input and output. When calling the command it specifies the maximum amount of keys and values to be retrieved. The kernel will modify the field to reflect the actual count of keys and values that were copied, which might be less than the maximum.

### `elem_flags`

This attribute is a bitmask of flags applied internally for each element lookup.

#### `BPF_F_LOCK`

If this flag is set, the command will acquire the spin-lock of the map value we are looking up. If the map contains no spin-lock in its value, `-EINVAL` will be returned by the command.

### `flags`

This attribute is a bitmask of flags for the batch operation as a whole, there are currently no valid flags defined.
