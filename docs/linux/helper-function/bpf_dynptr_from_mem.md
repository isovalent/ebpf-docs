# Helper function `bpf_dynptr_from_mem`

<!-- [FEATURE_TAG](bpf_dynptr_from_mem) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/263ae152e96253f40c2c276faad8629e096b3bad)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get a dynptr to local memory _data_.

_data_ must be a ptr to a map value. The maximum _size_ supported is DYNPTR_MAX_SIZE. _flags_ is currently unused.

### Returns

0 on success, -E2BIG if the size exceeds DYNPTR_MAX_SIZE, -EINVAL if flags is not 0.

`#!c static long (*bpf_dynptr_from_mem)(void *data, __u32 size, __u64 flags, struct bpf_dynptr *ptr) = (void *) 197;`
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
