# Helper function `bpf_xdp_store_bytes`

<!-- [FEATURE_TAG](bpf_xdp_store_bytes) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/3f364222d032eea6b245780e845ad213dab28cdd)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Store _len_ bytes from buffer _buf_ into the frame associated to _xdp_md_, at _offset_.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_xdp_store_bytes)(struct xdp_md *xdp_md, __u32 offset, void *buf, __u32 len) = (void *) 190;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
