# Helper function `bpf_msg_cork_bytes`

<!-- [FEATURE_TAG](bpf_msg_cork_bytes) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/91843d540a139eb8070bcff8aa10089164436deb)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
For socket policies, prevent the execution of the verdict eBPF program for message _msg_ until _bytes_ (byte number) have been accumulated.

This can be used when one needs a specific number of bytes before a verdict can be assigned, even if the data spans multiple **sendmsg**() or **sendfile**() calls. The extreme case would be a user calling **sendmsg**() repeatedly with 1-byte long message segments. Obviously, this is bad for performance, but it is still valid. If the eBPF program needs _bytes_ bytes to validate a header, this helper can be used to prevent the eBPF program to be called again until _bytes_ have been accumulated.

### Returns

0

`#!c static long (*bpf_msg_cork_bytes)(struct sk_msg_md *msg, __u32 bytes) = (void *) 62;`
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
