---
title: "Syscall command 'BPF_MAP_GET_NEXT_ID'"
description: "This page documents the 'BPF_MAP_GET_NEXT_ID' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_GET_NEXT_ID` command

<!-- [FEATURE_TAG](BPF_MAP_GET_NEXT_ID) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/34ad5580f8f9c86cb273ebea25c149613cd1667e)
<!-- [/FEATURE_TAG] -->

This syscall command is used to iterate over all loaded maps.

## Return type

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Usage

This syscall command will populate the `next_id` field with the ID of the "next" map which will have a higher number than `start_id`. If no map IDs are known, `start_id` can be left at `0`. If no maps exist higher than `start_id`, `next_id` is set to `-1` and the syscall will return an `-ENOENT` error code.

So to iterate or discover all loaded maps: 

1. call this command repeatably with the same attribute pointer and the attributes initialized at zero
2. move `next_id` to `start_id` between each call
3. record all `next_id` values
4. stop when we get an error

The IDs returned by this command can be used with the [`BPF_MAP_GET_FD_BY_ID`](BPF_MAP_GET_FD_BY_ID.md) syscall command to get a file descriptor to the actual map.

## Attributes

### `start_id`

The ID from which we wish to start iterating. `next_id` will always be higher than this field.

### `next_id`

This field will be set to the next map ID, or `-1` if no next map exists.
