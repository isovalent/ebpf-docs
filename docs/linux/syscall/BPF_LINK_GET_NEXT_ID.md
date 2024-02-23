---
title: "Syscall command 'BPF_LINK_GET_NEXT_ID'"
description: "This page documents the 'BPF_LINK_GET_NEXT_ID' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_LINK_GET_NEXT_ID` command

<!-- [FEATURE_TAG](BPF_LINK_GET_NEXT_ID) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/2d602c8cf40d65d4a7ac34fe18648d8778e6e594)
<!-- [/FEATURE_TAG] -->

This syscall command is used to iterate over all loaded links.

## Return type

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Usage

This syscall command will populate the `next_id` field with the ID of the "next" link which will have a higher number than `start_id`. If no link IDs are known, `start_id` can be left at `0`. If no links exist higher than `start_id`, `next_id` is set to `-1` and the syscall will return an `-ENOENT` error code.

So to iterate or discover all loaded links: 

1. call this command repeatably with the same attribute pointer and the attributes initialized at zero
2. move `next_id` to `start_id` between each call
3. record all `next_id` values
4. stop when we get an error

The IDs returned by this command can be used with the [`BPF_LINK_GET_FD_BY_ID`](BPF_LINK_GET_FD_BY_ID.md) syscall command to get a file descriptor to the actual link.

## Attributes

### `start_id`

The ID from which we wish to start iterating. `next_id` will always be higher than this field.

### `next_id`

This field will be set to the next link ID, or `-1` if no next link exists.
