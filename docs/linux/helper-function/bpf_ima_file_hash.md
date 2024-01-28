# Helper function `bpf_ima_file_hash`

<!-- [FEATURE_TAG](bpf_ima_file_hash) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/174b16946e39ebd369097e0f773536c91a8c1a4c)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Returns a calculated IMA hash of the _file_. If the hash is larger than _size_, then only _size_ bytes will be copied to _dst_

### Returns

The **hash_algo** is returned on success, **-EOPNOTSUP** if the hash calculation failed or **-EINVAL** if invalid arguments are passed.

`#!c static long (*bpf_ima_file_hash)(struct file *file, void *dst, __u32 size) = (void *) 193;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
