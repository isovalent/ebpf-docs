---
title: "Helper Function 'bpf_bprm_opts_set'"
description: "This page documents the 'bpf_bprm_opts_set' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_bprm_opts_set`

<!-- [FEATURE_TAG](bpf_bprm_opts_set) -->
[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/3f6719c7b62f0327c9091e26d0da10e65668229e)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Set or clear certain options on _bprm_:

**BPF_F_BPRM_SECUREEXEC** Set the secureexec bit which sets the **AT_SECURE** auxv for glibc. The bit is cleared if the flag is not specified.

### Returns

**-EINVAL** if invalid _flags_ are passed, zero otherwise.

`#!c static long (*bpf_bprm_opts_set)(struct linux_binprm *bprm, __u64 flags) = (void *) 159;`
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
