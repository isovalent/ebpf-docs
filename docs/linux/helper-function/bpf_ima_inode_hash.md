---
title: "Helper Function 'bpf_ima_inode_hash' - eBPF Docs"
description: "This page documents the 'bpf_ima_inode_hash' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_ima_inode_hash`

<!-- [FEATURE_TAG](bpf_ima_inode_hash) -->
[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/27672f0d280a3f286a410a8db2004f46ace72a17)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Returns the stored IMA hash of the _inode_ (if it's available). If the hash is larger than _size_, then only _size_ bytes will be copied to _dst_

### Returns

The **hash_algo** is returned on success, **-EOPNOTSUP** if IMA is disabled or **-EINVAL** if invalid arguments are passed.

`#!c static long (*bpf_ima_inode_hash)(struct inode *inode, void *dst, __u32 size) = (void *) 161;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
