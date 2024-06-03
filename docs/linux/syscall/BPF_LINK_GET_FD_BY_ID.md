---
title: "Syscall command 'BPF_LINK_GET_FD_BY_ID'"
description: "This page documents the 'BPF_LINK_GET_FD_BY_ID' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_LINK_GET_FD_BY_ID` command

<!-- [FEATURE_TAG](BPF_LINK_GET_FD_BY_ID) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/2d602c8cf40d65d4a7ac34fe18648d8778e6e594)
<!-- [/FEATURE_TAG] -->

This syscall command is used to get a file descriptor to an already loaded link via its unique identifier.

## Return type

This command will return the file descriptor of the link (positive integer) or an error number (negative integer) if something went wrong.

## Usage

This command is used to get a file descriptor to link so that you might use it in other syscall commands. This mechanism is usually used by inspection tools in combination with the [`BPF_OBJ_GET_INFO_BY_FD`](BPF_OBJ_GET_INFO_BY_FD.md).

## Attributes

### `link_id`

This field holds the unique identifier of the link for which a file descriptor should be opened.
