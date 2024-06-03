---
title: "Helper Function 'bpf_sysctl_get_new_value'"
description: "This page documents the 'bpf_sysctl_get_new_value' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_sysctl_get_new_value`

<!-- [FEATURE_TAG](bpf_sysctl_get_new_value) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/4e63acdff864654cee0ac5aaeda3913798ee78f6)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get new value being written by user space to sysctl (before the actual write happens) and copy it as a string into provided by program buffer _buf_ of size _buf_len_.

User space may write new value at file position > 0.

The buffer is always NUL terminated, unless it's zero-sized.

### Returns

Number of character copied (not including the trailing NUL).

**-E2BIG** if the buffer wasn't big enough (_buf_ will contain truncated name in this case).

**-EINVAL** if sysctl is being read.

`#!c static long (* const bpf_sysctl_get_new_value)(struct bpf_sysctl *ctx, char *buf, unsigned long buf_len) = (void *) 103;`
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
