---
title: "Helper Function 'bpf_msg_pull_data'"
description: "This page documents the 'bpf_msg_pull_data' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_msg_pull_data`

<!-- [FEATURE_TAG](bpf_msg_pull_data) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/015632bb30daaaee64e1bcac07570860e0bf3092)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
For socket policies, pull in non-linear data from user space for _msg_ and set pointers _msg_**->data** and _msg_\ **->data_end** to _start_ and _end_ bytes offsets into _msg_, respectively.

If a program of type **BPF_PROG_TYPE_SK_MSG** is run on a _msg_ it can only parse data that the (**data**, **data_end**) pointers have already consumed. For **sendmsg**() hooks this is likely the first scatterlist element. But for calls relying on the **sendpage** handler (e.g. **sendfile**()) this will be the range (**0**, **0**) because the data is shared with user space and by default the objective is to avoid allowing user space to modify data while (or after) eBPF verdict is being decided. This helper can be used to pull in data and to set the start and end pointer to given values. Data will be copied if necessary (i.e. if data was not linear and if start and end pointers do not point to the same chunk).

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

All values for _flags_ are reserved for future usage, and must be left at zero.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_msg_pull_data)(struct sk_msg_md *msg, __u32 start, __u32 end, __u64 flags) = (void *) 63;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
