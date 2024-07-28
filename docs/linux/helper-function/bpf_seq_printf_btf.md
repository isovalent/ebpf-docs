---
title: "Helper Function 'bpf_seq_printf_btf'"
description: "This page documents the 'bpf_seq_printf_btf' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_seq_printf_btf`

<!-- [FEATURE_TAG](bpf_seq_printf_btf) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/eb411377aed9e27835e77ee0710ee8f4649958f3)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Use BTF to write to seq_write a string representation of _ptr_->ptr, using _ptr_->type_id as per bpf_snprintf_btf(). _flags_ are identical to those used for bpf_snprintf_btf.

### Returns

0 on success or a negative error in case of failure.

`#!c static long (* const bpf_seq_printf_btf)(struct seq_file *m, struct btf_ptr *ptr, __u32 ptr_size, __u64 flags) = (void *) 150;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
