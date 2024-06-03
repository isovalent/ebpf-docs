---
title: "Helper Function 'bpf_strtol'"
description: "This page documents the 'bpf_strtol' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_strtol`

<!-- [FEATURE_TAG](bpf_strtol) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/d7a4cb9b6705a89937d12c8158a35a3145dc967a)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Convert the initial part of the string from buffer _buf_ of size _buf_len_ to a long integer according to the given base and save the result in _res_.

The string may begin with an arbitrary amount of white space (as determined by **isspace**(3)) followed by a single optional '**-**' sign.

Five least significant bits of _flags_ encode base, other bits are currently unused.

Base must be either 8, 10, 16 or 0 to detect it automatically similar to user space **strtol**(3).

### Returns

Number of characters consumed on success. Must be positive but no more than _buf_len_.

**-EINVAL** if no valid digits were found or unsupported base was provided.

**-ERANGE** if resulting value was out of range.

`#!c static long (* const bpf_strtol)(const char *buf, unsigned long buf_len, __u64 flags, long *res) = (void *) 105;`
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
