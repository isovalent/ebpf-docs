---
title: "Helper Function 'bpf_kallsyms_lookup_name'"
description: "This page documents the 'bpf_kallsyms_lookup_name' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_kallsyms_lookup_name`

<!-- [FEATURE_TAG](bpf_kallsyms_lookup_name) -->
[:octicons-tag-24: v5.16](https://github.com/torvalds/linux/commit/d6aef08a872b9e23eecc92d0e92393473b13c497)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get the address of a kernel symbol, returned in _res_. _res_ is set to 0 if the symbol is not found.

### Returns

On success, zero. On error, a negative value.

**-EINVAL** if _flags_ is not zero.

**-EINVAL** if string _name_ is not the same size as _name_sz_.

**-ENOENT** if symbol is not found.

**-EPERM** if caller does not have permission to obtain kernel address.

`#!c static long (* const bpf_kallsyms_lookup_name)(const char *name, int name_sz, int flags, __u64 *res) = (void *) 179;`
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
