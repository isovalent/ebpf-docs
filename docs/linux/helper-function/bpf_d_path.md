---
title: "Helper Function 'bpf_d_path'"
description: "This page documents the 'bpf_d_path' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_d_path`

<!-- [FEATURE_TAG](bpf_d_path) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/6e22ab9da79343532cd3cde39df25e5a5478c692)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Return full path for given **struct path** object, which needs to be the kernel BTF _path_ object. The path is returned in the provided buffer _buf_ of size _sz_ and is zero terminated.



### Returns

On success, the strictly positive length of the string, including the trailing NUL character. On error, a negative value.

`#!c static long (*bpf_d_path)(struct path *path, char *buf, __u32 sz) = (void *) 147;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
