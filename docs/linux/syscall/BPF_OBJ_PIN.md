---
title: "Syscall command 'BPF_OBJ_PIN'"
description: "This page documents the 'BPF_OBJ_PIN' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_OBJ_PIN` command

<!-- [FEATURE_TAG](BPF_OBJ_PIN) -->
[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/b2197755b2633e164a439682fb05a9b5ea48f706)
<!-- [/FEATURE_TAG] -->

The `BPF_OBJ_PIN` command is used to [pin](../concepts/pinning.md) a BPF object to the BPF file system.

## Return value

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Usage

A common use case for creating such a pin is to transfer a reference to a BPF object from one process to another. The [`BPF_OBJ_GET`](BPF_OBJ_GET.md) syscall command can be used to get a file descriptor from pins created with this command.

Please the the [pinning concept page](../concepts/pinning.md) for more details.

## Attributes

### `filename`

This field indicates the filename on the pin we wish to create. It should be a pointer to a null terminated string. 

The filename must indicate a non-existent file in a directory that is within the "BPF file system" (typically mounted at `/sys/fs/bpf`).

The filename must be an absolute path, so not relative paths (paths including a `.` or `..`).

### `bpf_fd`

This field indicates the file descriptor for the BPF object you wish to pin.

### `file_flags`

This field indicates flags for the creation of the file. Currently there are no valid flags for this command.

