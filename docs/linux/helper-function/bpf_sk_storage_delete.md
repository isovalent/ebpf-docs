---
title: "Helper Function 'bpf_sk_storage_delete'"
description: "This page documents the 'bpf_sk_storage_delete' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_sk_storage_delete`

<!-- [FEATURE_TAG](bpf_sk_storage_delete) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/6ac99e8f23d4b10258406ca0dd7bffca5f31da9d)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Delete a bpf-local-storage from a _sk_.

### Returns

0 on success.

**-ENOENT** if the bpf-local-storage cannot be found. **-EINVAL** if sk is not a fullsock (e.g. a request_sock).

`#!c static long (*bpf_sk_storage_delete)(void *map, void *sk) = (void *) 108;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
