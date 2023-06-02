# Helper function `bpf_msg_push_data`

<!-- [FEATURE_TAG](bpf_msg_push_data) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/6fff607e2f14bd7c63c06c464a6f93b8efbabe28)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
For socket policies, insert _len_ bytes into _msg_ at offset _start_.

If a program of type **BPF_PROG_TYPE_SK_MSG** is run on a _msg_ it may want to insert metadata or options into the _msg_. This can later be read and used by any of the lower layer BPF hooks.

This helper may fail if under memory pressure (a malloc fails) in these cases BPF programs will get an appropriate error and BPF programs will need to handle them.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_msg_push_data)(struct sk_msg_md *msg, __u32 start, __u32 len, __u64 flags) = (void *) 90;`
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
