---
title: "Syscall command 'BPF_BTF_GET_FD_BY_ID'"
description: "This page documents the 'BPF_BTF_GET_FD_BY_ID' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_BTF_GET_FD_BY_ID` command

<!-- [FEATURE_TAG](BPF_BTF_GET_FD_BY_ID) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/78958fca7ead2f81b60a6827881c4866d1ed0c52)
<!-- [/FEATURE_TAG] -->


This syscall command is used to get a file descriptor to an already loaded BTF object via its unique identifier.

## Return type

This command will return the file descriptor of the BTF object (positive integer) or an error number (negative integer) if something went wrong.

## Usage

This command is used to get a file descriptor to BTF object so that you might use it in other syscall commands. This mechanism is usually used by inspection tools in combination with the [`BPF_OBJ_GET_INFO_BY_FD`](BPF_OBJ_GET_INFO_BY_FD.md).

## Attributes

### `btf_id`

This field holds the unique identifier of the BTF object for which a file descriptor should be opened.
