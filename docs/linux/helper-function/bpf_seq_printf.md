---
title: "Helper Function 'bpf_seq_printf'"
description: "This page documents the 'bpf_seq_printf' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_seq_printf`

<!-- [FEATURE_TAG](bpf_seq_printf) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/492e639f0c222784e2e0f121966375f641c61b15)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
**bpf_seq_printf**() uses seq_file **seq_printf**() to print out the format string. The _m_ represents the seq_file. The _fmt_ and _fmt_size_ are for the format string itself. The _data_ and _data_len_ are format string arguments. The _data_ are a **u64** array and corresponding format string values are stored in the array. For strings and pointers where pointees are accessed, only the pointer values are stored in the _data_ array. The _data_len_ is the size of _data_ in bytes - must be a multiple of 8.

Formats **%s**, **%p{i,I}{4,6}** requires to read kernel memory. Reading kernel memory may fail due to either invalid address or valid address but requiring a major memory fault. If reading kernel memory fails, the string for **%s** will be an empty string, and the ip address for **%p{i,I}{4,6}** will be 0. Not returning error to bpf program is consistent with what **bpf_trace_printk**() does for now.

### Returns

0 on success, or a negative error in case of failure:

**-EBUSY** if per-CPU memory copy buffer is busy, can try again by returning 1 from bpf program.

**-EINVAL** if arguments are invalid, or if _fmt_ is invalid/unsupported.

**-E2BIG** if _fmt_ contains too many format specifiers.

**-EOVERFLOW** if an overflow happened: The same object will be tried again.

`#!c static long (* const bpf_seq_printf)(struct seq_file *m, const char *fmt, __u32 fmt_size, const void *data, __u32 data_len) = (void *) 126;`
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
