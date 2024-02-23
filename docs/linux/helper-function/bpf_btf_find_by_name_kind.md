---
title: "Helper Function 'bpf_btf_find_by_name_kind'"
description: "This page documents the 'bpf_btf_find_by_name_kind' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_btf_find_by_name_kind`

<!-- [FEATURE_TAG](bpf_btf_find_by_name_kind) -->
[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/3d78417b60fba249cc555468cb72d96f5cde2964)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Find BTF type with given name and kind in vmlinux BTF or in module's BTFs.

### Returns

Returns btf_id and btf_obj_fd in lower and upper 32 bits.

`#!c static long (*bpf_btf_find_by_name_kind)(char *name, int name_sz, __u32 kind, int flags) = (void *) 167;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
