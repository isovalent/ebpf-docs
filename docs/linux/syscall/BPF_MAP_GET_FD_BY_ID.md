---
title: "Syscall command 'BPF_MAP_GET_FD_BY_ID'"
description: "This page documents the 'BPF_MAP_GET_FD_BY_ID' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_GET_FD_BY_ID` command

<!-- [FEATURE_TAG](BPF_MAP_GET_FD_BY_ID) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/bd5f5f4ecb78e2698dad655645b6d6a2f7012a8c)
<!-- [/FEATURE_TAG] -->


This syscall command is used to get a file descriptor to an already existing map via its unique identifier.

## Return type

This command will return the file descriptor of the map (positive integer) or an error number (negative integer) if something went wrong.

## Usage

This command is used to get a file descriptor to map so that you might use it in other syscall commands or map values. This mechanism is usually used by inspection tools in combination with the [`BPF_OBJ_GET_INFO_BY_FD`](BPF_OBJ_GET_INFO_BY_FD.md).

## Attributes

### `map_id`

This field holds the unique identifier of the map for which a file descriptor should be opened.
