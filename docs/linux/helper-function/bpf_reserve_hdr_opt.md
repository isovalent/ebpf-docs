# Helper function `bpf_reserve_hdr_opt`

<!-- [FEATURE_TAG](bpf_reserve_hdr_opt) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Reserve _len_ bytes for the bpf header option.  The space will be used by **bpf_store_hdr_opt**() later in **BPF_SOCK_OPS_WRITE_HDR_OPT_CB**.

If **bpf_reserve_hdr_opt**() is called multiple times, the total number of bytes will be reserved.

This helper can only be called during **BPF_SOCK_OPS_HDR_OPT_LEN_CB**.



### Returns

0 on success, or negative error in case of failure:

**-EINVAL** if a parameter is invalid.

**-ENOSPC** if there is not enough space in the header.

**-EPERM** if the helper cannot be used under the current _skops_**->op**.

`#!c static long (*bpf_reserve_hdr_opt)(struct bpf_sock_ops *skops, __u32 len, __u64 flags) = (void *) 144;`
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
