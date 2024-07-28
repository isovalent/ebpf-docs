---
title: "Syscall command 'BPF_OBJ_GET'"
description: "This page documents the 'BPF_OBJ_GET' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_OBJ_GET` command

<!-- [FEATURE_TAG](BPF_OBJ_GET) -->
[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/b2197755b2633e164a439682fb05a9b5ea48f706)
<!-- [/FEATURE_TAG] -->

The `BPF_OBJ_GET` command is used get a file descriptor to a BPF object from a [pin](../concepts/pinning.md).

## Return value

This command will return a file descriptor to the pinned BTF object on success (positive integer) or an error number (negative integer) if something went wrong.

## Usage

A common use case for opening such a pin is to transfer a reference to a BPF object from one process to another. The [`BPF_OBJ_PIN`](BPF_OBJ_PIN.md) syscall command can be used to pin a BPF object to the BPF file system so another process can get a reference to it with this syscall command.

Please the the [pinning concept page](../concepts/pinning.md) for more details.

## Attributes

### `filename`

This field indicates the filename of the pin we wish to open. It should be a pointer to a null terminated string. 

The filename must indicate a existing file in a directory that is within the "BPF file system" (typically mounted at `/sys/fs/bpf`).

The filename must be an absolute path, so not relative paths (paths including a `.` or `..`).

### `bpf_fd`

This field is unused for this command.

### `file_flags`

This field indicates flags that apply to the file descriptor opened with this syscall command.

#### `BPF_F_RDONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag is only applicable when opening a pin for a BPF map.

Setting this flag will make it so the process is only allowed to read from the map.

This flag is is mutually exclusive with the `BPF_F_WRONLY` flag. By default (if no flags are set) the map is opened in read and write mode.

#### `BPF_F_WRONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag is only applicable when opening a pin for a BPF map.

Setting this flag will make it so the process is only allowed to write to the map.

This flag is is mutually exclusive with the `BPF_F_RDONLY` flag. By default (if no flags are set) the map is opened in read and write mode.
