# Helper function `bpf_dynptr_write`

<!-- [FEATURE_TAG](bpf_dynptr_write) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/13bbbfbea7598ea9f8d9c3d73bf053bb57f9c4b2)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Write _len_ bytes from _src_ into _dst_, starting from _offset_ into _dst_.

_flags_ must be 0 except for skb-type dynptrs.

For skb-type dynptrs:     _  All data slices of the dynptr are automatically        invalidated after **bpf_dynptr_write**(). This is        because writing may pull the skb and change the        underlying packet buffer.

    _  For _flags_, please see the flags accepted by        **bpf_skb_store_bytes**().

### Returns

0 on success, -E2BIG if _offset_ + _len_ exceeds the length of _dst_'s data, -EINVAL if _dst_ is an invalid dynptr or if _dst_ is a read-only dynptr or if _flags_ is not correct. For skb-type dynptrs, other errors correspond to errors returned by **bpf_skb_store_bytes**().

`#!c static long (*bpf_dynptr_write)(const struct bpf_dynptr *dst, __u32 offset, void *src, __u32 len, __u64 flags) = (void *) 202;`
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
