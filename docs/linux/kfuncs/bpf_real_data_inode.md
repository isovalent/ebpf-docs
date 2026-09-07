---
title: "KFunc 'bpf_real_data_inode'"
description: "This page documents the 'bpf_real_data_inode' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_real_data_inode`

<!-- [FEATURE_TAG](bpf_real_data_inode) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/3f8c65b06fafc3f779abda5f7b81707411d05d4c)
<!-- [/FEATURE_TAG] -->

Get the real inode hosting a file's data.

## Definition

Resolve `file` to the inode that hosts its data. For a regular file on a union/overlay file system this is the underlying (upper or lower) inode that stores the data, not the overlay inode.

Data resolution only applies to regular files. For a non-regular file (e.g. a device node, FIFO or socket) on a union/overlay file system the overlay inode itself is returned; for any file on a non-union file system the inode attached to `file` is returned.

**Parameters**

`file`: file to resolve

**Returns**

The inode hosting `file`s data, or `NULL`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct inode *bpf_real_data_inode(struct file *file)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

