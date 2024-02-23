---
title: "Helper Function 'bpf_load_hdr_opt'"
description: "This page documents the 'bpf_load_hdr_opt' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_load_hdr_opt`

<!-- [FEATURE_TAG](bpf_load_hdr_opt) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Load header option.  Support reading a particular TCP header option for bpf program (**BPF_PROG_TYPE_SOCK_OPS**).

If _flags_ is 0, it will search the option from the _skops_**->skb_data**.  The comment in **struct bpf_sock_ops** has details on what skb_data contains under different _skops_**->op**.

The first byte of the _searchby_res_ specifies the kind that it wants to search.

If the searching kind is an experimental kind (i.e. 253 or 254 according to RFC6994).  It also needs to specify the "magic" which is either 2 bytes or 4 bytes.  It then also needs to specify the size of the magic by using the 2nd byte which is "kind-length" of a TCP header option and the "kind-length" also includes the first 2 bytes "kind" and "kind-length" itself as a normal TCP header option also does.

For example, to search experimental kind 254 with 2 byte magic 0xeB9F, the searchby_res should be [ 254, 4, 0xeB, 0x9F, 0, 0, .... 0 ].

To search for the standard window scale option (3), the _searchby_res_ should be [ 3, 0, 0, .... 0 ]. Note, kind-length must be 0 for regular option.

Searching for No-Op (0) and End-of-Option-List (1) are not supported.

_len_ must be at least 2 bytes which is the minimal size of a header option.

Supported flags:

* **BPF_LOAD_HDR_OPT_TCP_SYN** to search from the
  saved_syn packet or the just-received syn packet.



### Returns

> 0 when found, the header option is copied to _searchby_res_. The return value is the total length copied. On failure, a negative error code is returned:

**-EINVAL** if a parameter is invalid.

**-ENOMSG** if the option is not found.

**-ENOENT** if no syn packet is available when **BPF_LOAD_HDR_OPT_TCP_SYN** is used.

**-ENOSPC** if there is not enough space.  Only _len_ number of bytes are copied.

**-EFAULT** on failure to parse the header options in the packet.

**-EPERM** if the helper cannot be used under the current _skops_**->op**.

`#!c static long (*bpf_load_hdr_opt)(struct bpf_sock_ops *skops, void *searchby_res, __u32 len, __u64 flags) = (void *) 142;`
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
