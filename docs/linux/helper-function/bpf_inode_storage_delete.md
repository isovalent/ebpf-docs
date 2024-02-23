---
title: "Helper Function 'bpf_inode_storage_delete' - eBPF Docs"
description: "This page documents the 'bpf_inode_storage_delete' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_inode_storage_delete`

<!-- [FEATURE_TAG](bpf_inode_storage_delete) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/8ea636848aca35b9f97c5b5dee30225cf2dd0fe6)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Delete a bpf_local_storage from an _inode_.

### Returns

0 on success.

**-ENOENT** if the bpf_local_storage cannot be found.

`#!c static int (*bpf_inode_storage_delete)(void *map, void *inode) = (void *) 146;`
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
