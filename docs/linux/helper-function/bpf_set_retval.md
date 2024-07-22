---
title: "Helper Function 'bpf_set_retval'"
description: "This page documents the 'bpf_set_retval' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_set_retval`

<!-- [FEATURE_TAG](bpf_set_retval) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/b44123b4a3dcad4664d3a0f72c011ffd4c9c4d93)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Set the BPF program's return value that will be returned to the upper layers.

This helper is currently supported by cgroup programs and only by the hooks where BPF program's return value is returned to the userspace via errno.

Note that there is the following corner case where the program exports an error via bpf_set_retval but signals success via 'return 1':

&nbsp;&nbsp;&nbsp;&nbsp;bpf_set_retval(-EPERM); return 1;

In this case, the BPF program's return value will use helper's -EPERM. This still holds true for cgroup/bind{4,6} which supports extra 'return 3' success case.



### Returns

0 on success, or a negative error in case of failure.

`#!c static int (* const bpf_set_retval)(int retval) = (void *) 187;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
