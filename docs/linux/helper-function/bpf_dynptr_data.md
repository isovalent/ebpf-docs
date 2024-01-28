# Helper function `bpf_dynptr_data`

<!-- [FEATURE_TAG](bpf_dynptr_data) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/34d4ef5775f776ec4b0d53a02d588bf3195cada6)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get a pointer to the underlying dynptr data.

_len_ must be a statically known value. The returned data slice is invalidated whenever the dynptr is invalidated.

skb and xdp type dynptrs may not use bpf_dynptr_data. They should instead use bpf_dynptr_slice and bpf_dynptr_slice_rdwr.

### Returns

Pointer to the underlying dynptr data, NULL if the dynptr is read-only, if the dynptr is invalid, or if the offset and length is out of bounds.

`#!c static void *(*bpf_dynptr_data)(const struct bpf_dynptr *ptr, __u32 offset, __u32 len) = (void *) 203;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
