# Helper function `bpf_csum_update`

<!-- [FEATURE_TAG](bpf_csum_update) -->
[:octicons-tag-24: v4.9](https://github.com/torvalds/linux/commit/36bbef52c7eb646ed6247055a2acd3851e317857)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Add the checksum _csum_ into _skb_**->csum** in case the driver has supplied a checksum for the entire packet into that field. Return an error otherwise. This helper is intended to be used in combination with **bpf_csum_diff**(), in particular when the checksum needs to be updated after data has been written into the packet through direct packet access.

### Returns

The checksum on success, or a negative error code in case of failure.

`#!c static __s64 (*bpf_csum_update)(struct __sk_buff *skb, __wsum csum) = (void *) 40;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
