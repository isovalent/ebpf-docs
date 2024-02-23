---
title: "Helper Function 'bpf_store_hdr_opt'"
description: "This page documents the 'bpf_store_hdr_opt' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_store_hdr_opt`

<!-- [FEATURE_TAG](bpf_store_hdr_opt) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Store header option.  The data will be copied from buffer _from_ with length _len_ to the TCP header.

The buffer _from_ should have the whole option that includes the kind, kind-length, and the actual option data.  The _len_ must be at least kind-length long.  The kind-length does not have to be 4 byte aligned.  The kernel will take care of the padding and setting the 4 bytes aligned value to th->doff.

This helper will check for duplicated option by searching the same option in the outgoing skb.

This helper can only be called during **BPF_SOCK_OPS_WRITE_HDR_OPT_CB**.



### Returns

0 on success, or negative error in case of failure:

**-EINVAL** If param is invalid.

**-ENOSPC** if there is not enough space in the header. Nothing has been written

**-EEXIST** if the option already exists.

**-EFAULT** on failure to parse the existing header options.

**-EPERM** if the helper cannot be used under the current _skops_**->op**.

`#!c static long (*bpf_store_hdr_opt)(struct bpf_sock_ops *skops, const void *from, __u32 len, __u64 flags) = (void *) 143;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
