---
title: "Syscall command 'BPF_PROG_GET_FD_BY_ID'"
description: "This page documents the 'BPF_PROG_GET_FD_BY_ID' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_PROG_GET_FD_BY_ID` command

<!-- [FEATURE_TAG](BPF_PROG_GET_FD_BY_ID) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/b16d9aa4c2b90af8d2c3201e245150f8c430c3bc)
<!-- [/FEATURE_TAG] -->

This syscall command is used to get a file descriptor to an already loaded program via its unique identifier.

## Return type

This command will return the file descriptor of the program (positive integer) or an error number (negative integer) if something went wrong.

## Usage

This command is used to get a file descriptor to program so that you might use it in other syscall commands or map values. This mechanism is usually used by inspection tools in combination with the [`BPF_OBJ_GET_INFO_BY_FD`](BPF_OBJ_GET_INFO_BY_FD.md).

## Attributes

### `prog_id`

This field holds the unique identifier of the program for which a file descriptor should be opened.
