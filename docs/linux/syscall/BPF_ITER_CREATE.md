---
title: "Syscall command 'BPF_ITER_CREATE'"
description: "This page documents the 'BPF_ITER_CREATE' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_ITER_CREATE` command

<!-- [FEATURE_TAG](BPF_ITER_CREATE) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/ac51d99bf81caac8d8881fe52098948110d0de68)
<!-- [/FEATURE_TAG] -->

This syscall creates a pseudo-file for an iterator.

## Return type

This syscall returns the file descriptor of the pseudo-file.

## Usage

This syscall works specifically with iterator [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md#iterator) programs. This type of program allows you to write a BPF program type which iterates of certain kernel resources such as BPF maps, sockets, VMAs, ect.

Iteration is triggered by the reading of a pseudo-file. Where typical pseudo-files such as the once found in `/proc` and `/sys` are implemented by the kernel directly, the pseudo-files for BPF iterators are produced by the BPF program when invoked.

Typical pseudo-files have a path on the file system, a file descriptor to them is gotten via the `open` syscall. For BPF iterators this is not the case. The pseudo-file does not exist as path, rather, this syscall command is used on the BPF link of a iterator which will return a file descriptor to this file.

## Attributes

```c
union bpf_attr {
    struct {
        __u32		link_fd;
        __u32		flags;
    } iter_create;
};
```

### `link_fd`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/ac51d99bf81caac8d8881fe52098948110d0de68)

The file descriptor of the BPF link for the iterator program.

### `flags`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/ac51d99bf81caac8d8881fe52098948110d0de68)

Flags to be passed to VFS layer when creating the inode.
