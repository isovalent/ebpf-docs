---
title: "Helper Function 'bpf_read_branch_records'"
description: "This page documents the 'bpf_read_branch_records' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_read_branch_records`

<!-- [FEATURE_TAG](bpf_read_branch_records) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/fff7b64355eac6e29b50229ad1512315bc04b44e)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
For an eBPF program attached to a perf event, retrieve the branch records (**struct perf_branch_entry**) associated to _ctx_ and store it in the buffer pointed by _buf_ up to size _size_ bytes.

### Returns

On success, number of bytes written to _buf_. On error, a negative value.

The _flags_ can be set to **BPF_F_GET_BRANCH_RECORDS_SIZE** to instead return the number of bytes required to store all the branch entries. If this flag is set, _buf_ may be NULL.

**-EINVAL** if arguments invalid or **size** not a multiple of **sizeof**(**struct perf_branch_entry**).

**-ENOENT** if architecture does not support branch records.

`#!c static long (* const bpf_read_branch_records)(struct bpf_perf_event_data *ctx, void *buf, __u32 size, __u64 flags) = (void *) 119;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
