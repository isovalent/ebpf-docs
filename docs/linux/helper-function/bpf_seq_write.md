---
title: "Helper Function 'bpf_seq_write'"
description: "This page documents the 'bpf_seq_write' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_seq_write`

<!-- [FEATURE_TAG](bpf_seq_write) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/492e639f0c222784e2e0f121966375f641c61b15)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
**bpf_seq_write**() uses seq_file **seq_write**() to write the data. The _m_ represents the seq_file. The _data_ and _len_ represent the data to write in bytes.

### Returns

0 on success, or a negative error in case of failure:

**-EOVERFLOW** if an overflow happened: The same object will be tried again.

`#!c static long (* const bpf_seq_write)(struct seq_file *m, const void *data, __u32 len) = (void *) 127;`
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
