---
title: "Helper Function 'bpf_msg_apply_bytes' - eBPF Docs"
description: "This page documents the 'bpf_msg_apply_bytes' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_msg_apply_bytes`

<!-- [FEATURE_TAG](bpf_msg_apply_bytes) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/2a100317c9ebc204a166f16294884fbf9da074ce)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
For socket policies, apply the verdict of the eBPF program to the next _bytes_ (number of bytes) of message _msg_.

For example, this helper can be used in the following cases:

* A single **sendmsg**() or **sendfile**() system call
  contains multiple logical messages that the eBPF program is   supposed to read and for which it should apply a verdict. * An eBPF program only cares to read the first _bytes_ of a
  _msg_. If the message has a large payload, then setting up   and calling the eBPF program repeatedly for all bytes, even   though the verdict is already known, would create unnecessary   overhead.

When called from within an eBPF program, the helper sets a counter internal to the BPF infrastructure, that is used to apply the last verdict to the next _bytes_. If _bytes_ is smaller than the current data being processed from a **sendmsg**() or **sendfile**() system call, the first _bytes_ will be sent and the eBPF program will be re-run with the pointer for start of data pointing to byte number _bytes_ **+ 1**. If _bytes_ is larger than the current data being processed, then the eBPF verdict will be applied to multiple **sendmsg**() or **sendfile**() calls until _bytes_ are consumed.

Note that if a socket closes with the internal counter holding a non-zero value, this is not a problem because data is not being buffered for _bytes_ and is sent as it is received.

### Returns

0

`#!c static long (*bpf_msg_apply_bytes)(struct sk_msg_md *msg, __u32 bytes) = (void *) 61;`
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
